package formatter

import (
	"bytes"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	gray    = 35
	blue    = 36
)

type Colored struct{}

// Format renders a single log entry
func (f Colored) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.printColored(b, entry)
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f Colored) printColored(b *bytes.Buffer, entry *logrus.Entry) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	fmt.Fprintf(b, "\x1b[%dm%-44s\x1b[0m", levelColor, entry.Message)
}
