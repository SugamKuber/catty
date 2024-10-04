package database

import (
	"context"
	"fmt"
	"os"

	// "github.com/go-redis/redis/v8"
	"server/internal/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase() types.Database {

	mongoURI := os.Getenv("DB_URI")
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Println("Failed to connect to MongoDB: %v", err)
	}

	if err := mongoClient.Ping(context.TODO(), nil); err != nil {
		fmt.Println("Failed to ping MongoDB: %v", err)
	}

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     os.Getenv("REDIS_DB_ADDR"),
	// 	Password: os.Getenv("REDIS_DB_PASS"),
	// 	DB:       1,
	// })

	// if err := redisClient.Ping(context.TODO()).Err(); err != nil {
	// 	fmt.Println("Failed to connect to Redis: %v", err)
	// }

	return types.Database{
		MongoClient: mongoClient,
		// RedisClient: redisClient,
	}
}
