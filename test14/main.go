package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{} = make(chan int)

	switch value := v.(type) {
	case int:
		fmt.Printf("Integer: %v\n", value)
	case float64:
		fmt.Printf("Float64: %v\n", value)
	case string:
		fmt.Printf("String: %v\n", value)
	case bool:
		fmt.Printf("Bool: %v\n", value)
	default:
		var r = reflect.TypeOf(value)
		fmt.Printf("Other: %v\n", r)
	}
}
