package keymap

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up        key.Binding
	Down      key.Binding
	Compact   key.Binding
	Back      key.Binding
	Exit      key.Binding
	OpenIssue key.Binding
	Left      key.Binding
	Right     key.Binding
	Refresh   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Exit, k.Compact, k.Left, k.Right,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Exit, k.Compact, k.Back, k.OpenIssue},
		{k.Up, k.Down, k.Left, k.Right, k.Refresh},
	}
}

var KeyMap = keyMap{
	Back: key.NewBinding(
		key.WithKeys("esc", "backspace"),
		key.WithHelp("esc", "back"),
	),
	OpenIssue: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Open Issue"),
	),
	Exit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "exit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Compact: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "compact"),
	),
	Left: key.NewBinding(
		key.WithKeys("ctrl+left"),
		key.WithHelp("ctrl+←", "left query qte"),
		key.WithDisabled(),
	),
	Right: key.NewBinding(
		key.WithKeys("ctrl+right"),
		key.WithHelp("ctrl+→", "right query"),
		key.WithDisabled(),
	),
	Refresh: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "refresh"),
	),
}
