package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

type Person struct {
	Name string
	Age  int
}

func typeOf(x any) string {
	return reflect.TypeOf(x).String()
}

func main() {
	{
		x := 1
		fmt.Println(typeOf(x))
	}

	{
		x := "hello"
		fmt.Println(typeOf(x))
	}

	{
		x := true
		fmt.Println(typeOf(x))
	}

	{
		x := make(chan Person, 3)
		fmt.Println(typeOf(x))
	}
}
