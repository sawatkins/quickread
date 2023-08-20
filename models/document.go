package models

type Document struct {
	Id        int64  `json:"id"`
	Filename  string `json:"filename"`
	Url       string `json:"url"`
	CreatedOn string `json:"created_on"`
}
