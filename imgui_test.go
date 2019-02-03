package imgui_test

import (
	"testing"

	"github.com/inkyblackness/imgui-go"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	version := imgui.Version()
	assert.Equal(t, "1.67", version)
}
