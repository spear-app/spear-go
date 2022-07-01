package main

import (
	"github.com/spear-app/spear-go/pkg/handlers"
	"os/exec"
	"time"
)

func main() {
	time1 := time.Now()
	exec.Command("/usr/bin/sleep", "7")
	time2 := time.Now()
	duration, err := handlers.SubtractTime(time1, time2)
	if err != nil {
		println(duration)
	}
}
