package timeutil

import (
	"fmt"
	"time"
)

// ParseHHMM parses a "HH:MM" string into hours and minutes.
func ParseHHMM(s string) (int, int, error) {
	var h, m int
	_, err := fmt.Sscanf(s, "%d:%d", &h, &m)
	return h, m, err
}

// SlotFits reports whether a slot of duration d fits within the range [from, to].
func SlotFits(from, to time.Time, d time.Duration) bool {
	return !from.Add(d).After(to)
}
