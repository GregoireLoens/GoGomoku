package ai

import (
	"fmt"
)

func Start(comChan chan string) {
	msg := <- comChan
	fmt.Println("On get " + msg)
	comChan <- "toto"
}
