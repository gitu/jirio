package style

import "github.com/charmbracelet/lipgloss"

const (
	black     = lipgloss.Color("#000000")
	blue      = lipgloss.Color("6")
	greenblue = lipgloss.Color("#00A095")
	pink      = lipgloss.Color("#E760FC")
	darkred   = lipgloss.Color("#FF0000")
	darkgreen = lipgloss.Color("#00FF00")
	grey      = lipgloss.Color("#737373")
	red       = lipgloss.Color("#FF5353")
	yellow    = lipgloss.Color("#DBBD70")
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
