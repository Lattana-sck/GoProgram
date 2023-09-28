package types

type ResponsData struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    Owner  `json:"owner"`
	HtmlURL  string `json:"html_url"`
	CloneURL string `json:"clone_url"`
}

type Owner struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	NodeID           string `json:"node_id"`
	AvatarURL        string `json:"avatar_url"`
	GravatarID       string `json:"gravatar_id"`
	URL              string `json:"url"`
	HtmlURL          string `json:"html_url"`
	FollowersURL     string `json:"followers_url"`
	FollowingURL     string `json:"following_url"`
	GistsURL         string `json:"gists_url"`
	StarredURL       string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURL string `json:"organizations_url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type             string `json:"type"`
	SiteAdmin        bool   `json:"site_admin"`
}