package models

type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	Action string `json:"action"`
}

type Org struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	GravatarID string `json:"gravatar_id"`
	Url        string `json:"url"`
	AvatarURL  string `json:"avatar_url"`
}

type Event struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	Public    bool    `json:"public"`
	CreatedAt string  `json:"created_at"`
	Org       Org     `json:"org"`
}
