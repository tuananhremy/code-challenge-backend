package app

import (
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

// Create user
func (ds *DataStorage) Create(user *User) error {
	return ds.mysqlDB.Create(user).Error
}

func (ds *DataStorage) CreateBooking(booking *Booking) error {
	return ds.mysqlDB.Create(booking).Error
}

func (ds *DataStorage) GetSeatByNumber(number string) (*Seat, error) {
	var seat Seat
	err := ds.mysqlDB.Where("number = ?", number).First(&seat).Error
	if err != nil {
		return nil, err
	}
	return &seat, err
}

// Find seats
func (ds *DataStorage) FindSeats() ([]Seat, error) {
	var seats []Seat
	err := ds.mysqlDB.Find(&seats).Error
	return seats, err
}

// gorm transaction
func (ds *DataStorage) Transaction(fn func(ds *DataStorage) error) error {
	return ds.mysqlDB.Transaction(func(tx *gorm.DB) error {
		return fn(&DataStorage{mysqlDB: tx})
	})
}
