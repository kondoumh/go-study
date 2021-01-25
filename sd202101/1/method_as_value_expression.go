package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	fv := Hex(1024).String
	fmt.Println(fv())

	fe := Hex.String
	fmt.Println(fe(1024))
}
