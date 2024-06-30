package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v; want: %v", got, want)
	}
}

func StringContains(t *testing.T, got, wantedSubstring string) {
	t.Helper()
	if !strings.Contains(got, wantedSubstring) {
		t.Errorf("got: %q; want to contains: %q", got, wantedSubstring)
	}
}
