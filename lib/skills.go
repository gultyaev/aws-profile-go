package lib

type Skill struct {
	Title string `json:"title"`
}

type Skills []Skill

type SkillsCollection struct {
	Collection string
	Value      Skills
}
