package accessors

import (
	"datingAppBE/accessors/model"
	"datingAppBE/accessors/mongodb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginEntityAccessor interface {
	init(string) LoginEntityAccessor
	Insert(model.LoginEntity) error
}

type loginEntityMongoAccessor struct {
	accessor *mongodb.Accessor
}

func CreateLoginEntityAccessor(hostname string) LoginEntityAccessor {
	var accessor LoginEntityAccessor = new(loginEntityMongoAccessor)
	return accessor.init(hostname)
}

func (accessor *loginEntityMongoAccessor) init(hostname string) LoginEntityAccessor {
	accessor.accessor = mongodb.CreateAccessor(hostname, "schooldb", "login")
	return accessor
}

func (accessor *loginEntityMongoAccessor) Insert(loginEntity model.LoginEntity) error {
	result, err := accessor.accessor.Collection.InsertOne(accessor.accessor.Context, &loginEntity)
	if err != nil {
		return err
	}

	if result.InsertedID == primitive.NilObjectID {
		return fmt.Errorf("failed when insert loginEntity")
	}

	return nil
}
