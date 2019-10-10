package main

import "syscall/js"

func main() {
	window := js.Global()

	message := window.Get("document").Call("querySelector", "#message")

	message.Set("innerHTML", "Hello, WebAssembly")
}
