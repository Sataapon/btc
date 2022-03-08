package time

import (
	"time"
)

const Layout = time.RFC3339

type Time struct {
	t time.Time
}

func New(timeStr string) (Time, error) {
	t, err := time.Parse(Layout, timeStr)
	if err != nil {
		return Time{}, err
	}
	return Time{t: t}, nil
}

func (t Time) String() string {
	return t.t.Format(Layout)
}

func (t Time) After(ot Time) bool {
	return t.t.After(ot.t)
}

func (t Time) Before(ot Time) bool {
	return t.t.Before(ot.t)
}
