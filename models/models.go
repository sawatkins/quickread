package models

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username,omitempty"`
	CreatedOn string `json:"created_on"`
}

type PDFDocument struct {
	Id        string  	`json:"id"`
	Filename  string 	`json:"filename"`
	Url       string 	`json:"url"`
	CreatedOn string 	`json:"created_on"`
}
