package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	timeFormat = "2006-01-02 15:04"
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

func (h *Handler) BookSeat(c *gin.Context) {
	var request struct {
		SeatNumber string `json:"seat_number"`
		UserID     uint   `json:"user_id"`
		FromTime   string `json:"from_time"`
		ToTime     string `json:"to_time"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	seat, err := h.ds.GetSeatByNumber(request.SeatNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seat not found"})
		return
	}

	fromTime, err := time.Parse(timeFormat, request.FromTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid from_time format"})
		return
	}

	toTime, err := time.Parse(timeFormat, request.ToTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid to_time format"})
		return
	}

	booking := &Booking{
		UserID:   request.UserID,
		SeatID:   seat.ID,
		FromTime: fromTime,
		ToTime:   toTime,
	}

	if err := h.ds.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book seat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat booked successfully"})
}
