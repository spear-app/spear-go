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
	/*var outb, errb bytes.Buffer
	errb.Reset()
	outb.Reset()
	cmd.Stdout = &outb
	cmd.Stderr = &errb*/
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err.Error())
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		str := string(tmp)
		if len(str) == 1024 {
			fmt.Print("str len:", len(str), "\noutput:\n", str)
			break
		}
		/*fmt.Print("str len:", len(str), "\noutput:\n", str)
		if err != nil {
			break
		}*/
	}
	/*fmt.Println("sleeping-----------")
	cmd2 := exec.Command("sleep", "20")
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("slept")*/
	/*for true {
		//fmt.Println("out:", outb.String(), "err:", errb.String())

		if len(outb.String()) > 0 {
			fmt.Println("out:", outb.String(), "err:", errb.String())
		}
	}
	fmt.Println("out:", outb.String(), "err:", errb.String())*/
	cmd.Wait()

	/*pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		fmt.Println("killing the process")
		err := syscall.Kill(-pgid, 15)
		if err != nil {
			log.Fatal("failed to kill")
		} else {
			fmt.Println("process killed")
		}
	}

	cmd.Wait()*/

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
