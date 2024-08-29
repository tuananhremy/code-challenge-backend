package log

import (
	"reflect"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func Test_jstFormatter_Format(t *testing.T) {
	var (
		dt = time.Date(2024, 3, 1, 10, 0, 0, 0, time.UTC)
	)

	type fields struct {
		Formatter log.Formatter
	}
	type args struct {
		e *log.Entry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Formatter: &log.TextFormatter{},
			},
			args: args{
				e: &log.Entry{
					Level:   log.InfoLevel,
					Message: "test",
					Time:    dt,
				},
			},
			want:    []byte("time=\"2024-03-01T19:00:00+09:00\" level=info msg=test\n"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &jstFormatter{
				Formatter: tt.fields.Formatter,
			}
			got, err := f.Format(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("jstFormatter.Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jstFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
