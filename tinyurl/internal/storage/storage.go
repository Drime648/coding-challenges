package storage

import (
	"context"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)


type storage struct {
	redisClient *redis.Client
}


var (
	storeService = &storage{}
	ctx = context.Background()
)

const cacheDuration = 24 * time.Hour


func InitStorage() error {
	redisClient := redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})

	err := testConnection(redisClient)
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis initialized succesfully")

	storeService.redisClient = redisClient
	
	return nil
}

func testConnection(redisClient *redis.Client) error {
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Error with connecting to Redis: %v\n", err)
	}
	return nil
}


func SaveUrl(originalUrl string, shortenedUrl string) error {

	err := storeService.redisClient.Set(ctx, shortenedUrl, originalUrl, cacheDuration).Err() //key is shortened, value is original
	if err != nil {
		return fmt.Errorf("Failed to save shortened url %s for original url %s | Error: %v", shortenedUrl, originalUrl, err)
	}
	return nil
}

func GetUrl(shortenedUrl string) (string, error) {

	originalUrl, err := storeService.redisClient.Get(ctx, shortenedUrl).Result()
	if err != nil {
		return "", fmt.Errorf("Failed to retrieve original URL for the shortened url %s | Error: %v", shortenedUrl, err)
	}
	return originalUrl, nil
}



