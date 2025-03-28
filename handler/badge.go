package handler

import (
	"context"
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

func setBadgeHeaders(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
}

func (h *BadgeHandler) processAndRenderBadge(c *gin.Context, counterOp func(ctx context.Context, url string) (int64, int64, error)) {
	var params models.BadgeParams

	if err := c.ShouldBindQuery(&params); err != nil {
		c.String(http.StatusBadRequest, "invalid parameters: %v", err)
		return
	}

	params.SetDefaults()

	decodedURL, err := url.QueryUnescape(params.URL)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid URL encoding")
		return
	}

	todayCount, totalCount, _ := counterOp(c.Request.Context(), decodedURL)
	if err != nil {
		fmt.Printf("failed to process request: %v\n", err)
		c.String(http.StatusInternalServerError, "unable to process request, counter issue")
		return
	}

	badgeText := fmt.Sprintf("%d / %d", todayCount, totalCount)

	setBadgeHeaders(c)

	if err := badge.Render(params.Title, badgeText, badge.Color(params.CountBg), c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "unable to render badge")
		return
	}
}

func (h *BadgeHandler) HandleIncrBadge(c *gin.Context) {
	h.processAndRenderBadge(c, h.store.IncrementCounters)
}

func (h *BadgeHandler) HandleShowBadge(c *gin.Context) {
	h.processAndRenderBadge(c, h.store.ShowCounters)
}
