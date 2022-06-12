package golang_generic

import (
	"fmt"
	"strconv"
	"testing"
)

func NewArray[T any](b, k int) [][]T {
	var array = make([][]T, b)
	for i := range array {
		array[i] = make([]T, k)
	}
	return array
}

func Kotak(n int) [][]int {
	array := NewArray[int](n, n)
	angka := 1
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			array[i][k] = angka
			angka += 5
		}
	}
	return array
}

func Kotak2(n int) [][]string {
	array := NewArray[string](n, n)
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			array[a][b] = strconv.Itoa(a) + "," + strconv.Itoa(b)
		}
	}
	return array
}

func TestArray(t *testing.T) {
	arr := Kotak(5)
	fmt.Println(arr)
	fmt.Println()

	arr2 := Kotak2(4)
	fmt.Println(arr2)
}
