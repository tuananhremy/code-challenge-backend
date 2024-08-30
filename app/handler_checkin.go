package app

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CheckinService struct {
	ds        *DataStorage
	jwtSecret string
}

func NewCheckInService(ds *DataStorage, jwtSecret string) *CheckinService {
	return &CheckinService{
		ds:        ds,
		jwtSecret: jwtSecret,
	}
}

func (h *CheckinService) CheckIn(c *gin.Context) {
	var checkIn struct {
		SeatID    uint `json:"seat_id"`
		UserID    uint `json:"user_id"`
		BookingID uint `json:"booking_id"`
	}
	if err := c.ShouldBindJSON(&checkIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the booking exists and is not already checked in

	booking, err := h.ds.QueryBooking(checkIn.BookingID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking"})
		}
		return
	}
	if booking.SeatID != checkIn.SeatID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking does not match seat"})
		return
	}

	if booking.UserID != checkIn.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking does not match user"})
		return
	}

	if booking.CheckedIn {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking already checked in"})
		return
	}

	if time.Now().Sub(booking.StartTime) > 10*time.Minute {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check-in time exceeded"})
		return
	}

	// Check in the booking
	err = h.ds.ReseverBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check in"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Check-in successful"})
}

func (h *CheckinService) ReleaseBooking() {
	log.Printf("start release booking")
	for {
		err := h.ds.ReleaseBooking()
		if err != nil {
			// Log the error
			log.Printf("release fail")
		}
		time.Sleep(1)
	}
}
