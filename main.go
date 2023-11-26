package main

import (
	"context"
	"fmt"
	jira "github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	search        textinput.Model
	searchResults table.Model
	cache         *JiraCache
}

func initialModel(cache *JiraCache) model {
	m := model{
		cache:         cache,
		search:        textinput.New(),
		searchResults: table.New(),
	}

	m.search.Placeholder = "Search"
	m.search.Focus()
	m.search.CharLimit = 80
	m.search.Width = 80
	m.search.PromptStyle = focusedStyle
	m.search.TextStyle = focusedStyle
	m.search.Cursor.Style = cursorStyle

	m.searchResults.SetColumns([]table.Column{
		{Title: "Key", Width: 10},
		{Title: "Summary", Width: 70},
	})

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "alt+left", "alt+right":
			return m, tea.Batch()
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	_, cmd := m.search.Update(msg)
	search, err := m.cache.Search(m.search.Value())
	if err == nil {
		var rows []table.Row
		for _, hit := range search.Hits {
			issue := m.cache.Issues[hit.ID]
			rows = append(rows, table.Row{issue.Key, issue.Fields.Summary})
		}
		m.searchResults.SetRows(rows)
	}
	return tea.Batch(cmd)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(m.search.View())
	b.WriteString("\n\n")

	//b.WriteString(m.searchResults.View())

	return b.String()
}

func main() {
	initialize()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	cache, err := buildCache(ctx)
	if err != nil {
		fmt.Printf("could not start cache loader: %s\n", err)
		os.Exit(1)
	}
	if _, err := tea.NewProgram(initialModel(cache)).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

func buildCache(ctx context.Context) (j *JiraCache, err error) {
	j = &JiraCache{
		Issues:     make(map[string]jira.Issue),
		maxResults: viper.GetInt("jira.max_results"),
		projects:   make([]string, 0),
	}

	j.mapping = bleve.NewIndexMapping()
	j.index, err = bleve.NewMemOnly(j.mapping)
	if err != nil {
		return
	}

	if viper.GetBool("jira.fake") {

		issues := []jira.Issue{
			{
				Key: "TEST-1",
				Fields: &jira.IssueFields{
					Summary: "Test Issue 1",
				},
			},
			{
				Key: "TEST-2",
				Fields: &jira.IssueFields{
					Summary: "Test Issue 2",
				},
			},
			{
				Key: "TEST-3",
				Fields: &jira.IssueFields{
					Summary: "Test Issue 3",
				},
				Changelog: &jira.Changelog{
					Histories: nil,
				},
			},
		}
		err = j.addIssues(issues)
		return
	}

	tp := jira.BearerAuthTransport{
		Token: viper.GetString("jira.token"),
	}

	client, err := jira.NewClient(viper.GetString("jira.url"), tp.Client())
	if err != nil {
		return
	}
	j.client = client

	err = j.FetchIssues(ctx, viper.GetString("jira.initial_updated_since"))
	if err != nil {
		return
	}

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(viper.GetDuration("jira.refresh_interval")):
				err := j.FetchIssues(ctx, viper.GetString("jira.refresh_updated_since"))
				if err != nil {
					log.Fatal("could not fetch issues", err)
				}
			}
		}
	}(ctx)

	return
}

type JiraCache struct {
	Issues     map[string]jira.Issue
	client     *jira.Client
	projects   []string
	maxResults int
	index      bleve.Index
	mapping    *mapping.IndexMappingImpl
}

func (j *JiraCache) Search(search string) (*bleve.SearchResult, error) {
	result, err := j.index.Search(bleve.NewSearchRequest(bleve.NewQueryStringQuery(search)))
	if err != nil {
		log.Error("could not search", err)
		return result, err
	}
	return result, err
}

func (j *JiraCache) FetchIssues(ctx context.Context, since string) error {
	if j.client == nil {
		log.Error("client not initialized")
		return fmt.Errorf("client not initialized")
	}
	query := fmt.Sprintf(`project in (%v) AND updated >= %s ORDER BY updated DESC`,
		toQuotedList(j.projects), since)
	issues, _, err := j.client.Issue.Search(ctx, query,
		&jira.SearchOptions{MaxResults: j.maxResults})

	if err != nil {
		log.Error("could not fetch issues", err)
		return err
	}

	err = j.addIssues(issues)
	if err != nil {
		log.Error("could not add issues to cache", err)
		return err
	}
	return nil
}

func (j *JiraCache) addIssues(issues []jira.Issue) error {
	for _, issue := range issues {
		j.Issues[issue.Key] = issue
		err := j.index.Index(issue.Key, issue)
		if err != nil {
			return err
		}
	}
	return nil
}

// toQuotedList returns a string of projects in the format "PROJECT1","PROJECT2",...
func toQuotedList(projects []string) string {
	var projectsString string
	for _, project := range projects {
		projectsString += "\"" + project + "\","
	}
	return strings.TrimRight(projectsString, ",")
}

func initialize() {
	log.SetReportCaller(true)
	viper.SetEnvPrefix("JIRIO")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/jirio/")
	viper.AddConfigPath("$HOME/.jirio")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn("could not read config file", err)
	}

}
