package main

import "time"

const (
	TypePush        = "PushEvent"
	TypeCreate      = "CreateEvent"
	TypeWatch       = "WatchEvent"
	TypeIssue       = "IssueEvent"
	TypePullRequest = "PullRequestEvent"
)

var Productive = []string{
	TypePush,
	TypeCreate,
	TypeIssue,
	TypePullRequest,
}

type TypedEvent struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		Id           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarId   string `json:"gravatar_id"`
		Url          string `json:"url"`
		AvatarUrl    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
	Public    bool      `json:"public"`
}

func (e *TypedEvent) IsProductive() bool {
	for _, p := range Productive {
		if p == e.Type {
			return true
		}
	}
	return false
}
