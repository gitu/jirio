package jirio

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/gitu/jirio/internal/tui/style"
	"strings"
)

type Page int8

const (
	Unset Page = iota
	Search
)

func getShortHelp(bindings []key.Binding) string {
	var output string
	for _, km := range bindings {
		output += style.KeyHelpKey.Render(km.Help().Key) + " " + style.KeyHelpDescription.Render(km.Help().Desc) + "  "
	}
	output = strings.TrimSpace(output)
	return output
}

func changeKeyHelp(k *key.Binding, h string) {
	k.SetHelp(k.Help().Key, h)
}

func GetPage() {

}
