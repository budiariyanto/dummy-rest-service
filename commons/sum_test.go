package commons

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddNum(t *testing.T) {
	result := AddNums(1, 2)
	assert.Equal(t, 3, result, "Should be 3")
}
