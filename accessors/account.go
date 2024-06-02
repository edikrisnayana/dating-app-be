package accessors

import (
	"datingAppBE/accessors/model"
	"datingAppBE/accessors/mongodb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountEntityAccessor interface {
	init(string) AccountEntityAccessor
	Insert(model.AccountEntity) error
}

type accountEntityMongoAccessor struct {
	accessor *mongodb.Accessor
}

func CreateAccountEntityAccessor(hostname string) AccountEntityAccessor {
	var accessor AccountEntityAccessor = new(accountEntityMongoAccessor)
	return accessor.init(hostname)
}

func (accessor *accountEntityMongoAccessor) init(hostname string) AccountEntityAccessor {
	accessor.accessor = mongodb.CreateAccessor(hostname, "datingAppDB", "account")
	return accessor
}

func (accessor *accountEntityMongoAccessor) Insert(accountEntity model.AccountEntity) error {
	result, err := accessor.accessor.Collection.InsertOne(accessor.accessor.Context, &accountEntity)
	if err != nil {
		return err
	}

	if result.InsertedID == primitive.NilObjectID {
		return fmt.Errorf("failed when insert accountEntity")
	}

	return nil
}
