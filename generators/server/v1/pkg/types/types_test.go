package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestParseDuration tests the ParseDuration function.
func TestParseDuration(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    time.Duration
		wantErr bool
	}{
		{
			name:    "test_1",
			value:   "1ms",
			want:    time.Duration(1e6),
			wantErr: false,
		},
		{
			name:    "test_2",
			value:   "1s",
			want:    time.Duration(1e9),
			wantErr: false,
		},
		{
			name:    "test_3",
			value:   "1m",
			want:    time.Duration(6e10),
			wantErr: false,
		},
		{
			name:    "test_4",
			value:   "1h",
			want:    time.Duration(3.6e12),
			wantErr: false,
		},
		{
			name:    "test_8",
			value:   "1",
			want:    time.Duration(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDuration(tt.value)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseDateTime(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "empty",
			value:    "",
			expected: time.Time{},
			wantErr:  true,
		},
		{
			name:     "invalid",
			value:    "2017-01",
			expected: time.Time{},
			wantErr:  true,
		},
		{
			name:     "valid",
			value:    "2017-01-01T00:00:00Z",
			expected: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseDateTime(c.value)
			assert.Equal(t, c.expected, got)
			if c.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		value   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "should return time with no error",
			value:   "15:04:05",
			want:    time.Date(0, time.January, 1, 15, 4, 5, 0, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTime(tt.value)
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
