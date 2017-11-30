package ai

import (
	"strconv"
)

var GameBoard [][]int

type Position struct {
	x int
	y int
}

var LastEnemyPosition = Position{-1, -1}
var LastPlayerPosition = Position{-1, -1}

func hasPlayed(pos Position) bool {
	return pos.x != -1 && pos.y != -1
}

func otherPlayer(player int) int {
	if player == 1 {
		return 2
	} else {
		return 1
	}
}

func calcWeightByStar(position Position, turn int) int {
	var weight = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	func(weightToSet int) {
		for x := position.x; x >= 0 && x > position.x-4; x-- {
			if GameBoard[x][position.y] == turn {
				weightToSet++
			} else if GameBoard[x][position.y] == otherPlayer(turn) {
				break
			}
		}
	}(weight[0])
	func(weightToSet int) {
		for i := 0; i < 4; i++ {
			if position.x-i >= 0 && position.y+i < len(GameBoard) {
				if GameBoard[position.x-i][position.y+i] == turn {
					weightToSet++
				} else if GameBoard[position.x-i][position.y+i] == otherPlayer(turn) {
					break
				}
			}
		}
	}(weight[1])
	func(weightToSet int) {
		for y := position.y; y >= 0 && y < position.y+4; y++ {
			if GameBoard[position.x][y] == turn {
				weightToSet++
			} else if GameBoard[position.x][y] == otherPlayer(turn) {
				break
			}
		}
	}(weight[2])
	func(weightToSet int) {
		for i := 0; i < 4; i++ {
			if position.x+i >= 0 && position.y+i < len(GameBoard) {
				if GameBoard[position.x+i][position.y+i] == turn {
					weightToSet++
				} else if GameBoard[position.x-i][position.y+i] == otherPlayer(turn) {
					break
				}
			}
		}
	}(weight[3])
	func(weightToSet int) {
		for x := position.x; x >= 0 && x < position.x+4; x++ {
			if GameBoard[x][position.y] == turn {
				weightToSet++
			} else if GameBoard[x][position.y] == otherPlayer(turn) {
				break
			}
		}
	}(weight[4])
	func(weightToSet int) {
		for i := 0; i < 4; i++ {
			if position.x+i >= 0 && position.y-i > 0 {
				if GameBoard[position.x+i][position.y-i] == turn {
					weightToSet++
				} else if GameBoard[position.x+i][position.y-i] == otherPlayer(turn) {
					break
				}
			}
		}
	}(weight[5])
	func(weightToSet int) {
		for y := position.y; y >= 0 && y > position.y-4; y-- {
			if GameBoard[position.x][y] == turn {
				weightToSet++
			} else if GameBoard[position.x][y] == otherPlayer(turn) {
				break
			}
		}
	}(weight[6])
	func(weightToSet int) {
		for i := 0; i < 4; i++ {
			if position.x-i >= 0 && position.y-i > 0 {
				if GameBoard[position.x-i][position.y-i] == turn {
					weightToSet++
				} else if GameBoard[position.x-i][position.y-i] == otherPlayer(turn) {
					break
				}
			}
		}
	}(weight[7])
	return weight[0] ^ 3 + weight[1] ^ 3 + weight[2] ^ 3 + weight[3] ^ 3 +
			weight[4] ^ 3 + weight[5] ^ 3 + weight[6] ^ 3 + weight[7] ^ 3
}

func firstTurn() Position {
	return Position{len(GameBoard) / 2, len(GameBoard) / 2}
}

func turn() Position {
	var enemyWeight int = -1
	var playerWeight int = -1
	if hasPlayed(LastPlayerPosition) {
		playerWeight = calcWeightByStar(LastPlayerPosition, 1)
	}
	if hasPlayed(LastEnemyPosition) {
		enemyWeight = calcWeightByStar(LastEnemyPosition, 2)
	}
	if enemyWeight == -1 && playerWeight == -1 {
		return firstTurn()
	}

}

func returnChan(comChan chan<- string, x int, y int) {
	GameBoard[x][y] = 2
	sX := strconv.Itoa(x + 1)
	sY := strconv.Itoa(y + 1)
	comChan <- sX + "," + sY
	LastPlayerPosition.x = x
	LastPlayerPosition.y = y
}

func Start(comChan chan<- string) {
	var pos = turn()
	returnChan(comChan, pos.x, pos.y)
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
