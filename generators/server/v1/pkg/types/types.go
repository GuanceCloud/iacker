package types

import "time"

// ParseDuration parses a string and returns the time.Duration value it
// represents. If the string cannot be parsed, it returns 0 and an error.
func ParseDuration(value string) (time.Duration, error) {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return time.Duration(0), err
	}
	return duration, nil
}

// ParseDateTime parses a date and time string in RFC3339 format.
//
// The returned time will be in UTC.
//
// This function is provided for convenience. It is a wrapper around
// time.Parse, with the time.RFC3339 format.
func ParseDateTime(value string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// ParseTime parses a time string and returns a time object.
// value is a string in the format "15:04:05Z07:00".
// If the string is not in that format, an error is returned.
func ParseTime(value string) (time.Time, error) {
	t, err := time.Parse("15:04:05", value)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
