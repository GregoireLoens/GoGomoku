package main

import (
	"fmt"
	"gogomoku/ai"
)

func main() {
	fmt.Print("main")
	go ai.Start()
}
