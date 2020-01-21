// Package chronos provides utilities for working with the chronos time system.
//
// The chronos time system splits the day into 16 * 16 * 16 * 16 moments instead of 24 * 60 * 60 seconds.
// This is then represented using a four digit hexadecimal number.
// For example, *0000* is the start of the day, *8000* is half way through the day and *FFFF* is the last moment of the day.
//
package chronos

import (
	"fmt"
	"strconv"
	"time"
)

// Chronos represents a chronos. It represents how far through a day is.
type Chronos uint16

const (
	// MaxChronos is the maximum value of a chronos.
	MaxChronos Chronos = 1<<16 - 1
	// MinChronos is the minimum value of a chronos.
	MinChronos Chronos = 0

	// maxNano is the maximum number of nanoseconds in a day.
	maxNano int64 = 24 * 60 * 60 * 1000000000
)

// Now returns the current time as a Chronos.
func Now() Chronos {
	now := time.Now()
	return FromTime(now)
}

// FromTime returns a Chronos from a time.Time. It ignores the date.
func FromTime(t time.Time) Chronos {
	midnight := t.UTC().Truncate(24 * time.Hour)
	nanoseconds := t.Sub(midnight).Nanoseconds()
	moments := (nanoseconds * (int64(MaxChronos) + 1)) / maxNano
	return Chronos(moments)
}

// FromDuration returns a Chronos from a time.Duration. Loops over days.
func FromDuration(d time.Duration) Chronos {
	moments := (d.Nanoseconds() * int64(MaxChronos)) / maxNano
	return Chronos(moments)
}

// ToTime creates a time.Time from a Chronos set to today's date. The timezone
// will be UTC.
func (c Chronos) ToTime() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, int(c)*(int(maxNano)/int(MaxChronos)), time.UTC)
}

// ToDuration creates a time.Duration from a Chronos. Loops over days.
func (c Chronos) ToDuration() time.Duration {
	nanoseconds := int64(c) * (maxNano / int64(MaxChronos))
	return time.Duration(nanoseconds)
}

// Parse returns a Chronos from a chronos string. The string will be truncated
// to a length of 4.
func Parse(s string) (c Chronos, err error) {
	precision := len(s)
	if precision > 4 {
		precision = 4
	}
	moments, err := strconv.ParseUint(fmt.Sprintf("%04s", s[:precision]), 16, 16)
	return Chronos(moments), err
}

// String returns a string representing the chronos in the form `CCCC`.
func (c Chronos) String() string {
	return fmt.Sprintf("%04X", uint16(c))
}

// Add returns the Chronos c+d. It loops over days.
func (c Chronos) Add(d Chronos) Chronos {
	return c + d
}

// Sub returns the Chronos c-d. It loops over days.
func (c Chronos) Sub(d Chronos) Chronos {
	return c - d
}
