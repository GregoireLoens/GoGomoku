package ai

import (
	"math"
	"fmt"
)

func isWinningPoint(origin Position, player int) bool {
	var nbPoint = 0

	for i := 4; i > -4; i++ {
		var tmpY = 0
		tmpY = origin.Y + i
		if posIsValid(Position{origin.X, tmpY}) {
			if GameBoard[origin.X][tmpY] == player {
				nbPoint++
			} else {
				nbPoint = 0
			}
			if nbPoint == 4 {
				return true
			}
		}
	}
	for i := 4; i > -4; i++ {
		var tmpX = 0
		tmpX = origin.X + i
		if posIsValid(Position{tmpX, origin.Y}) {
			if GameBoard[tmpX][origin.Y] == player {
				nbPoint++
			} else {
				nbPoint = 0
			}
			if nbPoint == 4 {
				return true
			}
		}
	}
	for i := 4; i > -4; i++ {
		var tmpX = 0
		var tmpY = 0
		tmpX = origin.X + i
		tmpY = origin.Y + i
		if posIsValid(Position{tmpX, tmpY}) {
			if GameBoard[tmpX][tmpY] == player {
				nbPoint++
			} else {
				nbPoint = 0
			}
			if nbPoint == 4 {
				return true
			}
		}
	}
	for i := 4; i > -4; i++ {
		var tmpX = 0
		var tmpY = 0
		tmpX = origin.X + i
		tmpY = origin.Y - i
		if posIsValid(Position{tmpX, tmpY}) {
			if GameBoard[tmpX][tmpY] == player {
				nbPoint++
			} else {
				nbPoint = 0
			}
			if nbPoint == 4 {
				return true
			}
		}
	}
	return false
}

func computeMapWeight() int64 {
	var weightTotal int64
	var pos Position
	for x := 0; x < len(GameBoard); x++ {
		pos.X = x
		for y := 0; y < len(GameBoard); y++ {
			pos.Y = y
			if (GameBoard[x][y] == 1) {
				weightTotal += int64(calcWeightOfCase(pos, GameBoard[x][y]))
			} else if (GameBoard[x][y] == 2) {
				weightTotal -= int64(calcWeightOfCase(pos, GameBoard[x][y]))
			}
		}
	}
	return weightTotal
}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 2
	}
}

func computeBestPosition(origin Position, deep int, max bool) int64 {
	if deep == 0 {
		return computeMapWeight()
	} else {
		if isWinningPoint(origin, boolToInt(max)) {
			if max {
				return math.MaxInt64
			} else {
				return - math.MaxInt64
			}
		}

		fmt.Printf("%d\n", boolToInt(max))
		GameBoard[origin.X][origin.Y] = boolToInt(max)
		var weights []int64
		var gameBoardLen = len(GameBoard)
		for x := 0; x < gameBoardLen; x++ {
			for y := 0; y < gameBoardLen; y++ {
				if GameBoard[x][y] != 0 {
					var pos = [8]Position{
						{x - 1, y + 0},
						{x - 1, y + 1},
						{x + 0, y + 1},
						{x + 1, y + 1},
						{x + 1, y + 0},
						{x + 1, y - 1},
						{x + 0, y - 1},
						{x - 1, y - 1},
					}

					for i := 0; i < 8; i ++ {
						if posIsAvailable(pos[i]) {
							weights = append(weights, computeBestPosition(pos[i], deep-1, !max))
						}
					}
				}
			}
		}
		GameBoard[origin.X][origin.Y] = 0
		var weight = weights[0]
		if max { // MAXIMUM
			for i := range weights {
				if weights[i] > weight {
					weight = weights[i]
				}
			}
		} else { // MINIMUM
			for i := range weights {
				if weights[i] < weight {
					weight = weights[i]
				}
			}
		}
		return weight
	}
}
