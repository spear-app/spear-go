package main

import (
	"fmt"
	"github.com/spear-app/spear-go/pkg/handlers"
)

func main() {
	/*err := handlers.GetSpeakersAndDurationInFiles()
	if err != nil {
		log.Println(err.Error())
	}*/
	/*timeInFile, err := strconv.ParseFloat("3.141", 32)
	fmt.Println(timeInFile, err)*/
	output, err := handlers.GetSpeakerFromDiartOutput("/home/rahma/sound_output/live_recording.rttm", 563)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("output: ", output)
}
