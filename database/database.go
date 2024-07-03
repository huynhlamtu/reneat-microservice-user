package database

import (
	"context"
	"fmt"
	"reneat-microservice-user/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Init() (*mongo.Database, error) {
	if db == nil {
		cfg := config.GetConfig()
		database := cfg.GetString("database.db_name")
		host := cfg.GetString("database.host")
		port := cfg.GetString("database.port")
		user := cfg.GetString("database.username")
		password := cfg.GetString("database.password")
		ssl := cfg.GetBool("database.ssl")

		var uri string
		if user != "" && password != "" {
			if ssl {
				uri = fmt.Sprintf(`mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&readPreference=secondaryPreferred`, user, password, host)
			} else {
				uri = fmt.Sprintf(`mongodb://%s:%s@%s:%s/?authMechanism=SCRAM-SHA-256`, user, password, host, port)
			}
		} else {
			if ssl {
				uri = fmt.Sprintf(`mongodb+srv://%s/?retryWrites=true&w=majority&readPreference=secondaryPreferred`, host)
			} else {
				uri = fmt.Sprintf(`mongodb://%s:%s`, host, port)
			}
		}

		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()

		optionsClient := options.Client()
		optionsClient.ApplyURI(uri)

		client, err := mongo.Connect(ctx, optionsClient)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("MongoDB ping failed: %w", err)
		} else {
			println("Database connected successfully")
		}

		db = client.Database(database)
	}

	return db, nil
}

func GetInstance() *mongo.Database {
	return db
}
