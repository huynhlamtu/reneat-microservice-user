package models

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reneat-microservice-user/constant"
	"reneat-microservice-user/database"
	timeHelper "reneat-microservice-user/helpers/time"
	"time"
)

type UserRegister struct {
	ClientUuid  string    `json:"client_uuid, omitempty" bson:"client_uuid"`
	Uuid        string    `json:"uuid,omitempty" bson:"uuid"`
	Email       string    `json:"email,omitempty" bson:"email"`
	Otp         string    `json:"otp" bson:"otp"`
	ExpiredTime time.Time `json:"expired_time" bson:"expired_time"`
	IsBlock     int       `json:"is_block" bson:"is_block"`
	IsActive    int       `json:"is_active" bson:"is_active"`
	IsDelete    int       `json:"is_delete" bson:"is_delete"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (u *UserRegister) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("user_registers")
}

func (u *UserRegister) Find(conditions map[string]interface{}, opts ...*options.FindOptions) ([]*UserRegister, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	cursor, err := coll.Find(context.TODO(), conditions, opts...)
	if err != nil {
		return nil, err
	}

	var users []*UserRegister
	for cursor.Next(context.TODO()) {
		var elem UserRegister
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		users = append(users, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return users, nil
}

func (u *UserRegister) Pagination(ctx context.Context, conditions map[string]interface{}, modelOptions ...ModelOption) ([]*UserRegister, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	modelOpt := ModelOption{}
	findOptions := modelOpt.GetOption(modelOptions)
	cursor, err := coll.Find(context.TODO(), conditions, findOptions)
	if err != nil {
		return nil, err
	}

	var users []*UserRegister
	for cursor.Next(context.TODO()) {
		var elem UserRegister
		err := cursor.Decode(&elem)
		if err != nil {
			log.Println("[Decode] PopularCuisine:", err)
			log.Println("-> #", elem.Uuid)
			continue
		}

		users = append(users, &elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	_ = cursor.Close(context.TODO())

	return users, nil
}

func (u *UserRegister) Distinct(conditions map[string]interface{}, fieldName string, opts ...*options.DistinctOptions) ([]interface{}, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	values, err := coll.Distinct(context.TODO(), fieldName, conditions, opts...)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (u *UserRegister) FindOne(conditions map[string]interface{}) (*UserRegister, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *UserRegister) Insert() (interface{}, error) {
	coll := u.Model()

	if u.Uuid == "" {
		u.Uuid = uuid.New().String()
	}

	if u.CreatedAt.IsZero() {
		u.CreatedAt = timeHelper.NowUTC()
	}

	resp, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (u *UserRegister) InsertMany(Users []interface{}) ([]interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertMany(context.TODO(), Users)
	if err != nil {
		return nil, err
	}

	return resp.InsertedIDs, nil
}

func (u *UserRegister) Update() (int64, error) {
	coll := u.Model()

	condition := make(map[string]interface{})
	condition["uuid"] = u.Uuid

	u.UpdatedAt = timeHelper.NowUTC()
	updateStr := make(map[string]interface{})
	updateStr["$set"] = u

	resp, err := coll.UpdateOne(context.TODO(), condition, updateStr)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *UserRegister) UpdateByCondition(condition map[string]interface{}, data map[string]interface{}) (int64, error) {
	coll := u.Model()

	resp, err := coll.UpdateOne(context.TODO(), condition, data)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *UserRegister) UpdateMany(conditions map[string]interface{}, updateData map[string]interface{}) (int64, error) {
	coll := u.Model()
	resp, err := coll.UpdateMany(context.TODO(), conditions, updateData)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *UserRegister) Count(ctx context.Context, condition map[string]interface{}) (int64, error) {
	coll := u.Model()

	condition["is_delete"] = constant.UNDELETE

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}
