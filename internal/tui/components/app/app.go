package app

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/gitu/jirio/internal/jiracache"
	"github.com/gitu/jirio/internal/tui/components/header"
	"github.com/gitu/jirio/internal/tui/constants"
	"github.com/gitu/jirio/internal/tui/jirio"
	"github.com/spf13/viper"
	"os/exec"
	"runtime"
	"strings"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	cursorStyle  = focusedStyle.Copy()
)

type Model struct {
	search        textinput.Model
	searchResults table.Model
	cache         jiracache.JiraCache
	header        header.Model
	currentPage   jirio.Page
}

func InitialModel(cache jiracache.JiraCache) Model {

	initialHeader := header.New(
		constants.LogoString,
		cache.Url(),
		"XXX",
	)
	m := Model{
		cache:         cache,
		search:        textinput.New(),
		searchResults: table.New(),
		header:        initialHeader,
	}

	m.search.Placeholder = "Search"
	m.search.Focus()
	m.search.CharLimit = 80
	m.search.Width = 80
	m.search.PromptStyle = focusedStyle
	m.search.TextStyle = focusedStyle
	m.search.Cursor.Style = cursorStyle

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	m.searchResults.SetStyles(s)

	m.searchResults.SetColumns([]table.Column{
		{Title: "Key", Width: 10},
		{Title: "Summary", Width: 40},
	})
	m.searchResults.SetHeight(10)

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.searchResults.SelectedRow() != nil {
				err := open(viper.GetString("jira.url") + "/browse/" + m.searchResults.SelectedRow()[0])
				if err != nil {
					log.Error("could not open browser", err)
				}

				return m, tea.Batch(tea.Printf("Let's go to %s!", m.searchResults.SelectedRow()[0]))
			}
			break
		case "backspace":
			if !m.search.Focused() {
				m.search.Focus()
				m.searchResults.Blur()
				m.searchResults.SetCursor(0)
			}
		case "tab":
			if m.search.Focused() {
				m.search.Blur()
				m.searchResults.Focus()
			} else {
				m.searchResults.Blur()
				m.search.Focus()
			}
			return m, tea.Batch()
		case "down":
			if m.search.Focused() {
				m.search.Blur()
				m.searchResults.Focus()
				return m, tea.Batch()
			}
		case "up":
			if m.searchResults.Focused() && m.searchResults.Cursor() == 0 {
				m.searchResults.Blur()
				m.search.Focus()
				return m, tea.Batch()
			}
		case "alt+left", "alt+right":
			return m, tea.Batch()
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmd, cmdt tea.Cmd
	m.search, cmd = m.search.Update(msg)
	search, err := m.cache.Search(m.search.Value())
	if err == nil {
		var rows []table.Row
		for _, hit := range search.Hits {
			issue, _ := m.cache.GetIssue(hit.ID)
			rows = append(rows, table.Row{issue.Key, issue.Fields.Summary})
		}
		m.searchResults.SetRows(rows)
		m.searchResults, cmdt = m.searchResults.Update(msg)
		return tea.Batch(cmd, cmdt)
	}
	return tea.Batch(cmd)

}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(m.header.View())
	b.WriteString("\n")

	b.WriteString(m.search.View())
	b.WriteString("\n")

	b.WriteString(baseStyle.Render(m.searchResults.View()))
	b.WriteString("\n")

	return b.String()
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
