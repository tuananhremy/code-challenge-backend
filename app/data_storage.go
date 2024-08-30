package app

import (
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	DataStorage struct {
		mysqlDB *gorm.DB
	}
)

func NewDataStorage(dns string) *DataStorage {
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DataStorage{
		mysqlDB: db,
	}
}

// Upsert user
func (ds *DataStorage) Upsert(user *User) error {
	err := ds.mysqlDB.Save(user).Error
	if err != nil {
		if strings.EqualFold("UNIQUE constraint failed: users.email", err.Error()) {
			return nil
		}
		return err
	}
	return nil
}

func (ds *DataStorage) QueryBooking(bookingId uint) (*Booking, error) {
	result := Booking{}
	err := ds.mysqlDB.Model(&Booking{}).Raw(`
    SELECT id, user_id, seat_id, start_time, end_time, checked_in 
    FROM bookings 
    WHERE id = ? 
    AND NOT EXISTS (
        SELECT 1 
        FROM bookings AS b 
        WHERE b.seat_id = bookings.seat_id 
        AND b.id != bookings.id 
        AND (
            (b.start_time < bookings.end_time AND b.end_time > bookings.start_time)
        )
    )
`, bookingId).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ds *DataStorage) ReseverBooking(booking *Booking) error {
	err := ds.mysqlDB.Exec("UPDATE bookings SET checked_in = true WHERE id = ?", booking.ID).Error
	if err != nil {
		return err
	}
	// Insert the check-in record into the database
	err = ds.mysqlDB.Exec("INSERT INTO check_ins (booking_id) VALUES (?)", booking.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorage) ReleaseBooking() error {
	// Find bookings that have not been checked in and are more than 10 minutes past the start time
	var bookings []Booking
	err := ds.mysqlDB.Raw(`
        SELECT id
        FROM bookings
        WHERE checked_in = false
        AND start_time < ?
    `, time.Now().Add(-10*time.Second)).Scan(&bookings).Error
	if err != nil {
		return err
	}

	// Release each booking
	for _, booking := range bookings {
		// Update the booking status
		err = ds.mysqlDB.Exec("DELETE from bookings WHERE id = ?", booking.ID).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// Create user
func (ds *DataStorage) Create(user *User) error {
	return ds.mysqlDB.Create(user).Error
}

func (ds *DataStorage) CreateBooking(booking *Booking) error {
	return ds.mysqlDB.Create(booking).Error
}

// get user by email
func (ds *DataStorage) GetUserByEmail(email string) (*User, error) {
	var user User
	err := ds.mysqlDB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (ds *DataStorage) GetSeatByNumber(number string) (*Seat, error) {
	var seat Seat
	err := ds.mysqlDB.Where("number = ?", number).First(&seat).Error
	if err != nil {
		return nil, err
	}
	return &seat, err
}

func (ds *DataStorage) Booking(id uint) (*User, error) {
	var user User
	err := ds.mysqlDB.Where("id = ?", id).First(&user).Error
	return &user, err
}

// Find seats
func (ds *DataStorage) FindSeats() ([]Seat, error) {
	var seats []Seat
	err := ds.mysqlDB.Find(&seats).Error
	return seats, err
}

// Find bookings that start_time and end_time of request is overlap with start_time and end_time of bookings
func (ds *DataStorage) FindOverlapBookings(startTime, endTime time.Time) ([]Booking, error) {
	var bookings []Booking
	err := ds.mysqlDB.Where("start_time <= ? AND end_time >= ?", endTime, startTime).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, err
}

// gorm transaction
func (ds *DataStorage) Transaction(fn func(ds *DataStorage) error) error {
	return ds.mysqlDB.Transaction(func(tx *gorm.DB) error {
		return fn(&DataStorage{mysqlDB: tx})
	})
}
