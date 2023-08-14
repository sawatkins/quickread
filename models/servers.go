package models

// Server model
type Server struct {
	Id string `bson:"id" json:"id"`
	UserId string `bson:"user_id" json:"user_id"`
	Url string `bson:"url,omitempty" json:"url,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	Players int `bson:"players,omitempty" json:"players,omitempty"`
	MaxPlayers int `bson:"max_players,omitempty" json:"max_players,omitempty"`
	StartingMap string `bson:"starting_map" json:"starting_map"`
	Region string `bson:"region" json:"region"`
	Status string `bson:"status" json:"status"` // stopped, starting, active, stopping
	Public bool `bson:"public" json:"public"`
	CreatedOn string `bson:"created_on" json:"created_on"`
}