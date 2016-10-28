package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	FirstName string
	LastName  string
	Age       int
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Printf("Field Name: %s, Field Value: %v\n", typeField.Name, valueField.Interface())
	}
}

func main() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}

	f.reflect()
}
