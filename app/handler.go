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

func (h *Handler) Login(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := &User{
		Email: request.Email,
		Name:  request.Name,
	}

	err := h.ds.Upsert(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func (h *Handler) ListAvailableSeats(c *gin.Context) {
	var request struct {
		FromTime time.Time `json:"from_time" binding:"required"`
		ToTime   time.Time `json:"to_time" binding:"required"`
	}
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	seats, err := h.ds.FindAvailableSeats(request.FromTime, request.ToTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, seats)
}

func (h *Handler) BookSeat(c *gin.Context) {
	var request struct {
		SeatNumber string `json:"seat_number"`
		UserEmail  string `json:"user_email"`
		FromTime   string `json:"from_time"`
		ToTime     string `json:"to_time"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.ds.GetUserByEmail(request.UserEmail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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
		UserID:    user.ID,
		SeatID:    seat.ID,
		StartTime: fromTime,
		EndTime:   toTime,
	}

	if err := h.ds.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book seat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Seat booked successfully"})
}
