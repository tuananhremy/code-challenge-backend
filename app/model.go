package app

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
}

type Seat struct {
	gorm.Model
	Number string `json:"number"`
}

type Booking struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	SeatID   uint      `json:"seat_id"`
	FromTime time.Time `json:"from_time"`
	ToTime   time.Time `json:"to_time"`
}
