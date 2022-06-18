package models

type Projects []Project

type Project struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	DescriptionShort string   `json:"descriptionShort"`
	Images           []string `json:"images"`
	Links            []struct {
		Url  string `json:"url"`
		Type string `json:"type"`
	} `json:"links"`
}
