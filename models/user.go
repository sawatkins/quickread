package models

// User model
type User struct {
	Id        string `json:"id"`
	Username  string `json:"username,omitempty" json:"username,omitempty"`
	CreatedOn string `json:"created_on" json:"created_on"`
}
