package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func startConv() *exec.Cmd {
	fmt.Println("starting conversation .........")
	cmd := exec.Command("bash", "-c", "source "+"/home/rahma/spear-go/pkg/scripts/diart_run4.sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("sleeping-----------")
	cmd2 := exec.Command("sleep", "20")
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("slept")
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		fmt.Println("killing the process")
		err := syscall.Kill(-pgid, 15)
		if err != nil {
			log.Fatal("failed to kill")
		} else {
			fmt.Println("process killed")
		}
	}

	cmd.Wait()
	/*err := cmd.Process.Kill()

	if err != nil {
		fmt.Println("failed to kill the process")
	} else {
		fmt.Println("process killed")
	}*/

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
