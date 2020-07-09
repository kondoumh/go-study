package main

import "fmt"

func main() {
	var i interface{}
	i = 4
	fmt.Println(i)
	i = 4.5
	fmt.Println(i)
	i = "foo"
	fmt.Println(i)

	str, _ := i.(string)
	fmt.Println(str)
}
