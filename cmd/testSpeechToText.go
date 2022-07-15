package main

import "github.com/spear-app/spear-go/pkg/handlers"

func main() {
	str, _ := handlers.GetText("/home/rahma/speech_files/7.wav")
	println(str)
}
