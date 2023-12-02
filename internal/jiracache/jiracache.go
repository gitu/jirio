package jiracache

import (
	"context"
	"fmt"
	jira "github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/charmbracelet/log"
	"strings"
	"time"
)

type JiraCache interface {
	Search(search string) (*bleve.SearchResult, error)
	GetIssue(key string) (jira.Issue, error)
	Url() string
}

type jiraCache struct {
	issues  map[string]jira.Issue
	client  *jira.Client
	index   bleve.Index
	mapping *mapping.IndexMappingImpl
	config  JiraConfig
}

type JiraConfig struct {
	Url                  string
	Token                string
	Projects             []string
	MaxResults           int
	UseMockData          bool
	InitialUpdatedSince  string
	RefreshInterval      time.Duration
	RefreshIntervalSince string
}

func NewCache(ctx context.Context, config JiraConfig) (JiraCache, error) {
	j := &jiraCache{
		issues: make(map[string]jira.Issue),
		config: config,
	}
	var err error

	j.mapping = bleve.NewIndexMapping()
	j.index, err = bleve.NewMemOnly(j.mapping)
	if err != nil {
		return nil, err
	}

	if j.config.UseMockData {
		issues := buildFakeIssues()
		err = j.addIssues(issues)
		return j, err
	}

	tp := jira.BearerAuthTransport{
		Token: config.Token,
	}

	client, err := jira.NewClient(config.Url, tp.Client())
	if err != nil {
		return nil, err
	}
	j.client = client

	err = j.FetchIssues(ctx, config.InitialUpdatedSince)
	if err != nil {
		return nil, err
	}

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(config.RefreshInterval):
				err := j.FetchIssues(ctx, config.RefreshIntervalSince)
				if err != nil {
					log.Fatal("could not fetch issues", err)
				}
			}
		}
	}(ctx)

	return j, nil
}

func (j *jiraCache) Url() string {
	return j.config.Url
}

func (j *jiraCache) GetIssue(key string) (jira.Issue, error) {
	return j.issues[key], nil
}

func (j *jiraCache) Search(search string) (*bleve.SearchResult, error) {
	query := bleve.NewMatchQuery(search)
	query.Fuzziness = 2

	searchRequestOption := bleve.NewSearchRequestOptions(
		query, 40, 0, false)
	result, err := j.index.Search(searchRequestOption)
	if err != nil {
		log.Error("could not search", err)
		return result, err
	}
	return result, err
}

func (j *jiraCache) FetchIssues(ctx context.Context, since string) error {
	if j.client == nil {
		log.Error("client not initialized")
		return fmt.Errorf("client not initialized")
	}
	query := fmt.Sprintf(`project in (%v) AND updated >= %s ORDER BY updated DESC`,
		toQuotedList(j.config.Projects), since)
	issues, _, err := j.client.Issue.Search(ctx, query,
		&jira.SearchOptions{MaxResults: j.config.MaxResults})

	if err != nil {
		log.Error("could not fetch issues", err)
		return err
	}

	err = j.addIssues(issues)
	if err != nil {
		log.Error("could not add issues to cache", err)
		return err
	}
	return nil
}

func (j *jiraCache) addIssues(issues []jira.Issue) error {
	for _, issue := range issues {
		j.issues[issue.Key] = issue
		err := j.index.Index(issue.Key, issue)
		if err != nil {
			return err
		}
	}
	return nil
}

// toQuotedList returns a string of projects in the format "PROJECT1","PROJECT2",...
func toQuotedList(projects []string) string {
	var projectsString string
	for _, project := range projects {
		projectsString += "\"" + project + "\","
	}
	return strings.TrimRight(projectsString, ",")
}
