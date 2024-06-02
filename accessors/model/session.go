package model

type SessionEntity struct {
	SessionId             string `bson:"_id,omitempty"`
	AccountId             string `bson:"accountId,omitempty"`
	SessionType           string `bson:"sessionType,omitempty"`
	SessionExpirationTime int64  `bson:"sessionExpirationTime,omitempty"`
}
