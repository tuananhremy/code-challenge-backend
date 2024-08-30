package app

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email"`
}

type Seat struct {
	gorm.Model
	Number   string `json:"number"`
	Booked   bool   `json:"booked"`
	BookedBy uint   `json:"booked_by"`
}

type Booking struct {
	gorm.Model
	UserID uint `json:"user_id"`
	SeatID uint `json:"seat_id"`
}
