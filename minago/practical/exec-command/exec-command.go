package main

import (
	"os/exec"
	"fmt"
)

func main() {
	out, err := exec.Command("uname", "-s").Output()
	if (err != nil) {
		fmt.Println("error! : ", err)
	} else {
		fmt.Printf("%s", string(out))
	}
}