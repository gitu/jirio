package header

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/gitu/jirio/internal/tui/style"
)

type Model struct {
	logo, jiraUrl, keyHelp string
	compact                bool
}

func New(logo string, jiraUrl, keyHelp string) (m Model) {
	return Model{logo: logo, jiraUrl: jiraUrl, keyHelp: keyHelp}
}

func (m Model) View() string {
	logoStyle := style.Logo.Copy()
	clusterUrl := style.ClusterUrl.Render(m.jiraUrl)
	if m.compact {
		return lipgloss.JoinHorizontal(
			lipgloss.Center,
			logoStyle.Padding(0).Margin(0).Render("JIRIO"),
			style.KeyHelp.Render(m.keyHelp),
			clusterUrl,
		)
	}
	logo := logoStyle.Render(m.logo)
	left := style.Header.Render(lipgloss.JoinVertical(lipgloss.Center, logo, clusterUrl))
	styledKeyHelp := style.KeyHelp.Render(m.keyHelp)
	return lipgloss.JoinHorizontal(lipgloss.Center, left, styledKeyHelp)
}

func (m Model) ViewHeight() int {
	return lipgloss.Height(m.View())
}

func (m *Model) SetKeyHelp(keyHelp string) {
	m.keyHelp = keyHelp
}

func (m *Model) ToggleCompact() {
	m.compact = !m.compact
}
