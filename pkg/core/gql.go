package core

import (
	"context"
	"time"

	"github.com/shurcooL/githubv4"
)

type Variables struct {
	User          User
	Home          Home
	PostsAfter    string
	Projects      []Project
	ProjectsAfter string
	Dark          bool
	Description   string
}

type Home struct {
	Body    string
	Content string
}

type User struct {
	Name      string
	AvatarURL string
}

type Project struct {
	Name        string
	Description string
	URL         string
	Stars       int
	Forks       int
	UpdatedAt   time.Time
	Topics      []Topic
}

type Topic struct {
	Name string
	URL  string
}

type Repositories struct {
	Nodes []struct {
		Name             string
		URL              string
		Description      string
		StargazerCount   int
		ForkCount        int
		PushedAt         time.Time
		RepositoryTopics struct {
			Nodes []struct {
				Topic struct {
					Name string
				}
				URL string
			}
		} `graphql:"repositoryTopics(first: $topics)"`
	}
	PageInfo struct {
		EndCursor   string
		HasNextPage bool
	}
}

func (s *Server) FetchData(ctx context.Context) (*Variables, error) {
	var query struct {
		User struct {
			Login      string
			AvatarURL  string
			Repository struct {
				Description string
				Object      struct {
					Tree struct {
						Entries []struct {
							Name   string
							Object struct {
								Blob struct {
									Text string
								} `graphql:"... on Blob"`
							}
						}
					} `graphql:"... on Tree"`
				} `graphql:"object(expression: $expression)"`
			} `graphql:"repository(name: $user)"`
			Repositories Repositories `graphql:"repositories(first: $repositories, isFork: false, privacy: PUBLIC, orderBy: {field: PUSHED_AT, direction: DESC})"`
		} `graphql:"user(login: $user)"`
	}

	variables := map[string]interface{}{
		"user":         githubv4.String(s.Config.GitHub.User),
		"repositories": githubv4.Int(10),
		"topics":       githubv4.Int(10),
		"expression":   githubv4.String("HEAD:"),
	}

	if err := s.GitHub.Query(ctx, &query, variables); err != nil {
		return nil, err
	}

	user := User{
		Name:      query.User.Login,
		AvatarURL: query.User.AvatarURL,
	}

	var home Home
	for _, entry := range query.User.Repository.Object.Tree.Entries {
		if entry.Name == "README.md" {
			home.Body = entry.Object.Blob.Text
		}
	}

	projectsAfter := query.User.Repositories.PageInfo.EndCursor
	if !query.User.Repositories.PageInfo.HasNextPage {
		projectsAfter = ""
	}

	vars := &Variables{
		User:          user,
		Home:          home,
		Projects:      make([]Project, 0, len(query.User.Repositories.Nodes)),
		ProjectsAfter: projectsAfter,
		Dark:          true,
		Description:   query.User.Repository.Description,
	}

	for _, node := range query.User.Repositories.Nodes {
		topics := make([]Topic, 0, len(node.RepositoryTopics.Nodes))

		for _, tNode := range node.RepositoryTopics.Nodes {
			topics = append(topics, Topic{
				Name: tNode.Topic.Name,
				URL:  tNode.URL,
			})
		}

		vars.Projects = append(vars.Projects, Project{
			Name:        node.Name,
			Description: node.Description,
			URL:         node.URL,
			Stars:       node.StargazerCount,
			Forks:       node.ForkCount,
			UpdatedAt:   node.PushedAt,
			Topics:      topics,
		})
	}

	return vars, nil
}
