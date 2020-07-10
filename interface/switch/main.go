package main

import "fmt"

func main() {
	var i interface{}
	i = 3
	checkType(i)
	i = 3.5
	checkType(i)
	i = "aaaaa"
	checkType(i)

	type MyType uint
	i = MyType(4)
	checkType(i)
}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other")
	}
}
