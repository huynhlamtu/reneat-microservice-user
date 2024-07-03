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

type UserRole struct {
	Uuid string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
}

type User struct {
	ClientUuid string     `json:"client_uuid, omitempty" bson:"client_uuid"`
	Uuid       string     `json:"uuid,omitempty" bson:"uuid"`
	FirstName  string     `json:"first_name" bson:"first_name"`
	LastName   string     `json:"last_name" bson:"last_name"`
	Email      string     `json:"email,omitempty" bson:"email"`
	Username   string     `json:"username,omitempty" bson:"username"`
	Password   string     `json:"password,omitempty" bson:"password"`
	IsBlock    int        `json:"is_block" bson:"is_block"`
	IsVerified int        `json:"is_verified" bson:"is_verified"`
	Tried      int        `json:"tried" bson:"tried"`
	Token      string     `json:"token" bson:"token"`
	Roles      []UserRole `json:"roles" bson:"roles"`
	IsAdmin    int        `json:"is_admin" bson:"is_admin"`
	IsActive   int        `json:"is_active" bson:"is_active"`
	IsDelete   int        `json:"is_delete" bson:"is_delete"`
	CreatedAt  time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" bson:"updated_at"`
	CreatedBy  *string    `json:"created_by" bson:"created_by"`
	UpdatedBy  *string    `json:"updated_by" bson:"updated_by"`
}

func (u *User) Model() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("users")
}

func (u *User) Find(conditions map[string]interface{}, opts ...*options.FindOptions) ([]*User, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	cursor, err := coll.Find(context.TODO(), conditions, opts...)
	if err != nil {
		return nil, err
	}

	var users []*User
	for cursor.Next(context.TODO()) {
		var elem User
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

func (u *User) Pagination(ctx context.Context, conditions map[string]interface{}, modelOptions ...ModelOption) ([]*User, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	modelOpt := ModelOption{}
	findOptions := modelOpt.GetOption(modelOptions)
	cursor, err := coll.Find(context.TODO(), conditions, findOptions)
	if err != nil {
		return nil, err
	}

	var users []*User
	for cursor.Next(context.TODO()) {
		var elem User
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

func (u *User) Distinct(conditions map[string]interface{}, fieldName string, opts ...*options.DistinctOptions) ([]interface{}, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE

	values, err := coll.Distinct(context.TODO(), fieldName, conditions, opts...)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func (u *User) FindOne(conditions map[string]interface{}) (*User, error) {
	coll := u.Model()

	conditions["is_delete"] = constant.UNDELETE
	err := coll.FindOne(context.TODO(), conditions).Decode(&u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Insert() (interface{}, error) {
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

func (u *User) InsertMany(Users []interface{}) ([]interface{}, error) {
	coll := u.Model()

	resp, err := coll.InsertMany(context.TODO(), Users)
	if err != nil {
		return nil, err
	}

	return resp.InsertedIDs, nil
}

func (u *User) Update() (int64, error) {
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

func (u *User) UpdateByCondition(condition map[string]interface{}, data map[string]interface{}) (int64, error) {
	coll := u.Model()

	resp, err := coll.UpdateOne(context.TODO(), condition, data)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *User) UpdateMany(conditions map[string]interface{}, updateData map[string]interface{}) (int64, error) {
	coll := u.Model()
	resp, err := coll.UpdateMany(context.TODO(), conditions, updateData)
	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, nil
}

func (u *User) Count(ctx context.Context, condition map[string]interface{}) (int64, error) {
	coll := u.Model()

	condition["is_delete"] = constant.UNDELETE

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}
