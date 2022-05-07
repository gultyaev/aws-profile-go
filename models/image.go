package models

type PresignImage struct {
	Url       string `json:"url"`
	UploadUrl string `json:"uploadUrl"`
}
