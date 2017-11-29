package ai

import (
	"fmt"
)

var GameBoard [][]int

func Start(comChan chan string) {
	comChan <- "YES"
	fmt.Println(GameBoard)
}
