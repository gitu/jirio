package main

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/gitu/jirio/internal/jiracache"
	"github.com/gitu/jirio/internal/tui/components/app"
	"github.com/spf13/viper"
	"os"
)

func main() {
	initialize()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	cache, err := jiracache.NewCache(ctx, jiracache.JiraConfig{
		Url:                  viper.GetString("jira.url"),
		Token:                viper.GetString("jira.token"),
		Projects:             viper.GetStringSlice("jira.projects"),
		MaxResults:           viper.GetInt("jira.maxResults"),
		UseMockData:          viper.GetBool("jira.fake"),
		InitialUpdatedSince:  viper.GetString("jira.initialUpdatedSince"),
		RefreshInterval:      viper.GetDuration("jira.refreshInterval"),
		RefreshIntervalSince: viper.GetString("jira.refreshIntervalSince"),
	})
	if err != nil {
		fmt.Printf("could not start cache loader: %s\n", err)
		os.Exit(1)
	}
	model := app.InitialModel(cache)
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

func initialize() {
	log.SetReportCaller(true)
	viper.SetEnvPrefix("JIRIO")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/jirio/")
	viper.AddConfigPath("$HOME/.jirio")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn("could not read config file", err)
	}
}
