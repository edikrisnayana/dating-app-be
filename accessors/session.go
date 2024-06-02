package accessors

import (
	"datingAppBE/accessors/model"
	"datingAppBE/accessors/mongodb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionEntityAccessor interface {
	init(string) SessionEntityAccessor
	Insert(model.SessionEntity) error
}

type sessionEntityMongoAccessor struct {
	accessor *mongodb.Accessor
}

func CreateSessionEntityAccessor(hostname string) SessionEntityAccessor {
	var accessor SessionEntityAccessor = new(sessionEntityMongoAccessor)
	return accessor.init(hostname)
}

func (accessor *sessionEntityMongoAccessor) init(hostname string) SessionEntityAccessor {
	accessor.accessor = mongodb.CreateAccessor(hostname, "datingAppDB", "session")
	return accessor
}

func (accessor *sessionEntityMongoAccessor) Insert(sessionEntity model.SessionEntity) error {
	result, err := accessor.accessor.Collection.InsertOne(accessor.accessor.Context, &sessionEntity)
	if err != nil {
		return err
	}

	if result.InsertedID == primitive.NilObjectID {
		return fmt.Errorf("failed when insert sessionEntity")
	}

	return nil
}
