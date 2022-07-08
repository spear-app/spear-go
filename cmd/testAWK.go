package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*err := handlers.GetSpeakersAndDurationInFiles()
	if err != nil {
		log.Println(err.Error())
	}*/
	timeInFile, err := strconv.ParseFloat("3.14159265", 8)
	fmt.Println(timeInFile, err.Error())

}
