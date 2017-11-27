package gogomoku_main

import (
	"fmt"
	"gogomoku_ai"
)

func main() {
	fmt.Print("main")
	go gogomoku_ai.Start()
}
