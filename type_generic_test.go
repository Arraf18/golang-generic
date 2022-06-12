package golang_generic

import (
	"fmt"
	"testing"
)

//type baru namanya Bag, sebagai Array
type Bag[T any] []T

//function yang digunakan untuk print type Bag
func PrintBag[T any](bag Bag[T]) {
	//karna Bag adalah Array, maka untuk print nya harus pakai for
	for _, value := range bag {
		fmt.Print(value, "\t")
	}
	fmt.Println()
}

//git test
func TestBag(t *testing.T) {
	number := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(number)

	names := Bag[string]{"Rio", "Arie", "Bambang", "Aji", "Joko"}
	PrintBag(names)
}
