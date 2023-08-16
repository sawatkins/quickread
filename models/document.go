package models

type Document struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	User      string `json:"username,omitempty" json:"username,omitempty"`
	CreatedOn string `json:"created_on" json:"created_on"`
}
