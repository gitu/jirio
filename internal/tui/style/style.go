package style

import "github.com/charmbracelet/lipgloss"

const (
	black  = lipgloss.Color("#000000")
	blue   = lipgloss.Color("6")
	pink   = lipgloss.Color("#E760FC")
	yellow = lipgloss.Color("#DBBD70")
)

var (
	Regular = lipgloss.NewStyle()
	Bold    = Regular.Copy().Bold(true)
	Logo    = Regular.Copy().Padding(0, 0).Foreground(yellow)
	JiraUrl = Bold.Copy()
	KeyHelp = Regular.Copy().Padding(0, 1)
	Header  = Regular.Copy().Padding(0, 1).Border(lipgloss.RoundedBorder(), true)

	FocusedStyle       = Regular.Copy().Foreground(pink)
	TableHeaderStyle   = Regular.Copy().Foreground(blue)
	TableCellStyle     = Regular.Copy()
	TableSelectedStyle = Regular.Copy().Foreground(black).Background(blue)
)
