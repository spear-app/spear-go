package main

import "github.com/spear-app/spear-go/pkg/handlers"

func main() {
	err := handlers.PlayAudio("/home/rahma/conversation_audio/1.wav")
	if err != nil {
		println(err)
	}
}
