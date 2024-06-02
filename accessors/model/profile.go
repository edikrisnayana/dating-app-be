package model

type ProfileEntity struct {
	AccountId      string `bson:"_id,omitempty"`
	FirstName      string `bson:"firstName,omitempty"`
	LastName       string `bson:"lastName,omitempty"`
	Gender         int64  `bson:"gender,omitempty"`
	Location       string `bson:"location,omitempty"`
	ProfilePictUrl int64  `bson:"profilePictUrl,omitempty"`
}
