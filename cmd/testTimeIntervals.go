package main

import (
	"errors"
	"os/exec"
	"time"
)

func main() {
	time1 := time.Now()
	exec.Command("/usr/bin/sleep", "7")
	time2 := time.Now()
	duration, err := SubtractTime(time1, time2)
	if err != nil {
		println(duration)
	}
}
func SubtractTime(time1 time.Time, time2 time.Time) (int, error) {
	hour1, min1, second1 := time1.Clock()
	hour2, min2, second2 := time2.Clock()
	if hour2-hour1 != 0 {
		return 0, errors.New("max conversation time is 15 minutes")
	}
	duration := (min2*60 + second2) - (min1*60 + second1)
	if duration <= 0 {
		return 0, errors.New("invalid time duration")
	}
	return duration, nil
}
