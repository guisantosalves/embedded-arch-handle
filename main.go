package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		// pos x + 1
		x, y := robotgo.GetMousePos()
		robotgo.MoveSmooth(x+10, y)
		robotgo.MoveSmooth(x-10, y)
		time.Sleep(30 * time.Second)
	}
}
