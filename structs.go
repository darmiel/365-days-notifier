package main

import "time"

const (
	TypePush        = "PushEvent"
	TypeCreate      = "CreateEvent"
	TypeWatch       = "WatchEvent"
	TypeIssue       = "IssueEvent"
	TypePullRequest = "PullRequestEvent"
)

type CodeActivity interface {
	IsCode() bool
	GetID() string
	GetType() string
	GetCreatedAt() time.Time
	IsPublic() bool
}

type TypedEvent struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Public    bool      `json:"public"`
}

func (e *TypedEvent) IsCode() bool {
	return false
}
func (e *TypedEvent) GetID() string {
	return e.Id
}
func (e *TypedEvent) GetType() string {
	return e.Type
}
func (e *TypedEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *TypedEvent) IsPublic() bool {
	return e.Public
}

// ------------------------------------------------

type EventPush struct {
	*TypedEvent
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
	Payload struct {
		PushId       int    `json:"push_id"`
		Size         int    `json:"size"`
		DistinctSize int    `json:"distinct_size"`
		Ref          string `json:"ref"`
		Head         string `json:"head"`
		Before       string `json:"before"`
		Commits      []struct {
			Sha    string `json:"sha"`
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Message  string `json:"message"`
			Distinct bool   `json:"distinct"`
			Url      string `json:"url"`
		} `json:"commits"`
	} `json:"payload"`
}

func (e *EventPush) IsCode() bool {
	return true
}
func (e *EventPush) GetID() string {
	return e.Id
}
func (e *EventPush) GetType() string {
	return e.Type
}
func (e *EventPush) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *EventPush) IsPublic() bool {
	return e.Public
}

// ------------------------------------------------

type EventCreate struct {
	*TypedEvent
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
	Payload struct {
		Ref          interface{} `json:"ref"`
		RefType      string      `json:"ref_type"`
		MasterBranch string      `json:"master_branch"`
		Description  string      `json:"description"`
		PusherType   string      `json:"pusher_type"`
	} `json:"payload"`
	Public bool `json:"public"`
}

func (e *EventCreate) IsCode() bool {
	return true
}
func (e *EventCreate) GetID() string {
	return e.Id
}
func (e *EventCreate) GetType() string {
	return e.Type
}
func (e *EventCreate) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *EventCreate) IsPublic() bool {
	return e.Public
}

// ------------------------------------------------

type EventWatch struct {
	*TypedEvent
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
	Payload struct {
		Action string `json:"action"`
	} `json:"payload"`
	Public bool `json:"public"`
}

func (e *EventWatch) IsCode() bool {
	return false
}
func (e *EventWatch) GetID() string {
	return e.Id
}
func (e *EventWatch) GetType() string {
	return e.Type
}
func (e *EventWatch) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *EventWatch) IsPublic() bool {
	return e.Public
}

// ------------------------------------------------

type EventIssue struct {
	*TypedEvent
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
	Payload struct {
		Action string `json:"action"`
		Issue  struct {
			Url           string `json:"url"`
			RepositoryUrl string `json:"repository_url"`
			LabelsUrl     string `json:"labels_url"`
			CommentsUrl   string `json:"comments_url"`
			EventsUrl     string `json:"events_url"`
			HtmlUrl       string `json:"html_url"`
			Id            int    `json:"id"`
			NodeId        string `json:"node_id"`
			Number        int    `json:"number"`
			Title         string `json:"title"`
			User          struct {
				Login             string `json:"login"`
				Id                int    `json:"id"`
				NodeId            string `json:"node_id"`
				AvatarUrl         string `json:"avatar_url"`
				GravatarId        string `json:"gravatar_id"`
				Url               string `json:"url"`
				HtmlUrl           string `json:"html_url"`
				FollowersUrl      string `json:"followers_url"`
				FollowingUrl      string `json:"following_url"`
				GistsUrl          string `json:"gists_url"`
				StarredUrl        string `json:"starred_url"`
				SubscriptionsUrl  string `json:"subscriptions_url"`
				OrganizationsUrl  string `json:"organizations_url"`
				ReposUrl          string `json:"repos_url"`
				EventsUrl         string `json:"events_url"`
				ReceivedEventsUrl string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Labels            []interface{} `json:"labels"`
			State             string        `json:"state"`
			Locked            bool          `json:"locked"`
			Assignee          interface{}   `json:"assignee"`
			Assignees         []interface{} `json:"assignees"`
			Milestone         interface{}   `json:"milestone"`
			Comments          int           `json:"comments"`
			CreatedAt         time.Time     `json:"created_at"`
			UpdatedAt         time.Time     `json:"updated_at"`
			ClosedAt          interface{}   `json:"closed_at"`
			AuthorAssociation string        `json:"author_association"`
			ActiveLockReason  interface{}   `json:"active_lock_reason"`
			Body              interface{}   `json:"body"`
			Reactions         struct {
				Url        string `json:"url"`
				TotalCount int    `json:"total_count"`
				ThumbsUp   int    `json:"+1"`
				ThumbsDown int    `json:"-1"`
				Laugh      int    `json:"laugh"`
				Hooray     int    `json:"hooray"`
				Confused   int    `json:"confused"`
				Heart      int    `json:"heart"`
				Rocket     int    `json:"rocket"`
				Eyes       int    `json:"eyes"`
			} `json:"reactions"`
			TimelineUrl           string      `json:"timeline_url"`
			PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
		} `json:"issue"`
	} `json:"payload"`
	Public bool `json:"public"`
	Org    struct {
		Id         int    `json:"id"`
		Login      string `json:"login"`
		GravatarId string `json:"gravatar_id"`
		Url        string `json:"url"`
		AvatarUrl  string `json:"avatar_url"`
	} `json:"org"`
}

func (e *EventIssue) IsCode() bool {
	return true
}
func (e *EventIssue) GetID() string {
	return e.Id
}
func (e *EventIssue) GetType() string {
	return e.Type
}
func (e *EventIssue) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *EventIssue) IsPublic() bool {
	return e.Public
}

// ------------------------------------------------

type EventPullRequest struct {
	*TypedEvent
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
	Payload struct {
		Action      string `json:"action"`
		Number      int    `json:"number"`
		PullRequest struct {
			Url      string `json:"url"`
			Id       int    `json:"id"`
			NodeId   string `json:"node_id"`
			HtmlUrl  string `json:"html_url"`
			DiffUrl  string `json:"diff_url"`
			PatchUrl string `json:"patch_url"`
			IssueUrl string `json:"issue_url"`
			Number   int    `json:"number"`
			State    string `json:"state"`
			Locked   bool   `json:"locked"`
			Title    string `json:"title"`
			User     struct {
				Login             string `json:"login"`
				Id                int    `json:"id"`
				NodeId            string `json:"node_id"`
				AvatarUrl         string `json:"avatar_url"`
				GravatarId        string `json:"gravatar_id"`
				Url               string `json:"url"`
				HtmlUrl           string `json:"html_url"`
				FollowersUrl      string `json:"followers_url"`
				FollowingUrl      string `json:"following_url"`
				GistsUrl          string `json:"gists_url"`
				StarredUrl        string `json:"starred_url"`
				SubscriptionsUrl  string `json:"subscriptions_url"`
				OrganizationsUrl  string `json:"organizations_url"`
				ReposUrl          string `json:"repos_url"`
				EventsUrl         string `json:"events_url"`
				ReceivedEventsUrl string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Body               string        `json:"body"`
			CreatedAt          time.Time     `json:"created_at"`
			UpdatedAt          time.Time     `json:"updated_at"`
			ClosedAt           time.Time     `json:"closed_at"`
			MergedAt           interface{}   `json:"merged_at"`
			MergeCommitSha     interface{}   `json:"merge_commit_sha"`
			Assignee           interface{}   `json:"assignee"`
			Assignees          []interface{} `json:"assignees"`
			RequestedReviewers []interface{} `json:"requested_reviewers"`
			RequestedTeams     []interface{} `json:"requested_teams"`
			Labels             []interface{} `json:"labels"`
			Milestone          interface{}   `json:"milestone"`
			Draft              bool          `json:"draft"`
			CommitsUrl         string        `json:"commits_url"`
			ReviewCommentsUrl  string        `json:"review_comments_url"`
			ReviewCommentUrl   string        `json:"review_comment_url"`
			CommentsUrl        string        `json:"comments_url"`
			StatusesUrl        string        `json:"statuses_url"`
			Head               struct {
				Label string `json:"label"`
				Ref   string `json:"ref"`
				Sha   string `json:"sha"`
				User  struct {
					Login             string `json:"login"`
					Id                int    `json:"id"`
					NodeId            string `json:"node_id"`
					AvatarUrl         string `json:"avatar_url"`
					GravatarId        string `json:"gravatar_id"`
					Url               string `json:"url"`
					HtmlUrl           string `json:"html_url"`
					FollowersUrl      string `json:"followers_url"`
					FollowingUrl      string `json:"following_url"`
					GistsUrl          string `json:"gists_url"`
					StarredUrl        string `json:"starred_url"`
					SubscriptionsUrl  string `json:"subscriptions_url"`
					OrganizationsUrl  string `json:"organizations_url"`
					ReposUrl          string `json:"repos_url"`
					EventsUrl         string `json:"events_url"`
					ReceivedEventsUrl string `json:"received_events_url"`
					Type              string `json:"type"`
					SiteAdmin         bool   `json:"site_admin"`
				} `json:"user"`
				Repo struct {
					Id       int    `json:"id"`
					NodeId   string `json:"node_id"`
					Name     string `json:"name"`
					FullName string `json:"full_name"`
					Private  bool   `json:"private"`
					Owner    struct {
						Login             string `json:"login"`
						Id                int    `json:"id"`
						NodeId            string `json:"node_id"`
						AvatarUrl         string `json:"avatar_url"`
						GravatarId        string `json:"gravatar_id"`
						Url               string `json:"url"`
						HtmlUrl           string `json:"html_url"`
						FollowersUrl      string `json:"followers_url"`
						FollowingUrl      string `json:"following_url"`
						GistsUrl          string `json:"gists_url"`
						StarredUrl        string `json:"starred_url"`
						SubscriptionsUrl  string `json:"subscriptions_url"`
						OrganizationsUrl  string `json:"organizations_url"`
						ReposUrl          string `json:"repos_url"`
						EventsUrl         string `json:"events_url"`
						ReceivedEventsUrl string `json:"received_events_url"`
						Type              string `json:"type"`
						SiteAdmin         bool   `json:"site_admin"`
					} `json:"owner"`
					HtmlUrl          string      `json:"html_url"`
					Description      string      `json:"description"`
					Fork             bool        `json:"fork"`
					Url              string      `json:"url"`
					ForksUrl         string      `json:"forks_url"`
					KeysUrl          string      `json:"keys_url"`
					CollaboratorsUrl string      `json:"collaborators_url"`
					TeamsUrl         string      `json:"teams_url"`
					HooksUrl         string      `json:"hooks_url"`
					IssueEventsUrl   string      `json:"issue_events_url"`
					EventsUrl        string      `json:"events_url"`
					AssigneesUrl     string      `json:"assignees_url"`
					BranchesUrl      string      `json:"branches_url"`
					TagsUrl          string      `json:"tags_url"`
					BlobsUrl         string      `json:"blobs_url"`
					GitTagsUrl       string      `json:"git_tags_url"`
					GitRefsUrl       string      `json:"git_refs_url"`
					TreesUrl         string      `json:"trees_url"`
					StatusesUrl      string      `json:"statuses_url"`
					LanguagesUrl     string      `json:"languages_url"`
					StargazersUrl    string      `json:"stargazers_url"`
					ContributorsUrl  string      `json:"contributors_url"`
					SubscribersUrl   string      `json:"subscribers_url"`
					SubscriptionUrl  string      `json:"subscription_url"`
					CommitsUrl       string      `json:"commits_url"`
					GitCommitsUrl    string      `json:"git_commits_url"`
					CommentsUrl      string      `json:"comments_url"`
					IssueCommentUrl  string      `json:"issue_comment_url"`
					ContentsUrl      string      `json:"contents_url"`
					CompareUrl       string      `json:"compare_url"`
					MergesUrl        string      `json:"merges_url"`
					ArchiveUrl       string      `json:"archive_url"`
					DownloadsUrl     string      `json:"downloads_url"`
					IssuesUrl        string      `json:"issues_url"`
					PullsUrl         string      `json:"pulls_url"`
					MilestonesUrl    string      `json:"milestones_url"`
					NotificationsUrl string      `json:"notifications_url"`
					LabelsUrl        string      `json:"labels_url"`
					ReleasesUrl      string      `json:"releases_url"`
					DeploymentsUrl   string      `json:"deployments_url"`
					CreatedAt        time.Time   `json:"created_at"`
					UpdatedAt        time.Time   `json:"updated_at"`
					PushedAt         time.Time   `json:"pushed_at"`
					GitUrl           string      `json:"git_url"`
					SshUrl           string      `json:"ssh_url"`
					CloneUrl         string      `json:"clone_url"`
					SvnUrl           string      `json:"svn_url"`
					Homepage         string      `json:"homepage"`
					Size             int         `json:"size"`
					StargazersCount  int         `json:"stargazers_count"`
					WatchersCount    int         `json:"watchers_count"`
					Language         string      `json:"language"`
					HasIssues        bool        `json:"has_issues"`
					HasProjects      bool        `json:"has_projects"`
					HasDownloads     bool        `json:"has_downloads"`
					HasWiki          bool        `json:"has_wiki"`
					HasPages         bool        `json:"has_pages"`
					ForksCount       int         `json:"forks_count"`
					MirrorUrl        interface{} `json:"mirror_url"`
					Archived         bool        `json:"archived"`
					Disabled         bool        `json:"disabled"`
					OpenIssuesCount  int         `json:"open_issues_count"`
					License          struct {
						Key    string `json:"key"`
						Name   string `json:"name"`
						SpdxId string `json:"spdx_id"`
						Url    string `json:"url"`
						NodeId string `json:"node_id"`
					} `json:"license"`
					AllowForking  bool   `json:"allow_forking"`
					Visibility    string `json:"visibility"`
					Forks         int    `json:"forks"`
					OpenIssues    int    `json:"open_issues"`
					Watchers      int    `json:"watchers"`
					DefaultBranch string `json:"default_branch"`
				} `json:"repo"`
			} `json:"head"`
			Base struct {
				Label string `json:"label"`
				Ref   string `json:"ref"`
				Sha   string `json:"sha"`
				User  struct {
					Login             string `json:"login"`
					Id                int    `json:"id"`
					NodeId            string `json:"node_id"`
					AvatarUrl         string `json:"avatar_url"`
					GravatarId        string `json:"gravatar_id"`
					Url               string `json:"url"`
					HtmlUrl           string `json:"html_url"`
					FollowersUrl      string `json:"followers_url"`
					FollowingUrl      string `json:"following_url"`
					GistsUrl          string `json:"gists_url"`
					StarredUrl        string `json:"starred_url"`
					SubscriptionsUrl  string `json:"subscriptions_url"`
					OrganizationsUrl  string `json:"organizations_url"`
					ReposUrl          string `json:"repos_url"`
					EventsUrl         string `json:"events_url"`
					ReceivedEventsUrl string `json:"received_events_url"`
					Type              string `json:"type"`
					SiteAdmin         bool   `json:"site_admin"`
				} `json:"user"`
				Repo struct {
					Id       int    `json:"id"`
					NodeId   string `json:"node_id"`
					Name     string `json:"name"`
					FullName string `json:"full_name"`
					Private  bool   `json:"private"`
					Owner    struct {
						Login             string `json:"login"`
						Id                int    `json:"id"`
						NodeId            string `json:"node_id"`
						AvatarUrl         string `json:"avatar_url"`
						GravatarId        string `json:"gravatar_id"`
						Url               string `json:"url"`
						HtmlUrl           string `json:"html_url"`
						FollowersUrl      string `json:"followers_url"`
						FollowingUrl      string `json:"following_url"`
						GistsUrl          string `json:"gists_url"`
						StarredUrl        string `json:"starred_url"`
						SubscriptionsUrl  string `json:"subscriptions_url"`
						OrganizationsUrl  string `json:"organizations_url"`
						ReposUrl          string `json:"repos_url"`
						EventsUrl         string `json:"events_url"`
						ReceivedEventsUrl string `json:"received_events_url"`
						Type              string `json:"type"`
						SiteAdmin         bool   `json:"site_admin"`
					} `json:"owner"`
					HtmlUrl          string      `json:"html_url"`
					Description      string      `json:"description"`
					Fork             bool        `json:"fork"`
					Url              string      `json:"url"`
					ForksUrl         string      `json:"forks_url"`
					KeysUrl          string      `json:"keys_url"`
					CollaboratorsUrl string      `json:"collaborators_url"`
					TeamsUrl         string      `json:"teams_url"`
					HooksUrl         string      `json:"hooks_url"`
					IssueEventsUrl   string      `json:"issue_events_url"`
					EventsUrl        string      `json:"events_url"`
					AssigneesUrl     string      `json:"assignees_url"`
					BranchesUrl      string      `json:"branches_url"`
					TagsUrl          string      `json:"tags_url"`
					BlobsUrl         string      `json:"blobs_url"`
					GitTagsUrl       string      `json:"git_tags_url"`
					GitRefsUrl       string      `json:"git_refs_url"`
					TreesUrl         string      `json:"trees_url"`
					StatusesUrl      string      `json:"statuses_url"`
					LanguagesUrl     string      `json:"languages_url"`
					StargazersUrl    string      `json:"stargazers_url"`
					ContributorsUrl  string      `json:"contributors_url"`
					SubscribersUrl   string      `json:"subscribers_url"`
					SubscriptionUrl  string      `json:"subscription_url"`
					CommitsUrl       string      `json:"commits_url"`
					GitCommitsUrl    string      `json:"git_commits_url"`
					CommentsUrl      string      `json:"comments_url"`
					IssueCommentUrl  string      `json:"issue_comment_url"`
					ContentsUrl      string      `json:"contents_url"`
					CompareUrl       string      `json:"compare_url"`
					MergesUrl        string      `json:"merges_url"`
					ArchiveUrl       string      `json:"archive_url"`
					DownloadsUrl     string      `json:"downloads_url"`
					IssuesUrl        string      `json:"issues_url"`
					PullsUrl         string      `json:"pulls_url"`
					MilestonesUrl    string      `json:"milestones_url"`
					NotificationsUrl string      `json:"notifications_url"`
					LabelsUrl        string      `json:"labels_url"`
					ReleasesUrl      string      `json:"releases_url"`
					DeploymentsUrl   string      `json:"deployments_url"`
					CreatedAt        time.Time   `json:"created_at"`
					UpdatedAt        time.Time   `json:"updated_at"`
					PushedAt         time.Time   `json:"pushed_at"`
					GitUrl           string      `json:"git_url"`
					SshUrl           string      `json:"ssh_url"`
					CloneUrl         string      `json:"clone_url"`
					SvnUrl           string      `json:"svn_url"`
					Homepage         string      `json:"homepage"`
					Size             int         `json:"size"`
					StargazersCount  int         `json:"stargazers_count"`
					WatchersCount    int         `json:"watchers_count"`
					Language         string      `json:"language"`
					HasIssues        bool        `json:"has_issues"`
					HasProjects      bool        `json:"has_projects"`
					HasDownloads     bool        `json:"has_downloads"`
					HasWiki          bool        `json:"has_wiki"`
					HasPages         bool        `json:"has_pages"`
					ForksCount       int         `json:"forks_count"`
					MirrorUrl        interface{} `json:"mirror_url"`
					Archived         bool        `json:"archived"`
					Disabled         bool        `json:"disabled"`
					OpenIssuesCount  int         `json:"open_issues_count"`
					License          struct {
						Key    string `json:"key"`
						Name   string `json:"name"`
						SpdxId string `json:"spdx_id"`
						Url    string `json:"url"`
						NodeId string `json:"node_id"`
					} `json:"license"`
					AllowForking  bool   `json:"allow_forking"`
					Visibility    string `json:"visibility"`
					Forks         int    `json:"forks"`
					OpenIssues    int    `json:"open_issues"`
					Watchers      int    `json:"watchers"`
					DefaultBranch string `json:"default_branch"`
				} `json:"repo"`
			} `json:"base"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Html struct {
					Href string `json:"href"`
				} `json:"html"`
				Issue struct {
					Href string `json:"href"`
				} `json:"issue"`
				Comments struct {
					Href string `json:"href"`
				} `json:"comments"`
				ReviewComments struct {
					Href string `json:"href"`
				} `json:"review_comments"`
				ReviewComment struct {
					Href string `json:"href"`
				} `json:"review_comment"`
				Commits struct {
					Href string `json:"href"`
				} `json:"commits"`
				Statuses struct {
					Href string `json:"href"`
				} `json:"statuses"`
			} `json:"_links"`
			AuthorAssociation   string      `json:"author_association"`
			AutoMerge           interface{} `json:"auto_merge"`
			ActiveLockReason    interface{} `json:"active_lock_reason"`
			Merged              bool        `json:"merged"`
			Mergeable           bool        `json:"mergeable"`
			Rebaseable          bool        `json:"rebaseable"`
			MergeableState      string      `json:"mergeable_state"`
			MergedBy            interface{} `json:"merged_by"`
			Comments            int         `json:"comments"`
			ReviewComments      int         `json:"review_comments"`
			MaintainerCanModify bool        `json:"maintainer_can_modify"`
			Commits             int         `json:"commits"`
			Additions           int         `json:"additions"`
			Deletions           int         `json:"deletions"`
			ChangedFiles        int         `json:"changed_files"`
		} `json:"pull_request"`
	} `json:"payload"`
	Public bool `json:"public"`
	Org    struct {
		Id         int    `json:"id"`
		Login      string `json:"login"`
		GravatarId string `json:"gravatar_id"`
		Url        string `json:"url"`
		AvatarUrl  string `json:"avatar_url"`
	} `json:"org"`
}

func (e *EventPullRequest) IsCode() bool {
	return true
}
func (e *EventPullRequest) GetID() string {
	return e.Id
}
func (e *EventPullRequest) GetType() string {
	return e.Type
}
func (e *EventPullRequest) GetCreatedAt() time.Time {
	return e.CreatedAt
}
func (e *EventPullRequest) IsPublic() bool {
	return e.Public
}
