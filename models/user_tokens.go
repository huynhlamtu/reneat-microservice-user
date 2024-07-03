package models

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"reneat-microservice-user/constant"
	"reneat-microservice-user/database"
	timeHelper "reneat-microservice-user/helpers/time"
	"time"
)

type UserToken struct {
	Uuid      string    `json:"uuid" bson:"uuid,omitempty"`
	UserUuid  string    `json:"user_uuid,omitempty" bson:"user_uuid,omitempty"`
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	IsActive  int       `json:"is_active" bson:"is_active"`
	IsDelete  int       `json:"is_delete" bson:"is_delete"`
	ExpiredAt time.Time `json:"expired_at" bson:"expired_at,omitempty"`
	CreatedBy string    `json:"created_by" bson:"created_by"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (userToken *UserToken) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("user_tokens")
}

func (userToken *UserToken) FindOne(conditions map[string]interface{}) (*UserToken, error) {
	coll := userToken.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&userToken)
	if err != nil {
		return nil, err
	}

	return userToken, nil
}

func (userToken *UserToken) Insert() (interface{}, error) {
	coll := userToken.Model()

	if userToken.Uuid == "" {
		userToken.Uuid = uuid.New().String()
	}

	if userToken.CreatedAt.IsZero() {
		userToken.CreatedAt = timeHelper.NowUTC()
	}

	resp, err := coll.InsertOne(context.TODO(), userToken)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (userToken *UserToken) Update() (int64, error) {
	coll := userToken.Model()

	condition := make(map[string]interface{})
	condition["uuid"] = userToken.Uuid

	userToken.UpdatedAt = timeHelper.NowUTC()
	updateStr := make(map[string]interface{})
	updateStr["$set"] = userToken

	resp, err := coll.UpdateOne(context.TODO(), condition, updateStr)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}
