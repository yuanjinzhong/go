package main

import (
	"fmt"
	"reflect"
)

var a string = "1"

var b *string = &a

func main() {
	fmt.Println(a, &b, *b, &*b, b)
	fmt.Printf("a type: %T\n", a)
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("*b type: %T\n", *b)
	fmt.Println(reflect.TypeOf(b).String())

}
