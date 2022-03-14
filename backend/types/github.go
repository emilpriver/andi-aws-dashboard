package types

type GithubLoginCallbackResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type GithubUser struct {
	Login      string `json:"login"`
	ID         int64  `json:"id"`
	NodeId     string `json:"node_id"`
	AvatarUrl  string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	URL        string `json:"url"`
	SiteAdmin  bool   `json:"site_admin"`
	Name       string `json:"name"`
	Company    string `json:"company"`
	Blog       string `json:"blog"`
	Location   string `json:"location"`
	Email      string `json:"email"`
	Hireable   string `json:"hireable"`
	Bio        string `json:"bio"`
}
