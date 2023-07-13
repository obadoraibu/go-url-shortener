package model

type URL struct {
	Id     string `json:"short"`
	Url    string `json:"url" binding:"required,url"`
	Expiry string `json:"expiry" binding:"required"`
}
