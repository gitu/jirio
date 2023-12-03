package app

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/gitu/jirio/internal/dev"
	"github.com/gitu/jirio/internal/jiracache"
	"github.com/gitu/jirio/internal/tui/components/header"
	"github.com/gitu/jirio/internal/tui/keymap"
	"github.com/gitu/jirio/internal/tui/style"
	"github.com/spf13/viper"
	"os/exec"
	"runtime"
	"sort"
	"strings"
)

type Model struct {
	search        textinput.Model
	searchResults table.Model
	header        header.Model
	cache         jiracache.JiraCache
	selectedQuery int
	queries       map[int]string
	querylist     []string
}

func InitialModel(cache jiracache.JiraCache) Model {

	initialHeader := header.New(help.New())
	m := Model{
		header:        initialHeader,
		cache:         cache,
		search:        textinput.New(),
		searchResults: table.New(),
		queries:       make(map[int]string),
	}

	m.search.Placeholder = "Search"
	m.search.Focus()
	m.search.CharLimit = 80
	m.search.Width = 80
	m.search.PromptStyle = style.FocusedStyle
	m.search.TextStyle = style.FocusedStyle
	m.search.Cursor.Style = style.FocusedStyle

	s := table.DefaultStyles()
	s.Header = style.TableHeaderStyle
	s.Selected = style.TableSelectedStyle
	s.Cell = style.TableCellStyle
	m.searchResults.SetStyles(s)

	m.searchResults.SetColumns([]table.Column{
		{Title: "Key", Width: 15},
		{Title: "Summary", Width: 80},
	})

	queries := cache.Queries()
	for _, k := range sortedKeys(queries) {
		m.queries[len(m.querylist)] = queries[k]
		m.querylist = append(m.querylist, k)
	}

	if len(m.querylist) > 1 {
		keymap.KeyMap.Right.SetEnabled(true)
	}
	m.header.SetSelectedQuery(m.queries[m.selectedQuery])

	m = m.loadAllIssues(0)
	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, keymap.KeyMap.Exit) {
			return m, tea.Quit
		} else if key.Matches(msg, keymap.KeyMap.Up) {
			if m.searchResults.Focused() && m.searchResults.Cursor() == 0 {
				m.searchResults.Blur()
				m.search.Focus()
			}
		} else if key.Matches(msg, keymap.KeyMap.Down) {
			if m.search.Focused() {
				m.search.Blur()
				m.searchResults.Focus()
			}
		} else if key.Matches(msg, keymap.KeyMap.Compact) {
			prevHeight := m.header.ViewHeight()
			m.header.ToggleCompact()
			m.searchResults.SetHeight(m.searchResults.Height() + prevHeight - m.header.ViewHeight())
			return m, nil
		} else if key.Matches(msg, keymap.KeyMap.Back) {
			if !m.search.Focused() {
				m.search.Focus()
				m.searchResults.Blur()
				m.searchResults.SetCursor(0)
			}
		} else if key.Matches(msg, keymap.KeyMap.OpenIssue) {
			if m.searchResults.SelectedRow() != nil {
				err := open(viper.GetString("jira.url") + "/browse/" + m.searchResults.SelectedRow()[0])
				if err != nil {
					log.Error("could not open browser", err)
				}

				return m, tea.Batch(tea.Printf("Let's go to %s!", m.searchResults.SelectedRow()[0]))
			}
		} else if key.Matches(msg, keymap.KeyMap.Left) {
			m.selectedQuery--
			m.checkQuerybounds()
			m.header.SetSelectedQuery(m.queries[m.selectedQuery])
			return m, m.searchIssues(m.search.Value())
		} else if key.Matches(msg, keymap.KeyMap.Right) {
			m.selectedQuery++
			m.checkQuerybounds()
			m.header.SetSelectedQuery(m.queries[m.selectedQuery])
			return m, m.searchIssues(m.search.Value())
		} else if key.Matches(msg, keymap.KeyMap.Refresh) {
			return m, m.refreshIssues()
		}
		break
	case tea.WindowSizeMsg:
		m.searchResults.SetWidth(msg.Width)
		m.searchResults.SetHeight(msg.Height - m.header.ViewHeight() - 6)
		break
	case searchResults:
		m = m.loadIssues(msg)
		break
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd

	// Handle character input and blinking
	prevVal := m.search.Value()
	m.search, cmd = m.search.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	if m.search.Value() != prevVal {
		cmds = append(cmds, m.searchIssues(m.search.Value()))
	}
	// handle search result updates
	m.searchResults, cmd = m.searchResults.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) loadIssues(msg searchResults) Model {
	m.searchResults.SetRows([]table.Row{})
	rows := []table.Row{}
	for _, issue := range msg.issues {
		rows = append(rows, table.Row{issue.Key, issue.Fields.Summary})
	}
	m.searchResults.SetRows(rows)
	return m
}

func (m *Model) checkQuerybounds() {
	keymap.KeyMap.Left.SetEnabled(true)
	keymap.KeyMap.Right.SetEnabled(true)
	if m.selectedQuery == 0 {
		keymap.KeyMap.Left.SetEnabled(false)
	}
	if m.selectedQuery == len(m.querylist)-1 {
		keymap.KeyMap.Right.SetEnabled(false)
	}
}

type searchResults struct {
	issues []jira.Issue
}

func (m Model) refreshIssues() tea.Cmd {
	return func() tea.Msg {
		err := m.cache.RefreshIssues()
		if err != nil {
			dev.Debug(fmt.Sprintf("ERROR: refreshIssues %v", err))
			return nil
		}
		return nil
	}
}

func (m Model) searchIssues(value string) tea.Cmd {
	return func() tea.Msg {
		if value == "" {
			return searchResults{issues: m.cache.GetAllIssues(m.querylist[m.selectedQuery])}
		}
		search, err := m.cache.Search(m.querylist[m.selectedQuery], value)
		if err != nil {
			dev.Debug(fmt.Sprintf("ERROR: searchIssues %v", err))
			return nil
		} else {
			issues := make([]jira.Issue, len(search.Hits))
			for i, issue := range search.Hits {
				issues[i] = m.cache.GetIssue(issue.ID)
			}
			return searchResults{issues: issues}
		}
	}
}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(m.header.View())
	b.WriteString("\n")

	b.WriteString(m.search.View())
	b.WriteString("\n")

	b.WriteString(m.searchResults.View())
	return b.String()
}

func (m Model) loadAllIssues(i int) Model {
	queryKey := m.querylist[i]
	msg := searchResults{
		issues: m.cache.GetAllIssues(queryKey),
	}
	m = m.loadIssues(msg)
	return m
}

func sortedKeys(queries map[string]string) []string {
	keys := make([]string, 0, len(queries))
	for k := range queries {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
