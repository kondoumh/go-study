package main

import (
	"encoding/binary"
	crand"crypto/rand"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		fmt.Println("fallback to time seed")
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	n := rand.Intn(100)
	fmt.Println(n)
}