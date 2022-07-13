package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	var command_output string
	cmd := exec.Command("bash", "-c", "python3 /home/rahma/sound_detection/SoundDetection.py print_prediction /home/rahma/sound_detection/testing_data/whatsapp.wav")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		log.Println(err.Error())
		return
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			// TODO kill process here
			log.Println(err.Error())
			break
		}
		str := string(tmp)
		if len(str) != 0 {
			command_output += str
			//log.Println(str)
		}
	}
	if strings.Contains(command_output, "doorbell") {
		fmt.Println("doorbell")
	} else if strings.Contains(command_output, "knock") {
		fmt.Println("knock")
	} else {
		fmt.Println("other")
	}
}
