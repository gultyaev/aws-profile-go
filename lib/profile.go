package lib

type Profile struct {
	Name       string           `json:"name"`
	Location   string           `json:"location"`
	Contacts   []ProfileContact `json:"contacts"`
	ProfileImg string           `json:"profileImg"`
}

type ProfileContact struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
