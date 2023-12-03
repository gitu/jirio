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
	"strings"
)

func main() {
	initialize()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	viper.SetDefault("jira.refreshInterval", "10m")
	viper.SetDefault("jira.refreshIntervalSince", "0s")
	viper.SetDefault("jira.maxResults", 1000)

	queries := make(map[string]jiracache.JiraQuery)

	if len(viper.GetStringSlice("jira.projects")) > 0 || viper.GetBool("jira.fake") {
		queries["all"] = jiracache.JiraQuery{
			Name:        "Selected Projects",
			Jql:         fmt.Sprintf("project in (%s)", toQuotedList(viper.GetStringSlice("jira.projects"))),
			Incremental: true,
		}
		if viper.GetBool("jira.fake") {
			queries["all_TTP"] = jiracache.JiraQuery{
				Name:        "AI",
				Jql:         "project in (AI)",
				Incremental: true,
			}
			queries["all_AVQ"] = jiracache.JiraQuery{
				Name:        "AVQ",
				Jql:         "project in (AVQ)",
				Incremental: false,
			}
		}
	}

	queryIds := viper.GetStringMap("jira.queries")
	for k, _ := range queryIds {
		queries[k] = jiracache.JiraQuery{
			Name:        viper.GetString("jira.queries." + k + ".name"),
			Jql:         viper.GetString("jira.queries." + k + ".jql"),
			Incremental: viper.GetBool("jira.queries." + k + ".incremental"),
		}
	}

	cache, err := jiracache.NewCache(ctx, jiracache.JiraConfig{
		Url:                viper.GetString("jira.url"),
		Token:              viper.GetString("jira.token"),
		MaxResults:         viper.GetInt("jira.maxResults"),
		UseMockData:        viper.GetBool("jira.fake"),
		RefreshInterval:    viper.GetDuration("jira.refreshInterval"),
		InitialIncremental: viper.GetString("jira.initialIncremental"),
		Queries:            queries,
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

// toQuotedList returns a string of projects in the format "PROJECT1","PROJECT2",...
func toQuotedList(projects []string) string {
	var projectsString string
	for _, project := range projects {
		projectsString += "\"" + project + "\","
	}
	return strings.TrimRight(projectsString, ",")
}
