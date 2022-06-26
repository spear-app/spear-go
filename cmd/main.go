package main

import (
	"github.com/spear-app/spear-go/pkg/handlers"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	handlers.Start()
}
func Hi() {

}
