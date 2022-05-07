package models

type Projects []Project

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Links       []struct {
		Url  string `json:"url"`
		Type string `json:"type"`
	} `json:"links"`
}
