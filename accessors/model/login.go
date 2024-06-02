package model

type LoginEntity struct {
	LoginId   string `bson:"_id,omitempty"`
	AccountId string `bson:"accountId,omitempty"`
	LoginType string `bson:"loginType,omitempty"`
}
