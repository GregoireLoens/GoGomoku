package ai

import (
	"strconv"
	"time"
)

var GameBoard [][]int

func returnChan(comChan chan<- string, x int, y int) {
	GameBoard[x][y] = 2
	sX := strconv.Itoa(x)
	sY := strconv.Itoa(y)
	comChan <- sX + "," + sY
}

func Start(comChan chan<- string) {
	time.Sleep(time.Second * 2)
}

func StartRandom(comChan chan<- string) {
	for x := range GameBoard {
		for y := range GameBoard[x] {
			if GameBoard[x][y] == 0 {
				returnChan(comChan, x, y)
				return
			}
		}
	}
}
