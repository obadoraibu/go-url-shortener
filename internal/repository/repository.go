package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type URLShortenerRepository struct {
	client *redis.Client
}

func (r URLShortenerRepository) Close() error {
	err := r.client.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewRepository() (*URLShortenerRepository, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPortStr := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	err = client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	logrus.Print("connected to redis, %s", err)

	return &URLShortenerRepository{
		client: client,
	}, nil
}
