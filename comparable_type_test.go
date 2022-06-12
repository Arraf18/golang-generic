package golang_generic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func IsSameInt(value1, value2 int64) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func IsSameString(value1, value2 string) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSame[string]("Roni", "Roni"))
	assert.True(t, IsSame[int](10, 10))
	assert.False(t, IsSame[int](10, 20))

	// non generic
	assert.True(t, IsSameInt(20, 20))
	// non generic string
	assert.False(t, IsSameString("Aqila", "Sandy"))
}
