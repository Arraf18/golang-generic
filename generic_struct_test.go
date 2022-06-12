package golang_generic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	FirstName T
	LastName  T
}

func (d Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d Data[T]) changeFirstName(firstName T) T {
	d.FirstName = firstName
	return firstName
}

func TestMethodData(t *testing.T) {
	data := Data[string]{
		FirstName: "Arie",
		LastName:  "Afriyanto",
	}

	//testing untuk test changeName
	changeName := data.changeFirstName("Ria")
	assert.Equal(t, "Ria", changeName)

	// testing untuk sayHello
	hello := data.SayHello("Arie")
	assert.Equal(t, "Hello Arie", hello)
}

func TestData(t *testing.T) {
	data := Data[string]{
		FirstName: "Arie",
		LastName:  "Afriyanto",
	}

	fmt.Println(data)

	data2 := Data[int]{
		FirstName: 20,
		LastName:  300,
	}
	fmt.Println(data2)
}
