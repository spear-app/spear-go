package main

import (
	"fmt"
	"github.com/spear-app/spear-go/pkg/handlers"
	"os/exec"
	"time"
)

func main() {
	startTime, err := handlers.StartConversation()
	tmp1 := startTime.Format("15:04:05")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("start time: ", tmp1)
	exec.Command("sleep", "5")
	audioPlayTime := time.Now()
	tmp2 := audioPlayTime.Format("15:04:05")
	err = handlers.PlayAudio("/home/rahma/conversation_audio/1.wav")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("end time: ", tmp2)
	duration, err := handlers.SubtractTime(startTime, audioPlayTime)
	if err != nil {
		err.Error()
	}
	fmt.Println("duration: ", duration)
}
