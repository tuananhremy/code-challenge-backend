package dateutil

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

// Date format
const (
	// Ex: YYYY/MM/DD
	FormatYYYYMMDD = "2006/01/02"
	// Ex: YYYY/M/DD
	FormatYYYYMDD = "2006/1/02"
	// Ex: YYYY/MM/D
	FormatYYYYMMD = "2006/01/2"
	// Ex: YYYY/M/D
	FormatYYYYMD = "2006/1/2"
	// Ex: YYYY/MM/DD HH:mm
	FormatYYYYMMDDHHMM     = "2006/01/02 15:04"
	FormatYYYYMMDDHHMMDash = "2006-01-02 15:04"
	FormatYYYYMMDDDash     = "2006-01-02"
	FormatYYYYMM           = "2006/01"
	FormatYYYYMMDDNoSlash  = "20060102"
	FormatMMDD             = "01/02"
	FormatYYYYMMDDHHMMSSTZ = "2006/01/02 15:04:05Z07:00"
)

// TokyoTimeOffset define JST
const (
	TokyoTimeOffset = 9 * 60 * 60
)

var (
	serverTimeZone    string
	once              sync.Once
	acceptDateFormats = []string{FormatYYYYMMDD, FormatYYYYMDD, FormatYYYYMMD, FormatYYYYMD}
	// LocJP Japan location
	LocJP = time.FixedZone("Asia/Tokyo", TokyoTimeOffset)
)

// SetTimeZone set the timezone is used to parse date
func SetTimeZone(tz string) {
	once.Do(func() {
		serverTimeZone = tz
	})
}

// ToFormat format to date string
func ToFormat(t time.Time, format string) string {
	return t.Format(format)
}

// ServerTimeLocation return server location configured by environment variable,
// it throws panic error when loading zone failed
func ServerTimeLocation() *time.Location {
	loc, err := time.LoadLocation(serverTimeZone)
	if err != nil {
		panic(err)
	}
	return loc
}

func Now() time.Time {
	return time.Now().In(ServerTimeLocation())
}

// ToDate parse the date with format FormatYYYYMMDD from string with timezone
// based on the environment variable
func ToDate(dateStr string) (*time.Time, error) {
	slashMarkStr, isSuccess := addSlashMarkDate(dateStr)
	if isSuccess {
		dateStr = slashMarkStr
	}
	date, err := time.ParseInLocation(FormatYYYYMMDD, dateStr, ServerTimeLocation())
	if err != nil {
		return nil, err
	}
	return &date, nil
}

func ToDateWithAcceptFormat(dateStr string) (*time.Time, error) {
	if !strings.Contains(dateStr, "/") {
		slashMarkStr, isSuccess := addSlashMarkDate(dateStr)
		if isSuccess {
			dateStr = slashMarkStr
		}
	}

	for _, layout := range acceptDateFormats {
		date, err := time.ParseInLocation(layout, dateStr, ServerTimeLocation())
		if err == nil {
			return &date, nil
		}
	}

	return nil, errors.New("date format invalid")
}

// EndOfMonth return time.Time with last day of month
func EndOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m+1, 0, 0, 0, 0, 0, date.Location())
}

// FirstOfMonth return time.Time with first day of month
func FirstOfMonth(date time.Time) time.Time {
	y, m, _ := date.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, date.Location())
}

// EndOfNthMonth return time.Time with last day of next nth month
func EndOfNthMonth(date time.Time, nextMonth int) time.Time {
	// Convert date to first day of month to enhance the AddDate
	// https://golang.org/pkg/time/#Time.AddDate
	// Because with October 31 [AddDate(0, 1, 0)]-> December 1 (expect November 30)
	// Convert to first month October 1 [AddDate(0, 1, 0)]-> November 1 [EndOfMonth]-> November 30
	return EndOfMonth(FirstOfMonth(date).AddDate(0, nextMonth, 0))
}

// FirstOfNthMonth return time.Time with first day of next nth month
func FirstOfNthMonth(date time.Time, nextMonth int) time.Time {
	return FirstOfMonth(date).AddDate(0, nextMonth, 0)
}

// NextOfNthMonth return time.Time with next nth month
func NextOfNthMonth(date time.Time, nextMonth int) time.Time {
	_, _, day := date.Date()
	convertedDate := FirstOfMonth(date).AddDate(0, nextMonth, 0)
	_, _, lastDay := EndOfMonth(convertedDate).Date()

	if day > lastDay {
		return EndOfMonth(convertedDate)
	}

	return convertedDate.AddDate(0, 0, day-1)
}

// addSlashMarkDate convert the text with format YYYYMMDD to YYYY/MM/DD
// exp: 20210320 -> 2021/03/20
func addSlashMarkDate(dateStr string) (string, bool) {
	//nolint
	if len(dateStr) != 8 {
		return dateStr, false
	}
	var (
		yyyy = dateStr[:4]
		mm   = dateStr[4:6]
		dd   = dateStr[6:]
	)
	return fmt.Sprintf("%v/%v/%v", yyyy, mm, dd), true
}

// MonthDuration calculate number of month from start date to end date. The formula as below:
// - if month of end date is greater than month of start date,the month duration plus one
// - if day of end date is greater or equal to day of start date, the month duration plus one
func MonthDuration(startDate, endDate time.Time) (uint, error) {
	if startDate.After(endDate) {
		return 0, fmt.Errorf("invalid date period")
	}
	var (
		nMonths    uint
		_, _, sd   = startDate.Date()
		_, _, ed   = endDate.Date()
		bStartDate = FirstOfMonth(startDate)
		bEndDate   = FirstOfMonth(endDate)
	)
	for ; bStartDate.Before(bEndDate); bStartDate = bStartDate.AddDate(0, 1, 0) {
		nMonths++
	}
	if ed >= sd {
		nMonths++
	}
	return nMonths, nil
}
