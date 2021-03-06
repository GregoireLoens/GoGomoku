package ai

import (
	"math"
	"container/list"
)

func isWinningPoint(origin Position, player int) bool {
	var nbPoint = 0

	for i := 4; i > -4; i-- {
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
	for i := 4; i > -4; i-- {
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
	for i := 4; i > -4; i-- {
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
	for i := 4; i > -4; i-- {
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

func computeMapWeight(player int) int64 {
	var weightTotal int64 = 0
	var pos Position
	for pos.X = 0; pos.X < len(GameBoard); pos.X++ {
		for pos.Y = 0; pos.Y < len(GameBoard); pos.Y++ {
			if GameBoard[pos.X][pos.Y] == player {
				weightTotal += int64(calcWeightOfCase(pos, player))
			} else if GameBoard[pos.X][pos.Y] == otherPlayer(player) {
				weightTotal -= int64(calcWeightOfCase(pos, otherPlayer(player)))
			}
		}
	}
	return weightTotal
}

func otherPlayer(p int) int {
	if p == 2 {
		return 1
	} else {
		return 2
	}
}

func computeBestPosition(origin Position, deep int, player int) int64 {
	if deep == 0 {
		return computeMapWeight(player)
	} else {
		if isWinningPoint(origin, player) {
			if player == 1 {
				return math.MaxInt64
			} else {
				return -math.MaxInt64
			}
		}

		var listPosition *list.List = list.New()

		GameBoard[origin.X][origin.Y] = player
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
							if !Any(listPosition, pos[i]) {
								listPosition.PushBack(pos[i])
							}
						}
					}
				}
			}
		}
		for e := listPosition.Front(); e != nil; e = e.Next() {
			weight := computeBestPosition(e.Value.(Position), deep-1, otherPlayer(player))
			weights = append(weights, weight)
			//debugMessage("Pos (X=" + strconv.Itoa(e.Value.(Position).X) + ";Y=" + strconv.Itoa(e.Value.(Position).Y) + ") : " + strconv.Itoa(int(weight)))
		}
		GameBoard[origin.X][origin.Y] = 0
		var weight = weights[0]
		if player == 1 { // MAXIMUM
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
