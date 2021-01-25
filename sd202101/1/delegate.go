package main

import "fmt"

type Chip struct {
	Number int
}
type Card struct {
	string
	Chip
	Number int
}

func (c *Chip) Scan() {
	fmt.Println(c.Number)
}

func main() {
	c := Card{
		string: "Credit",
		Chip: Chip{
			Number: 4242424242424242,
		},
		Number: 5454545454545454,
	}
	c.Scan()
}
