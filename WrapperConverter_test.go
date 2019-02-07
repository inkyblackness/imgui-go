package imgui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringBufferAllocation(t *testing.T) {
	tt := []struct {
		initialSize  int
		expectedSize int
	}{
		{1, 1},
		{2, 2},
		{10, 10},
		{0, 1},
		{-1, 1},
		{-10, 1},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d -> %d", tc.initialSize, tc.expectedSize), func(t *testing.T) {
			buf := newStringBuffer(tc.initialSize, "")
			defer buf.free()
			assert.Equal(t, tc.expectedSize, buf.size)
		})
	}
}

func TestStringBufferStorage(t *testing.T) {
	tt := []string{"", "a", "ab", "SomeLongerText"}

	for _, tc := range tt {
		t.Run("Value <"+tc+">", func(t *testing.T) {
			asBytes := []byte(tc)
			buf := newStringBuffer(len(asBytes)+1, tc)
			require.NotNil(t, buf, "buffer expected")
			defer buf.free()
			result := buf.toGo()
			assert.Equal(t, tc, result)
		})
	}
}

func TestStringBufferTruncation(t *testing.T) {
	buf := newStringBuffer(3, "abcd")
	defer buf.free()
	assert.Equal(t, "ab", buf.toGo())
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
			buf := newStringBuffer(len([]byte(tc.initialValue))+1, tc.initialValue)
			defer buf.free()
			buf.resizeTo(tc.newSize)
			assert.Equal(t, tc.expectedValue, buf.toGo())
		})
	}
}
