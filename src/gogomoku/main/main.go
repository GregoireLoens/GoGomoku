package main

import (
	"fmt"
	"../ai"
)

func main() {
	fmt.Print("main")
	go ai.Start()
}
