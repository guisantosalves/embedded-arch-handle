package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	fmt.Println("Iniciando anti-idle para Teams...")
	fmt.Println("Pressione Ctrl+C para parar o programa.")

	// intervalor para mover o mouse a cada 30 segundos
	intervalo := 5 * time.Second

	ticker := time.NewTicker(intervalo)

	// Quando a função terminar, executa Stop() para liberar recursos e evitar leak de memória
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// x, y := robotgo.GetMousePos()
			// robotgo.MoveSmooth(x+10, y)
			// robotgo.MoveSmooth(x-10, y)
			// time.Sleep(30 * time.Second)
			executeAction()
		}
	}
}

func executeAction() {
	robotgo.KeyTap("scroll_lock")
	time.Sleep(100 * time.Second)
	robotgo.KeyTap("scroll_lock")
}
