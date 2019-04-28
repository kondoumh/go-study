package main

import (
	"fmt"
	"encoding/base64"
	"crypto/rand"
	"io"
)

func main() {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	readableKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println(readableKey)
}