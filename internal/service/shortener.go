package service

import (
	"context"
	"github.com/obadoraibu/go-url-shortener/internal/model"
	"github.com/obadoraibu/go-url-shortener/internal/random"
	"math/rand"
)

func (s *URLShortenerService) Shorten(ctx context.Context, url *model.URL) (string, error) {
	if url.Id == "" {
		url.Id = random.Encode(rand.Uint64())
	}

	res, err := s.repo.CreateURL(url)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

func (s *URLShortenerService) Resolve(ctx context.Context, short string) (string, error) {
	url, err := s.repo.FindURLByShort(short)
	if err != nil {
		return "", err
	}
	return url, nil
}
