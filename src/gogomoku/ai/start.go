package ai

import (
	"strconv"
)

var GameBoard [][]int
var HasPlayed = false

type Position struct {
	X int
	Y int
}

const weightNotEmptyPower = 3

func otherPlayer(player int) int {
	if player == 1 {
		return 2
	}
	return 1
}

func posIsValid(pos Position) bool {
	if pos.X >= 0 && pos.X < len(GameBoard) && pos.Y >= 0 && pos.Y < len(GameBoard) {
		return true
	}
	return false
}

func posIsAvailable(pos Position) bool {
	if posIsValid(pos) && GameBoard[pos.X][pos.Y] == 0 {
		return true
	}
	return false
}

func calcWeightOfCase_(pos Position, player int, emptyCase *int, playerCase *int) bool {
	if posIsValid(pos) {
		caseToCheck := GameBoard[pos.X][pos.Y]
		if caseToCheck == player {
			*playerCase++
		} else if caseToCheck == 0 {
			*emptyCase++
		} else {
			return false
		}
		return true
	} else {
		return false
	}
}

type calcWeightOfLineFunc func(a int, b int) Position

func calcWeightOfLine(player int, weightToSet *int, lineFunc calcWeightOfLineFunc) {
	for a := 0; a < 5; a++ {
		var emptyCaseA = 0
		var emptyCaseB = 0
		var playerCaseA = 0
		var playerCaseB = 0
		var continueA = true
		var continueB = true

		for b := 0; b < 5; b++ {

			if !continueA && !continueB {
				break
			}
			if continueA {
				continueA = calcWeightOfCase_(lineFunc(a, b), player, &emptyCaseA, &playerCaseA)
			}
			if continueB {
				continueB = calcWeightOfCase_(lineFunc(-a, -b), player, &emptyCaseB, &playerCaseB)
			}
		}

		if continueA {
			var newWeight int
			var res = weightNotEmptyPower
			if emptyCaseA+playerCaseA > 4 {
				for i := 1; i < playerCaseA; i++ {
					res *= weightNotEmptyPower
				}
				newWeight = 5 - playerCaseA + res
				if *weightToSet < newWeight {
					*weightToSet = newWeight
				}
			}
		}
		if continueB {
			var newWeight int
			var res = weightNotEmptyPower
			if emptyCaseB+playerCaseB > 4 {
				for i := 1; i < playerCaseB; i++ {
					res *= weightNotEmptyPower
				}
				newWeight = 5 - playerCaseB + res
				if *weightToSet < newWeight {
					*weightToSet = newWeight
				}
			}
		}

	}
}

func calcWeightOfCase(origin Position, player int) int {
	var weight = [4]int{0, 0, 0, 0}
	calcWeightOfLine(player, &weight[0], func(a int, b int) Position {
		return Position{origin.X - b + a, origin.Y}
	})
	calcWeightOfLine(player, &weight[1], func(a int, b int) Position {
		return Position{origin.X - b + a, origin.Y + b - a}
	})
	calcWeightOfLine(player, &weight[2], func(a int, b int) Position {
		return Position{origin.X, origin.Y + b - a}
	})
	calcWeightOfLine(player, &weight[3], func(a int, b int) Position {
		return Position{origin.X + b - a, origin.Y + b - a}
	})
	return weight[0] + weight[1] + weight[2] + weight[3]
}

type bestPair struct {
	weight int
	pos    Position
}

func calcBestPositionAndWeight(player int, deep int) bestPair {
	// WEIGHT BOARD CREATION
	boardLen := len(GameBoard)
	weightBoard := make([][]int, boardLen)
	for x := 0; x < boardLen; x++ {
		weightBoard[x] = make([]int, boardLen)
		for y := 0; y < boardLen; y++ {
			weightBoard[x][y] = 0
		}
	}
	// LOOP OVER ALL
	for x := 0; x < boardLen; x++ {
		for y := 0; y < boardLen; y++ {
			if GameBoard[x][y] != 0 {

				// CALC ALL AROUND CASES

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
					if posIsAvailable(pos[i]) && weightBoard[pos[i].X][pos[i].Y] == 0 {
						weightBoard[pos[i].X][pos[i].Y] = calcWeightOfCase(pos[i], player) + calcWeightOfCase(pos[i], otherPlayer(player))
					}
				}
			}
		}
	}

	if deep == 0 {
		if player == 1 { // MAXIMUM
			var bestPosition = Position{0, 0}
			var bestWeight = -1
			for x := 0; x < boardLen; x++ {
				for y := 0; y < boardLen; y++ {
					if weightBoard[x][y] > bestWeight {
						bestWeight = weightBoard[x][y]
						bestPosition = Position{x, y}
					}
				}
			}
			return bestPair{bestWeight, bestPosition}
		} else { // MINIMUM
			var bestPosition = Position{0, 0}
			var bestWeight = 1000000
			for x := 0; x < boardLen; x++ {
				for y := 0; y < boardLen; y++ {
					if weightBoard[x][y] < bestWeight && weightBoard[x][y] > 0 {
						bestWeight = weightBoard[x][y]
						bestPosition = Position{x, y}
					}
				}
			}
			return bestPair{bestWeight, bestPosition}
		}
	} else {
		var bestPairTab []bestPair
		for x := 0; x < boardLen; x++ {
			for y := 0; y < boardLen; y++ {
				if weightBoard[x][y] > 0 {
					GameBoard[x][y] = player
					best := calcBestPositionAndWeight(otherPlayer(player), deep-1)
					best.pos = Position{x, y}
					bestPairTab = append(bestPairTab, best)
					GameBoard[x][y] = 0
				}
			}
		}
		if player == 1 { // MINIMUM
			var bestPair = bestPair{1000000, Position{0, 0}}
			for i := range bestPairTab {
				if bestPairTab[i].weight < bestPair.weight && bestPairTab[i].weight > 0 {
					bestPair = bestPairTab[i]
				}
			}
			return bestPair
		} else { // MAXIMUM
			var bestPair = bestPair{-1, Position{0, 0}}
			for i := range bestPairTab {
				if bestPairTab[i].weight > bestPair.weight {
					bestPair = bestPairTab[i]
				}
			}
			return bestPair
		}
	}
}

const maxDeep = 2

func turn() Position {
	if !HasPlayed { // If first turn
		return Position{len(GameBoard) / 2, len(GameBoard) / 2}
	}

	return calcBestPositionAndWeight(1, maxDeep).pos
}

func returnChan(comChan chan<- string, x int, y int) {
	sX := strconv.Itoa(x)
	sY := strconv.Itoa(y)
	comChan <- sX + "," + sY
	GameBoard[x][y] = 1
}

func Start(comChan chan<- string) {
	var pos = turn()
	returnChan(comChan, pos.X, pos.Y)
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
