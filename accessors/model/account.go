package model

type AccountEntity struct {
	AccountId string `bson:"_id,omitempty"`
	Password  string `bson:"password,omitempty"`
}
