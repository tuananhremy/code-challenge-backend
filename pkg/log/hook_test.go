package log

import (
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewNopHook(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want NopHook
	}{
		{
			name: "success",
			want: NopHook{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNopHook(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNopHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNopHook_Levels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		h    NopHook
		want []log.Level
	}{
		{
			name: "success",
			h:    NopHook{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NopHook{}
			if got := h.Levels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NopHook.Levels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNopHook_Fire(t *testing.T) {
	t.Parallel()

	type args struct {
		entry *log.Entry
	}
	tests := []struct {
		name    string
		h       NopHook
		args    args
		wantErr bool
	}{
		{
			name: "success",
			h:    NopHook{},
			args: args{
				entry: &log.Entry{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NopHook{}
			if err := h.Fire(tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("NopHook.Fire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type MockHook struct {
	entries []*log.Entry
}

func (h MockHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *MockHook) Fire(entry *log.Entry) error {
	h.entries = append(h.entries, entry)
	return nil
}

func TestAddHook(t *testing.T) {
	t.Parallel()

	var (
		h = &MockHook{}
	)

	AddHook(h)
	for _, level := range log.AllLevels {
		assert.Contains(t, logger.Hooks[level], h)
	}
}
