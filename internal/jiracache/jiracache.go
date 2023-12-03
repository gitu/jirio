package jiracache

import (
	"context"
	"fmt"
	jira "github.com/andygrunwald/go-jira/v2/onpremise"
	"github.com/blevesearch/bleve/v2"
	"github.com/charmbracelet/log"
	"time"
)

type JiraCache interface {
	Search(query, search string) (*bleve.SearchResult, error)
	GetIssue(key string) (jira.Issue, error)
	Url() string
	RefreshIssues() error
	Queries() map[string]string
}

type jiraQuery struct {
	key   string
	query JiraQuery
	index bleve.Index
}

type jiraCache struct {
	issues      map[string]jira.Issue
	client      *jira.Client
	config      JiraConfig
	queries     map[string]jiraQuery
	lastRefresh time.Time
}

func (j *jiraCache) Queries() map[string]string {
	queries := make(map[string]string)
	for k, v := range j.queries {
		queries[k] = v.query.Name
	}
	return queries
}

type JiraQuery struct {
	Name        string
	Jql         string
	Incremental bool
}

type JiraConfig struct {
	Url                string
	Token              string
	MaxResults         int
	UseMockData        bool
	InitialIncremental string
	RefreshInterval    time.Duration
	Queries            map[string]JiraQuery
}

func NewCache(ctx context.Context, config JiraConfig) (JiraCache, error) {
	j := &jiraCache{
		issues:  make(map[string]jira.Issue),
		config:  config,
		queries: make(map[string]jiraQuery),
	}
	var err error

	for k, v := range config.Queries {
		index, err := bleve.NewMemOnly(bleve.NewIndexMapping())
		if err != nil {
			log.Error("could not create index", err)
			return nil, err
		}
		j.queries[k] = jiraQuery{
			query: v,
			index: index,
		}
	}

	if j.config.UseMockData {
		issues := buildFakeIssues()
		err = j.addIssues("all", issues)
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

	go func(ctx context.Context) {
		err = j.RefreshIssues()
		if err != nil {
			log.Fatal("could not fetch issues", err)
		}
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(config.RefreshInterval):
				err := j.RefreshIssues()
				if err != nil {
					log.Fatal("could not fetch issues", err)
				}
			}
		}
	}(ctx)

	return j, nil
}

func (j *jiraCache) RefreshIssues() error {
	if j.client == nil {
		log.Error("client not initialized")
		return fmt.Errorf("client not initialized")
	}
	for _, query := range j.queries {
		issues, _, err := j.client.Issue.Search(context.Background(), query.query.Jql,
			&jira.SearchOptions{MaxResults: j.config.MaxResults, Expand: "changelog"})
		if err != nil {
			log.Error("could not fetch issues", err)
			return err
		}
		err = j.addIssues(query.query.Jql, issues)
		if err != nil {
			log.Error("could not add issues to cache", err)
			return err
		}
	}
	j.lastRefresh = time.Now()
	return nil
}

func (j *jiraCache) Url() string {
	return j.config.Url
}

func (j *jiraCache) GetIssue(key string) (jira.Issue, error) {
	return j.issues[key], nil
}

func (j *jiraCache) Search(query, search string) (*bleve.SearchResult, error) {
	bq := bleve.NewMatchQuery(search)
	bq.Fuzziness = 2

	searchRequestOption := bleve.NewSearchRequestOptions(bq, 100, 0, false)
	result, err := j.queries[query].index.Search(searchRequestOption)
	if err != nil {
		log.Error("could not search", err)
		return result, err
	}
	return result, err
}

func (j *jiraCache) FetchIssues(ctx context.Context, query jiraQuery) error {
	if j.client == nil {
		log.Error("client not initialized")
		return fmt.Errorf("client not initialized")
	}

	q := query.query.Jql

	if query.query.Incremental {
		fetchSince := j.config.InitialIncremental
		updatedSince := j.lastRefresh
		if !updatedSince.IsZero() {
			fetchSince = fmt.Sprintf("-%.0fh", time.Now().Sub(updatedSince).Hours()+2)
		}
		q = fmt.Sprintf("%s AND updated >= %s ORDER BY updated DESC", q, fetchSince)
		return nil
	}

	issues, _, err := j.client.Issue.Search(ctx, q, &jira.SearchOptions{MaxResults: j.config.MaxResults})

	if err != nil {
		log.Error("could not fetch issues", err)
		return err
	}

	err = j.addIssues(query.key, issues)
	if err != nil {
		log.Error("could not add issues to cache", err)
		return err
	}
	return nil
}

func (j *jiraCache) addIssues(query string, issues []jira.Issue) error {
	for _, issue := range issues {
		j.issues[issue.Key] = issue
		err := j.queries[query].index.Index(issue.Key, issue)
		if err != nil {
			return err
		}
	}
	return nil
}
