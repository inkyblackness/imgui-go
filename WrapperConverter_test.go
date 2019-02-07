package imgui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringBufferAllocation(t *testing.T) {
	tt := []struct {
		initialValue string
		expectedSize int
	}{
		{"", 1},
		{"a", 2},
		{"123456789", 10},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("<%s>", tc.initialValue), func(t *testing.T) {
			buf := newStringBuffer(tc.initialValue)
			defer buf.free()
			assert.Equal(t, tc.expectedSize, buf.size)
		})
	}
}

func TestStringBufferStorage(t *testing.T) {
	tt := []string{"", "a", "ab", "SomeLongerText"}

	for _, tc := range tt {
		t.Run("Value <"+tc+">", func(t *testing.T) {
			buf := newStringBuffer(tc)
			require.NotNil(t, buf, "buffer expected")
			defer buf.free()
			result := buf.toGo()
			assert.Equal(t, tc, result)
		})
	}
}

func TestStringBufferResize(t *testing.T) {
	tt := []struct {
		initialValue  string
		newSize       int
		expectedValue string
	}{
		{"", 10, ""},
		{"abcd", 10, "abcd"},
		{"abcd", 3, "ab"},
		{"efgh", 0, ""},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("<%s> -> %d", tc.initialValue, tc.newSize), func(t *testing.T) {
			buf := newStringBuffer(tc.initialValue)
			defer buf.free()
			buf.resizeTo(tc.newSize)
			assert.Equal(t, tc.expectedValue, buf.toGo())
		})
	}
}
