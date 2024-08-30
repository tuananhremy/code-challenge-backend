package app

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

type Seat struct {
	gorm.Model
	Number string `json:"number"`
}

type Booking struct {
	ID        int       `json:"id"`
	UserID    uint      `json:"user_id"`
	SeatID    uint      `json:"seat_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CheckedIn bool      `json:"checked_in"`
	CreatedAt time.Time `json:"created_at"`
}
