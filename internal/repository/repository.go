package repository

import (
	"context"
	"fmt"
	"github.com/obadoraibu/go-url-shortener/internal/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

type URLShortenerRepository struct {
	client *redis.Client
}

func (r URLShortenerRepository) CreateURL(url *model.URL) (*model.URL, error) {
	ctx := context.Background()
	ttl, err := time.ParseDuration(url.Expiry)
	if err != nil {
		return nil, err
	}
	err = r.client.Set(ctx, url.Id, url.Url, ttl).Err()
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (r URLShortenerRepository) FindURLByShort(url string) (string, error) {
	ctx := context.Background()
	value, err := r.client.Get(ctx, url).Result()
	if err != nil {
		return "", err
	}

	return value, nil
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
