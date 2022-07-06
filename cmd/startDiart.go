package main

import (
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

	if err := exec.CommandContext(ctx, "/home/rahma/miniconda3/condabin/conda", "activate", "diart").Run(); err != nil {
		// here we write what happens if the process took more than what we specified in the context.
		// and after interrupt signal
		fmt.Println("interrupt :((")
	}
	fmt.Println("successful, run the sound")
}
