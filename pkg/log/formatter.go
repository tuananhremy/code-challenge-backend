package log

import (
	"code-challenge/pkg/dateutil"
	log "github.com/sirupsen/logrus"
)

type jstFormatter struct {
	log.Formatter
}

func (f *jstFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.In(dateutil.LocJP)
	return f.Formatter.Format(e)
}
