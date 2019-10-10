package main

import "syscall/js"

func main() {
	window := js.Global()

	message := window.Get("document").Call("getElementById", "message")

	cb := js.NewCallback(func(args []js.Value) {
		message.Set("innerHTML", "Clicked!!")
	})

	message.Call("addEventListener", "click", cb)

	select {}
}