package main

import "github.com/spear-app/spear-go/pkg/handlers"

func main() {
	str, _ := handlers.GetText("/home/rahma/conversation_audio/2.wav")
	println(str)
}
