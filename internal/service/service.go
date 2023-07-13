package service

import (
	"github.com/obadoraibu/go-url-shortener/internal/model"
)

type URLShortenerService struct {
	repo URLShortenerRepository
}

type URLShortenerRepository interface {
	CreateURL(url *model.URL) (*model.URL, error)
	FindURLByShort(url string) (string, error)
}

func NewService(repo URLShortenerRepository) *URLShortenerService {
	return &URLShortenerService{
		repo: repo,
	}
}
