package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func Length[T]() {
}
func Length[T interface{}]() {
}
*/

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

// parameter itu = input data dari luar function
func kali(a int, b int, c int) int {
	hasil := a * b * c
	return hasil
}

func TestSample(t *testing.T) {
	// cara memanggil function
	hasil := kali(10, 3, 9)
	assert.Equal(t, 270, hasil)

	hasil = kali(40, 2, 8)
	assert.Equal(t, 640, hasil)
	assert.True(t, true)

	//type parameter
	var result1 = Length[string]("Arie")
	assert.Equal(t, "Arie", result1)

	var resultNumber = Length[int](100)
	assert.Equal(t, 100, resultNumber)

	var resultBool = Length[bool](true)
	assert.True(t, resultBool)

	var resultFloat = Length[float64](100.30)
	assert.Equal(t, 100.30, resultFloat)
}
