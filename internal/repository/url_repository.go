package repository

import (
	"context"
	"github.com/obadoraibu/go-url-shortener/internal/model"
	"github.com/obadoraibu/go-url-shortener/internal/random"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

func (r URLShortenerRepository) CreateURL(url *model.URL) (*model.URL, error) {
	ctx := context.Background()
	ttl, err := time.ParseDuration(url.Expiry)
	if err != nil {
		return nil, err
	}

	existingURL, err := r.FindURLByShort(url.Id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if existingURL != "" {
		url.Id = random.Encode(rand.Uint64())
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
