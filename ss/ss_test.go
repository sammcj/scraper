package ss

import (
	"strings"
	"testing"
)

func TestGetPrefix(t *testing.T) {
	tests := []struct {
		in  map[string]string
		pre string
		out string
		ok  bool
	}{
		{
			in: map[string]string{
				"media_box2d_eu": "test",
			},
			pre: "media_box2d_",
			out: "test",
			ok:  true,
		},
		{
			in: map[string]string{
				"media_box2d_eu_crc": "test",
			},
			pre: "media_box2d_",
		},
		{
			pre: "test",
		},
	}
	for _, tt := range tests {
		if got, ok := getPrefix(tt.in, tt.pre); got != tt.out || ok != tt.ok {
			t.Errorf("getPrefix(%v, %s) = (%s, %t); want (%s, %t)", tt.in, tt.pre, got, ok, tt.out, tt.ok)
		}
	}
}

func getPrefix(m map[string]string, s string) (string, bool) {
	if s == "" {
		return "", false
	}
	for k, v := range m {
		if strings.HasPrefix(k, s) {
			// Only match if the key is exactly one region suffix longer than the prefix
			// e.g. "media_box2d_eu" for prefix "media_box2d_"
			rest := k[len(s):]
			if !strings.Contains(rest, "_") {
				return v, true
			}
		}
	}
	return "", false
}
