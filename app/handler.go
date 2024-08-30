package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Handler struct {
		ds *DataStorage
	}
)

func NewHandler(ds *DataStorage) *Handler {
	return &Handler{
		ds: ds,
	}
}

func (h *Handler) ListAvailableSeats(c *gin.Context) {
	seats, err := h.ds.FindSeats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, seats)
}
