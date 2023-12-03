package header

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/gitu/jirio/internal/tui/keymap"
	"github.com/gitu/jirio/internal/tui/style"
)

type Model struct {
	selectedQuery string
	compact       bool
	helpModel     help.Model
}

func New(helpModel help.Model) (m Model) {
	return Model{selectedQuery: "none", helpModel: helpModel}
}

func (m Model) View() string {
	logoStyle := style.Logo.Copy()
	clusterUrl := style.JiraUrl.Render(m.selectedQuery)
	if m.compact {
		return lipgloss.JoinHorizontal(
			lipgloss.Center,
			logoStyle.Padding(0).Margin(0).Render("JIRIO"),
			style.KeyHelp.Render(m.helpModel.ShortHelpView(keymap.KeyMap.ShortHelp())),
			clusterUrl,
		)
	}
	logo := logoStyle.Render(
		"░▀▀█░▀█▀░█▀▄░▀█▀░█▀█\n" +
			"░░░█░░█░░█▀▄░░█░░█░█\n" +
			"░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀",
	)
	left := style.Header.Render(lipgloss.JoinVertical(lipgloss.Center, logo, clusterUrl))
	styledKeyHelp := style.KeyHelp.Render(m.helpModel.FullHelpView(keymap.KeyMap.FullHelp()))
	return lipgloss.JoinHorizontal(lipgloss.Center, left, styledKeyHelp)
}

func (m Model) ViewHeight() int {
	return lipgloss.Height(m.View())
}

func (m *Model) ToggleCompact() {
	m.compact = !m.compact
}

func (m *Model) SetSelectedQuery(s string) {
	m.selectedQuery = s
}
