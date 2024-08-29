package log

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLevel(t *testing.T) {
	type args struct {
		lvl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "debug",
			args: args{
				lvl: "debug",
			},
			wantErr: false,
		},
		{
			name: "trace",
			args: args{
				lvl: "trace",
			},
			wantErr: false,
		},
		{
			name: "info",
			args: args{
				lvl: "info",
			},
			wantErr: false,
		},
		{
			name: "warn",
			args: args{
				lvl: "warn",
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				lvl: "error",
			},
			wantErr: false,
		},
		{
			name: "fatal",
			args: args{
				lvl: "fatal",
			},
			wantErr: false,
		},
		{
			name: "panic",
			args: args{
				lvl: "panic",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				lvl: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetLevel(tt.args.lvl); (err != nil) != tt.wantErr {
				t.Errorf("SetLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDebug(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Debug(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.DebugLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestInfo(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Info(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.InfoLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestWarn(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Warn(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.WarnLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestError(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Error(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.ErrorLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestPanic(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")

	defer func() {
		recover()

		if assert.NotEmpty(t, h.entries) {
			assert.Equal(t, log.PanicLevel, h.entries[0].Level)
			assert.Equal(t, message, h.entries[0].Message)
		}
	}()

	Panic(message)

}

func TestDebugf(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Debugf("%s", message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.DebugLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestInfof(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Infof("%s", message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.InfoLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestWarnf(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Warnf("%s", message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.WarnLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestErrorf(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Errorf("%s", message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.ErrorLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestPanicf(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")

	defer func() {
		recover()

		if assert.NotEmpty(t, h.entries) {
			assert.Equal(t, log.PanicLevel, h.entries[0].Level)
			assert.Equal(t, message, h.entries[0].Message)
		}
	}()

	Panicf("%s", message)
}

func TestDebugln(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Debugln(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.DebugLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestInfoln(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Infoln(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.InfoLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestWarnln(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Warnln(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.WarnLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestErrorln(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")
	Errorln(message)

	if assert.NotEmpty(t, h.entries) {
		assert.Equal(t, log.ErrorLevel, h.entries[0].Level)
		assert.Equal(t, message, h.entries[0].Message)
	}
}

func TestPanicln(t *testing.T) {
	var (
		h = &MockHook{}

		message = "something"
	)

	AddHook(h)

	SetLevel("trace")

	defer func() {
		recover()

		if assert.NotEmpty(t, h.entries) {
			assert.Equal(t, log.PanicLevel, h.entries[0].Level)
			assert.Equal(t, message, h.entries[0].Message)
		}
	}()

	Panicln(message)
}

func TestWithError(t *testing.T) {
	entry := WithError(assert.AnError)
	assert.ErrorIs(t, entry.Data[log.ErrorKey].(error), assert.AnError)
}

func TestWithField(t *testing.T) {
	var (
		key = "key"
		val = "value"
	)

	entry := WithField(key, val)

	assert.Equal(t, val, entry.Data[key])
}

func TestWithContext(t *testing.T) {
	var (
		ctx = context.Background()
	)

	entry := WithContext(ctx)

	assert.Equal(t, ctx, entry.Context)
}

func TestWithFields(t *testing.T) {
	var (
		key1 = "key"
		val1 = "value"
		key2 = "key2"
		val2 = "value2"
	)

	entry := WithFields(log.Fields{
		key1: val1,
		key2: val2,
	})

	assert.Equal(t, val1, entry.Data[key1])
	assert.Equal(t, val2, entry.Data[key2])
}
