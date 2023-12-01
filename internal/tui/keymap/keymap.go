package keymap

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Compact key.Binding
	Back    key.Binding
	Exec    key.Binding
	Exit    key.Binding
}

var KeyMap = keyMap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Exec: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "exec"),
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
}
