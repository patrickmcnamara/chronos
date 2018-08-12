/*
Package chronos provides utilities for working with the chronos time system.

The chronos time system splits the day into 16 * 16 * 16 * 16 parts instead of
the usual 24 * 60 * 60. It displays this time using four hexadecimal digits. For
example, 0000 is the start of the day, 8000 is half way through the day and FFFF
is the last part of the day.
*/
package chronos

import (
	"fmt"
	"strconv"
	"time"
)

// Chronos represents a chronos. This is only a time during a day, the same time
// tomorrow is the exact same chronos. It does not store day or date at all and
// will loop in functions such as Add() or Equal().
type Chronos uint64

const (
	// MaxChronos is the maximum value of a chronos.
	MaxChronos Chronos = 4294967296
	// MinChronos is the minimum value of a chronos.
	MinChronos Chronos = 0

	// maxNano is the maximum number of nanoseconds in a day.
	maxNano uint64 = 24 * 60 * 60 * 1000000000
)

// Now returns the current chronos.
func Now() Chronos {
	now := time.Now()
	return FromTime(now)
}

// FromTime returns a chronos from a time.Time. Ignores the day or date.
func FromTime(t time.Time) Chronos {
	t = t.UTC()
	midnight := t.Truncate(24 * time.Hour)
	proportion := float64(time.Since(midnight).Nanoseconds()) / float64(maxNano)
	chronos := uint64(proportion * float64(MaxChronos))
	return Chronos(chronos % uint64(MaxChronos))
}

// Parse returns a chronos from a chronos string.
func Parse(s string) Chronos {
	chronos, _ := strconv.ParseUint(s, 16, 64)
	chronos *= 1 << (4 * uint(len(s)))
	return Chronos(chronos % uint64(MaxChronos))
}

// FromDuration returns a chronos from a time.Duration. Loops over days.
func FromDuration(d time.Duration) Chronos {
	proportion := float64(d.Nanoseconds()) / float64(maxNano)
	chronos := uint64(proportion * float64(MaxChronos))
	return Chronos(chronos % uint64(MaxChronos))
}

// ToDuration creates a time.Duration from a chronos. Loops over days.
func (c Chronos) ToDuration() time.Duration {
	proportion := float64(c%MaxChronos) / float64(MaxChronos)
	nanoseconds := uint64(proportion * float64(maxNano))
	duration, _ := time.ParseDuration(fmt.Sprintf("%dns", nanoseconds))
	return duration
}

// Add adds two chronoses. This loops to the next or previous day if it goes
// over FFFF or below 0000. For example, 0000 - 0001 equals, because it looped
// back to the previous day.
func (c Chronos) Add(h Chronos) Chronos {
	return Chronos((c + h) % MaxChronos)
}

// Sub subtracts two chronoses. This loops to the next or previous day if it
// goes over FFFF or below 0000. For example, FFFF + 0001 equals 0000, because
// it looped over to the next day.
func (c Chronos) Sub(h Chronos) Chronos {
	return Chronos((c - h) % MaxChronos)
}

// Equal checks if two chronoses are equal. This loops to the next or previous
// day if it goes over FFFF or below 0000. For example, B000 + B000 is equal to
// 8000, because it looped over to the next day.
func (c Chronos) Equal(h Chronos) bool {
	return c%MaxChronos == h%MaxChronos
}

// String returns the standard precision chronos string.
func (c Chronos) String() string {
	return c.FullString()[:4]
}

// FullString returns the maximum precision chronos string.
func (c Chronos) FullString() string {
	return fmt.Sprintf("%08X", uint64(c))
}
