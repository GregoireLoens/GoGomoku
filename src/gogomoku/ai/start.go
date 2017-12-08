package ai

import (
	"strconv"
)

var GameBoard [][]int
var WeightGameBoard [2][][]int

type Position struct {
	X int
	Y int
}

var LastEnemyPosition = Position{-1, -1}
var LastPlayerPosition = Position{-1, -1}

const weightNotEmptyPower = 3

func hasPlayed(pos Position) bool {
	return pos.X != -1 && pos.Y != -1
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

type posWeightStruct struct {
	weight int64
	pos    Position
}

func turn() Position {
	var tab []posWeightStruct
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

				for i := 0; i < 8; i++ {
					if posIsAvailable(pos[i]) {
						weight := computeBestPosition(pos[i], 1, 1)
						debugMessage("Pos X is " + strconv.Itoa(pos[i].X) + " pos y is " + strconv.Itoa(pos[i].Y) + " the weight is " + strconv.Itoa(int(weight)))
						tab = append(tab, posWeightStruct{weight, pos[i]})
					}
				}
			}
		}
	}
	maxTab := tab[0]
	for i := range tab {
		if tab[i].weight > maxTab.weight {
			maxTab = tab[i]
		}
	}
	return maxTab.pos
}

func returnChan(comChan chan<- string, x int, y int) {
	sX := strconv.Itoa(x)
	sY := strconv.Itoa(y)
	comChan <- sX + "," + sY
	LastPlayerPosition.X = x
	LastPlayerPosition.Y = y
	GameBoard[x][y] = 1
}

func Start(comChan chan<- string) {
	var pos Position
	if !hasPlayed(LastEnemyPosition) && !hasPlayed(LastPlayerPosition) {
		pos = Position{len(GameBoard) / 2, len(GameBoard) / 2}
	} else {
		pos = turn()
	}
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
