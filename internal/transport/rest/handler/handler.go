package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/obadoraibu/go-url-shortener/internal/model"
	"golang.org/x/net/context"
)

type URLShortenerService interface {
	Shorten(ctx context.Context, url *model.URL) (string, error)
	Resolve(ctx context.Context, url string) (string, error)
}

type URLShortenerHandler struct {
	service URLShortenerService
}

func NewHandler(service URLShortenerService) *URLShortenerHandler {
	return &URLShortenerHandler{service: service}
}

func (h *URLShortenerHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/shorten", h.Shorten)
	router.GET("/:short", h.Redirect)

	return router
}
