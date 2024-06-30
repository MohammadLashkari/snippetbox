package main

import (
	"testing"
	"time"

	"github.com/MohammadLashkari/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 06, 29, 21, 05, 0, 0, time.UTC),
			want: "29 Jun 2024 at 21:05",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 06, 29, 21, 05, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "29 Jun 2024 at 20:05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanDate(tt.tm)
			assert.Equal(t, got, tt.want)
		})
	}

}
