package header

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/gitu/jirio/internal/tui/style"
)

type Model struct {
	logo, jiraUrl string
	compact       bool
	bindings      help.KeyMap
	helpModel     help.Model
}

func New(logo string, jiraUrl string, keyMap help.KeyMap, helpModel help.Model) (m Model) {
	return Model{logo: logo, jiraUrl: jiraUrl, bindings: keyMap, helpModel: helpModel}
}

func (m Model) View() string {
	logoStyle := style.Logo.Copy()
	clusterUrl := style.JiraUrl.Render(m.jiraUrl)
	if m.compact {
		return lipgloss.JoinHorizontal(
			lipgloss.Center,
			logoStyle.Padding(0).Margin(0).Render("JIRIO"),
			style.KeyHelp.Render(m.helpModel.ShortHelpView(m.bindings.ShortHelp())),
			clusterUrl,
		)
	}
	logo := logoStyle.Render(m.logo)
	left := style.Header.Render(lipgloss.JoinVertical(lipgloss.Center, logo, clusterUrl))
	styledKeyHelp := style.KeyHelp.Render(m.helpModel.FullHelpView(m.bindings.FullHelp()))
	return lipgloss.JoinHorizontal(lipgloss.Center, left, styledKeyHelp)
}

func (m Model) ViewHeight() int {
	return lipgloss.Height(m.View())
}

func (m *Model) ToggleCompact() {
	m.compact = !m.compact
}
