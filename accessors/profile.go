package accessors

import (
	"datingAppBE/accessors/model"
	"datingAppBE/accessors/mongodb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileEntityAccessor interface {
	init(string) ProfileEntityAccessor
	Insert(model.ProfileEntity) error
}

type profileEntityMongoAccessor struct {
	accessor *mongodb.Accessor
}

func CreateProfileEntityAccessor(hostname string) ProfileEntityAccessor {
	var accessor ProfileEntityAccessor = new(profileEntityMongoAccessor)
	return accessor.init(hostname)
}

func (accessor *profileEntityMongoAccessor) init(hostname string) ProfileEntityAccessor {
	accessor.accessor = mongodb.CreateAccessor(hostname, "schooldb", "profile")
	return accessor
}

func (accessor *profileEntityMongoAccessor) Insert(profileEntity model.ProfileEntity) error {
	result, err := accessor.accessor.Collection.InsertOne(accessor.accessor.Context, &profileEntity)
	if err != nil {
		return err
	}

	if result.InsertedID == primitive.NilObjectID {
		return fmt.Errorf("failed when insert profileEntity")
	}

	return nil
}
