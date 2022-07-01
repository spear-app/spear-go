package main

import (
	"github.com/spear-app/spear-go/pkg/handlers"
	"github.com/subosito/gotenv"
)

func Init() {
	gotenv.Load()
}

func main() {
	Init()
	handlers.Start()
}
