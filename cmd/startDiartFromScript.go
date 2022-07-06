package main

import (
	"fmt"
	"log"
	"os/exec"
)

func startConv() *exec.Cmd {
	fmt.Println("starting conversation .........")
	cmd := exec.Command("bash", "-c", "source "+"/home/rahma/spear-go/pkg/scripts/diart_run4.sh")
	if err := cmd.Run(); err != nil {
		log.Fatal(err.Error())
	}
	cmd2 := exec.Command("sleep", "20")
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	err := cmd.Process.Kill()

	if err != nil {
		fmt.Println("failed to kill the process")
	} else {
		fmt.Println("process killed")
	}
	return cmd
}
func main() {
	startConv()
	/*fmt.Println("sleep now")
	cmd2 := exec.Command("sleep", "20")
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("slept")
	err := cmd.Process.Kill()
	if err != nil {
		fmt.Println("failed to kill the process")
	} else {
		fmt.Println("process killed")
	}*/
}
