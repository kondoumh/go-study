package main

import "syscall/js"

func main() {
	window := js.Global()

	message := window.Get("document").Call("getElementById", "message")

	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		message.Set("innerHTML", "Clicked!!")
		return nil
	})

	message.Call("addEventListener", "click", cb)

	select {}
}