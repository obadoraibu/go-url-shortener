package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/obadoraibu/go-url-shortener/internal/model"
	"net/http"
	"time"
)

func (h *URLShortenerHandler) Shorten(c *gin.Context) {
	input := &model.URL{}

	if err := c.ShouldBindJSON(&input); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	if !expiryFormat(input.Expiry) {
		sendErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	short, err := h.service.Shorten(c, input)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"short": short,
	})
}

func (h *URLShortenerHandler) Redirect(c *gin.Context) {
	short := c.Param("short")

	url, err := h.service.Resolve(c, short)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}

func expiryFormat(s string) bool {
	_, err := time.ParseDuration(s)
	return err == nil
}
