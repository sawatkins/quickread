package models

// User model
type User struct {
	Id string `bson:"id"`
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Servers []Server `bson:"servers,omitempty" json:"servers,omitempty"`
	CreatedOn string `bson:"created_on" json:"created_on"`
}
