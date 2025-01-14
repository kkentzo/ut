package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Convert(t *testing.T) {
	target, err := time.Parse(time.RFC3339Nano, "2025-01-13T03:00:00+02:00")
	assert.NoError(t, err)

	kases := []struct {
		input          string
		expected       time.Time
		error_contains string
	}{
		// unix time in seconds
		{"1736730000", target, ""},
		// unix time in milliseconds
		{"1736730000000", target, ""},
		// unix time in microseconds
		{"1736730000000000", target, ""},
		// unix time in nanoseconds
		{"1736730000000000", target, ""},
		// ERRORS
		// invalid value
		{"foo", time.Time{}, "failed to convert"},
		// greater than seconds
		{"173673000", time.Time{}, "unrecognized resolution"},
		// smaller than nanoseconds
		{"17367300000000001", time.Time{}, "unrecognized resolution"},

	}

	for _, kase := range kases {
		out, err := convert(kase.input)
		assert.Equal(t, kase.expected, out, "failed input: "+kase.input)
		if err != nil {
			assert.ErrorContains(t, err, kase.error_contains)
		} else {
			assert.Empty(t, kase.error_contains)
		}
	}
}
