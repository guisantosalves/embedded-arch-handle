package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		// pos x + 1
		x, y := robotgo.GetMousePos()
		robotgo.MoveMouse(x+10, y)
		time.Sleep(5 * time.Second)

		// pos x - 1
		robotgo.MoveMouse(x-10, y)
		time.Sleep(5 * time.Second)
	}
}
