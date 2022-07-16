package main

import (
	"fmt"
	"github.com/spear-app/spear-go/pkg/handlers"
	"log"
)

func main() {
	text, err := handlers.GetText("/home/rahma/recorded_audio/1.opusconverted.wav")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(text)
}
