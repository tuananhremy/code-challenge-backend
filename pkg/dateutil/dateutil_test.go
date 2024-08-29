package dateutil

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestEndOfMonth(t *testing.T) {
	type args struct {
		date time.Time
	}
	timeFormat := "2006-01-02"
	beginOfMonth, _ := time.Parse(timeFormat, "2020-12-01")
	testCase1, _ := time.Parse(timeFormat, "2020-12-31")
	middleOfMonth, _ := time.Parse(timeFormat, "2020-11-15")
	testCase2, _ := time.Parse(timeFormat, "2020-11-30")
	endOfYear, _ := time.Parse(timeFormat, "2020-12-31")
	testCase3, _ := time.Parse(timeFormat, "2020-12-31")
	specialCaseFeb2020, _ := time.Parse(timeFormat, "2020-02-11")
	testCase4, _ := time.Parse(timeFormat, "2020-02-29")
	beginOfMonthWithlocalTimeZone, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-01 6:00:00", time.Local)
	testCase5, _ := time.ParseInLocation(timeFormat, "2020-02-29", time.Local)
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test case 1: args begin of month",
			args: args{
				date: beginOfMonth,
			},
			want: testCase1,
		},
		{
			name: "test case 2: args middle of month",
			args: args{
				date: middleOfMonth,
			},
			want: testCase2,
		},
		{
			name: "test case 3: args end of year",
			args: args{
				date: endOfYear,
			},
			want: testCase3,
		},
		{
			name: "test case 4: special case Feb 2020",
			args: args{
				date: specialCaseFeb2020,
			},
			want: testCase4,
		},
		{
			name: "test case 5: with local time zone",
			args: args{
				date: beginOfMonthWithlocalTimeZone,
			},
			want: testCase5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfMonth(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOfMonth(t *testing.T) {
	type args struct {
		date time.Time
	}

	timeFormat := "2006-01-02"
	beginOfMonth, _ := time.Parse(timeFormat, "2020-12-01")
	testCase1, _ := time.Parse(timeFormat, "2020-12-01")
	middleOfMonth, _ := time.Parse(timeFormat, "2020-11-15")
	testCase2, _ := time.Parse(timeFormat, "2020-11-01")
	endOfYear, _ := time.Parse(timeFormat, "2020-12-31")
	testCase3, _ := time.Parse(timeFormat, "2020-12-01")
	loc := time.FixedZone("-07", 8*60*60)
	endOfMonthWithLocalTimeZone, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-29 6:00:00", loc)
	testCase4, _ := time.ParseInLocation(timeFormat, "2020-02-01", loc)

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test case 1: args begin of month",
			args: args{
				date: beginOfMonth,
			},
			want: testCase1,
		},
		{
			name: "test case 2: args middle of month",
			args: args{
				date: middleOfMonth,
			},
			want: testCase2,
		},
		{
			name: "test case 3: args end of year",
			args: args{
				date: endOfYear,
			},
			want: testCase3,
		},
		{
			name: "test case 4: args end of month with local time zone",
			args: args{
				date: endOfMonthWithLocalTimeZone,
			},
			want: testCase4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOfMonth(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfNthMonth(t *testing.T) {
	type args struct {
		date      time.Time
		nextMonth int
	}

	timeFormat := "2006-01-02"
	beginOfMonth, _ := time.Parse(timeFormat, "2020-12-01")
	testCase1, _ := time.Parse(timeFormat, "2021-01-31")
	middleOfMonth, _ := time.Parse(timeFormat, "2020-11-15")
	testCase2, _ := time.Parse(timeFormat, "2021-02-28")
	endOfYear, _ := time.Parse(timeFormat, "2020-12-31")
	testCase3, _ := time.Parse(timeFormat, "2021-12-31")
	OctoberOfYear, _ := time.Parse(timeFormat, "2020-10-31")
	testCase4, _ := time.Parse(timeFormat, "2022-10-31")

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test case 1: args begin of month, next 1 month",
			args: args{
				date:      beginOfMonth,
				nextMonth: 1,
			},
			want: testCase1,
		},
		{
			name: "test case 2: args middle of month, next 3 month",
			args: args{
				date:      middleOfMonth,
				nextMonth: 3,
			},
			want: testCase2,
		},
		{
			name: "test case 3: args end of year, next 12 month",
			args: args{
				date:      endOfYear,
				nextMonth: 12,
			},
			want: testCase3,
		},
		{
			name: "test case 3: args end of year, next 24 month",
			args: args{
				date:      OctoberOfYear,
				nextMonth: 24,
			},
			want: testCase4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfNthMonth(tt.args.date, tt.args.nextMonth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfNextNthMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOfNthMonth(t *testing.T) {
	type args struct {
		date      time.Time
		nextMonth int
	}

	timeFormat := "2006-01-02"
	beginOfMonth, _ := time.Parse(timeFormat, "2020-12-01")
	testCase1, _ := time.Parse(timeFormat, "2021-01-01")
	middleOfMonth, _ := time.Parse(timeFormat, "2020-11-15")
	testCase2, _ := time.Parse(timeFormat, "2021-01-01")
	endOfYear, _ := time.Parse(timeFormat, "2020-12-31")
	testCase3, _ := time.Parse(timeFormat, "2021-12-01")

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test case 1: args begin of month, next 1 month",
			args: args{
				date:      beginOfMonth,
				nextMonth: 1,
			},
			want: testCase1,
		},
		{
			name: "test case 2: args middle of month, next 2 month",
			args: args{
				date:      middleOfMonth,
				nextMonth: 2,
			},
			want: testCase2,
		},
		{
			name: "test case 3: args end of year, next 12 month",
			args: args{
				date:      endOfYear,
				nextMonth: 12,
			},
			want: testCase3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOfNthMonth(tt.args.date, tt.args.nextMonth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOfNthMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextOfNthMonth(t *testing.T) {
	type args struct {
		date      time.Time
		nextMonth int
	}

	timeFormat := "2006-01-02"
	beginOfMonth, _ := time.Parse(timeFormat, "2020-12-01")
	testCase1, _ := time.Parse(timeFormat, "2021-01-01")
	middleOfMonth, _ := time.Parse(timeFormat, "2020-11-15")
	testCase2, _ := time.Parse(timeFormat, "2021-01-15")
	endOfYear, _ := time.Parse(timeFormat, "2020-12-31")
	testCase3, _ := time.Parse(timeFormat, "2021-12-31")
	middleOfYearAndMonth, _ := time.Parse(timeFormat, "2021-02-15")
	testCase4, _ := time.Parse(timeFormat, "2021-05-15")
	endOfJan, _ := time.Parse(timeFormat, "2020-01-31")
	testCase5, _ := time.Parse(timeFormat, "2020-02-29")

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test case 1: args begin of month, next 1 month",
			args: args{
				date:      beginOfMonth,
				nextMonth: 1,
			},
			want: testCase1,
		},
		{
			name: "test case 2: args middle of month, next 2 month",
			args: args{
				date:      middleOfMonth,
				nextMonth: 2,
			},
			want: testCase2,
		},
		{
			name: "test case 3: args end of year, next 12 month",
			args: args{
				date:      endOfYear,
				nextMonth: 12,
			},
			want: testCase3,
		},
		{
			name: "test case 3: args middle of year and month , next 3 month",
			args: args{
				date:      middleOfYearAndMonth,
				nextMonth: 3,
			},
			want: testCase4,
		},
		{
			name: "test case 4: args middle of year and month , next 0 month",
			args: args{
				date:      middleOfYearAndMonth,
				nextMonth: 0,
			},
			want: middleOfYearAndMonth,
		},
		{
			name: "test case 5: args end of Jan , next 1 month",
			args: args{
				date:      endOfJan,
				nextMonth: 1,
			},
			want: testCase5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextOfNthMonth(tt.args.date, tt.args.nextMonth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextOfNthMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ToDate_success_1",
			args: args{
				dateStr: "2021/03/20",
			},
			wantErr: false,
		},
		{
			name: "ToDate_success_2",
			args: args{
				dateStr: "20210320",
			},
			wantErr: false,
		},
		{
			name: "ToDate_failed_1",
			args: args{
				dateStr: "202103201",
			},
			wantErr: true,
		},
		{
			name: "ToDate_failed_2",
			args: args{
				dateStr: "2021/03/201",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ToDate(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addSlashMarkDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "addSlashMarkDate_success",
			args: args{
				dateStr: "20210310",
			},
			want:  "2021/03/10",
			want1: true,
		},
		{
			name: "addSlashMarkDate_failed",
			args: args{
				dateStr: "2021031",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := addSlashMarkDate(tt.args.dateStr)
			if got1 != tt.want1 {
				t.Errorf("addSlashMarkDate() got1 = %v, want %v", got1, tt.want1)
				return
			}
			if got1 && got != tt.want {
				t.Errorf("addSlashMarkDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonthDuration(t *testing.T) {
	type args struct {
		startDate time.Time
		endDate   time.Time
	}

	dateAtFunc := func(year, month, day int) time.Time {
		m := time.Month(month)
		return time.Date(year, m, day, 0, 0, 0, 0, time.Local)
	}

	tests := []struct {
		name       string
		args       args
		wantMonths uint
		wantError  bool
	}{
		{
			name: "testcase#1: start date and end date are at the same month, but diff day ",
			args: args{
				startDate: dateAtFunc(2021, 03, 01),
				endDate:   dateAtFunc(2021, 03, 05),
			},
			wantMonths: 1,
			wantError:  false,
		},
		{
			name: "testcase#2: start date and end date are the same date",
			args: args{
				startDate: dateAtFunc(2021, 03, 05),
				endDate:   dateAtFunc(2021, 03, 05),
			},
			wantMonths: 1,
			wantError:  false,
		},
		{
			name: "testcase#3: start date and end date are the same day, but diff month",
			args: args{
				startDate: dateAtFunc(2021, 03, 05),
				endDate:   dateAtFunc(2021, 04, 05),
			},
			wantMonths: 2,
			wantError:  false,
		},
		{
			name: "testcase#4: end date is greater than start date",
			args: args{
				startDate: dateAtFunc(2021, 04, 05),
				endDate:   dateAtFunc(2021, 03, 05),
			},
			wantMonths: 0,
			wantError:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ms, err := MonthDuration(tt.args.startDate, tt.args.endDate)
			if err != nil && !tt.wantError {
				t.Errorf("dateutil.MonthDuration(), unexpect error: %v", err)
				return
			}
			if tt.wantError {
				assert.Error(t, err, "dateutil.MonthDuration(), expected an error, but got nothing")
				return
			}
			if diff := cmp.Diff(ms, tt.wantMonths); diff != "" {
				t.Errorf("dateutil.MonthDuration(), got diff: %s", diff)
			}
		})
	}
}

func TestSetTimeZone(t *testing.T) {
	t.Parallel()
	type args struct {
		tz string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set timezone to UTC",
			args: args{
				tz: "UTC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetTimeZone(tt.args.tz)
			assert.Equalf(t, serverTimeZone, tt.args.tz, "SetTimeZone(%v)", tt.args.tz)
		})
	}
}

func TestToFormat(t *testing.T) {
	t.Parallel()
	type args struct {
		t      time.Time
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1: format to date string",
			args: args{
				t:      time.Date(2021, 03, 05, 0, 0, 0, 0, time.UTC),
				format: "2006-01-02",
			},
			want: "2021-03-05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ToFormat(tt.args.t, tt.args.format), "ToFormat(%v, %v)", tt.args.t, tt.args.format)
		})
	}
}

func TestToDateWithAcceptFormat(t *testing.T) {
	t.Parallel()
	type args struct {
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr error
	}{
		{
			name: "test case 1: date string with slash mark",
			args: args{
				dateStr: "2021/03/05",
			},
			want: func() *time.Time { t := time.Date(2021, 03, 05, 0, 0, 0, 0, time.UTC); return &t }(),
		},
		{
			name: "test case 2: date string without slash mark",
			args: args{
				dateStr: "20210305",
			},
			want: func() *time.Time { t := time.Date(2021, 03, 05, 0, 0, 0, 0, time.UTC); return &t }(),
		},
		{
			name: "test case 3: date string with invalid format",
			args: args{
				dateStr: "2021-03-05",
			},
			wantErr: errors.New("date format invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToDateWithAcceptFormat(tt.args.dateStr)
			if tt.wantErr != nil {
				assert.Equalf(t, err.Error(), tt.wantErr.Error(), "ToDateWithAcceptFormat()")
				return
			}
			assert.Equalf(t, tt.want, got, "ToDateWithAcceptFormat(%v)", tt.args.dateStr)
		})
	}
}
