package dev

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
	"log"
	"os"
)

// dev
func Debug(msg string) {
	if viper.GetBool("debug") {
		f, err := tea.LogToFile("jirio.log", "")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		log.Printf("%q", msg)
		defer f.Close()
	}
}
