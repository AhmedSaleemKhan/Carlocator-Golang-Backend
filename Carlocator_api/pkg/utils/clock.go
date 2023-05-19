package utils

import "time"

type Clock interface {
	NowUnix() int64
	NowUTC() time.Time
	// WARNING: You most likely do NOT want to use this method, you probably want to use NowUTC().
	// This function will use the hardware's timezone as the Location of the Time struct. This
	// could cause problems in the database (particularly if column is TIMESTAMP WITHOUT TIME ZONE).
	NowLocal() time.Time
}

type ClockImpl struct{}

// Assert that *clockImpl satisfies the Clock interface
var _ Clock = &ClockImpl{}

func NewClock() *ClockImpl {
	return &ClockImpl{}
}

func (c *ClockImpl) NowUnix() int64 {
	nowUTC := time.Now().UTC()
	return *ToUnix(&nowUTC)
}

func (c *ClockImpl) NowUTC() time.Time {
	return time.Now().UTC()
}

func (c *ClockImpl) NowLocal() time.Time {
	return time.Now()
}

// ToUnix converts a *Time (any locale/timezone will work) to a Unix *int64 timestamp in
// milliseconds
// Needed because time package's only has methods to convert to seconds (.Unix()) or nanoseconds
// (.UnixNano())
func ToUnix(t *time.Time) *int64 {
	if t == nil {
		return nil
	}
	millis := t.Round(time.Millisecond).UnixNano() / int64(time.Millisecond)
	return &millis
}

// ToTime converts a *int64 Unix timestamp in milliseconds to a *Time in UTC
func ToTime(unixMillis *int64) *time.Time {
	if unixMillis == nil {
		return nil
	}
	t := time.Unix(0, *unixMillis*int64(time.Millisecond)).In(time.UTC)
	return &t
}
