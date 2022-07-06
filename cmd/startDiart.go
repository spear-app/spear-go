package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	StartConversation()
}
func StartConversation() {
	//exec.Command("/home/rahma/miniconda3/condabin/conda", "activate","diart")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/home/rahma/miniconda3/condabin/conda", "init", "bash")
	cmdConda := exec.CommandContext(ctx, "/home/rahma/miniconda3/condabin/conda", "activate", "diart")
	cmdPython3 := exec.CommandContext(ctx, "/usr/bin/python3", "-m", "diart.stream", "microphone", "--output=sound_output")

	//cmd := exec.CommandContext(ctx, "sh", "/home/rahma/spear-go/pkg/scripts/diart_run.sh")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
	fmt.Println("---------------------------------------")
	var out2 bytes.Buffer
	var stderr2 bytes.Buffer
	cmdConda.Stdout = &out2
	cmdConda.Stderr = &stderr2
	err2 := cmdConda.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err2) + ": " + stderr2.String())
		return
	}
	fmt.Println("Result: " + out2.String())
	fmt.Println("---------------------------------------")
	var out3 bytes.Buffer
	var stderr3 bytes.Buffer
	cmdPython3.Stdout = &out3
	cmdPython3.Stderr = &stderr3
	err3 := cmdPython3.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err3) + ": " + stderr3.String())
		return
	}
	fmt.Println("Result: " + out3.String())
}
