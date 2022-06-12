package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	fmt.Println(param1)
	fmt.Println(param2)

	return param1, param2
}

func TestMultipleParam(t *testing.T) {
	r1, r2 := MultipleParameter[string, int]("Ini String", 10)
	assert.Equal(t, "Ini String", r1)
	assert.Equal(t, 20, r2)

	r1, r3 := MultipleParameter[string, string]("Aku", "Kamu")
	assert.Equal(t, "Aku", r1)
	assert.Equal(t, "Kamu", r3)
}
