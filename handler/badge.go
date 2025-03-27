package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/konstfish/hits/models"
	"github.com/konstfish/hits/storage"
	"github.com/narqo/go-badge"
)

type BadgeHandler struct {
	store storage.CounterStore
}

func NewBadgeHandler(store storage.CounterStore) *BadgeHandler {
	return &BadgeHandler{
		store: store,
	}
}

func (h *BadgeHandler) HandleBadge(c *gin.Context) {
	var params models.BadgeParams

	if err := c.ShouldBindQuery(&params); err != nil {
		c.String(http.StatusBadRequest, "Invalid parameters: %v", err)
		return
	}

	if params.TitleBg == "" {
		params.TitleBg = "#555555"
	}
	if params.CountBg == "" {
		params.CountBg = "#007ec6"
	}
	if params.Title == "" {
		params.Title = "hits"
	}

	decodedURL, err := url.QueryUnescape(params.URL)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid URL encoding")
		return
	}

	todayCount, totalCount, err := h.store.IncrementCounters(c.Request.Context(), decodedURL)
	if err != nil {
		c.String(http.StatusInternalServerError, "", err)
		return
	}

	badgeText := fmt.Sprintf("%d / %d", todayCount, totalCount)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	if err := badge.Render(params.Title, badgeText, badge.Color(params.CountBg), c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "", err)
		return
	}
}
