package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Usage = func() {
		usageTxt := `Usage example [option]
An example of customizing usage output

	-s, --s String argument, default: String help message
	-i, --i INTEGER argument, default: Int help message
	-b, --b BOOLEAN argument, default: Bool help message`

		fmt.Fprintf(os.Stderr, "%s\n", usageTxt)
	}
	s := flag.String("s", "hello, workd!", "String help message")
	i := flag.Int("i", 1, "Int help message")
	b := flag.Bool("b", false, "Bool help message")

	flag.Parse()

	fmt.Println("str: ", *s)
	fmt.Println("int: ", *i)
	fmt.Println("bool: ", *b)
}

