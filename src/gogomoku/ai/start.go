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
const weightAlarm = 82
const weightWarning = 29

func hasPlayed(pos Position) bool {
	return pos.X != -1 && pos.Y != -1
}

func otherPlayer(player int) int {
	if player == 1 {
		return 2
	} else {
		return 1
	}
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

func calcWeightOfCase(origin Position, player int) [4]int {
	var weight = [4]int{-1, -1, -1, -1}
	func () {
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
					continueA = calcWeightOfCase_(Position{origin.X - b + a, origin.Y}, player, &emptyCaseA, &playerCaseA)
				}
				if continueB {
					continueB = calcWeightOfCase_(Position{origin.X + b - a, origin.Y}, player, &emptyCaseB, &playerCaseB)
				}
			}

			if continueA {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseA + playerCaseA > 4 {
					for i:= 1; i < playerCaseA; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseA + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
			if continueB {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseB + playerCaseB > 4 {
					for i:= 1; i < playerCaseB; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseB + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}

		}
	}()
	func () {
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
					continueA = calcWeightOfCase_(Position{origin.X - b + a, origin.Y + b - a}, player, &emptyCaseA, &playerCaseA)
				}
				if continueB {
					continueB = calcWeightOfCase_(Position{origin.X + b - a, origin.Y - b + a}, player, &emptyCaseB, &playerCaseB)
				}
			}

			if continueA {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseA + playerCaseA > 4 {
					for i:= 1; i < playerCaseA; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseA + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
			if continueB {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseB + playerCaseB > 4 {
					for i:= 1; i < playerCaseB; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseB + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
		}
	}()
	func () {
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
					continueA = calcWeightOfCase_(Position{origin.X, origin.Y + b - a}, player, &emptyCaseA, &playerCaseA)
				}
				if continueB {
					continueB = calcWeightOfCase_(Position{origin.X, origin.Y - b + a}, player, &emptyCaseB, &playerCaseB)
				}
			}

			if continueA {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseA + playerCaseA > 4 {
					for i:= 1; i < playerCaseA; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseA + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
			if continueB {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseB + playerCaseB > 4 {
					for i:= 1; i < playerCaseB; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseB + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
		}
	}()
	func () {
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
					continueA = calcWeightOfCase_(Position{origin.X + b + a, origin.Y + b - a}, player, &emptyCaseA, &playerCaseA)
				}
				if continueB {
					continueB = calcWeightOfCase_(Position{origin.X - b - a, origin.Y - b + a}, player, &emptyCaseB, &playerCaseB)
				}
			}

			if continueA {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseA + playerCaseA > 4 {
					for i:= 1; i < playerCaseA; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseA + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
			if continueB {
				var newWeight int
				var res = weightNotEmptyPower
				if emptyCaseB + playerCaseB > 4 {
					for i:= 1; i < playerCaseB; i++{
						res *= weightNotEmptyPower
					}
					newWeight = 5 - playerCaseB + res
					if weight[0] < newWeight {
						weight[0] = newWeight
					}
				}
			}
		}
	}()
	return weight
}

func calcAllWeightOfCase(bestPosition *Position, bestWeight *int, origin Position, tab int, player int) {
	var weight = [8][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var pos = [8]Position{
		{origin.X - 1, origin.Y},
		{origin.X - 1, origin.Y + 1},
		{origin.X, origin.Y + 1},
		{origin.X + 1, origin.Y + 1},
		{origin.X + 1, origin.Y},
		{origin.X + 1, origin.Y - 1},
		{origin.X, origin.Y - 1},
		{origin.X - 1, origin.Y - 1},
	}

	for i := 0; i < 8; i ++ {
		if posIsAvailable(pos[i]) && WeightGameBoard[tab][pos[i].X][pos[i].Y] == 0 {
			weight[i] = calcWeightOfCase(pos[i], player)
		}
	}
	for i := range weight {
		for j := 0; j < 4; j++ {
			if weight[i][j] > *bestWeight {
				*bestWeight = weight[i][j]
				*bestPosition = pos[i]
			}
		}
		if posIsAvailable(pos[i]) {
			WeightGameBoard[tab][pos[i].X][pos[i].Y] = weight[i][0] + weight[i][1] + weight[i][2] + weight[i][3]
		}
	}
}

func calcBestPositionAndWeight(bestPosition *Position, bestWeight *int, tab int, player int) {
	for x := range GameBoard {
		for y := range GameBoard[x] {
			if GameBoard[x][y] == player {
				calcAllWeightOfCase(bestPosition, bestWeight, Position{x, y}, tab, player)
			}
		}
	}
}

func bestPositionInWeightBoard() Position {
	var bestWeight = -1
	var bestPosition Position

	for a := 0; a < 2; a++ {
		for x := range WeightGameBoard[a] {
			for y := range WeightGameBoard[a][x] {
				if WeightGameBoard[a][x][y] > bestWeight {
					bestWeight = WeightGameBoard[a][x][y]
					bestPosition = Position{x, y}
				}
			}
		}
	}
	return bestPosition
}

func turn() Position {
	if !hasPlayed(LastPlayerPosition) && !hasPlayed(LastEnemyPosition) { // If first turn
		return Position{len(GameBoard) / 2, len(GameBoard) / 2}
	}

	var bestPlayerPosition = Position{0, 0}
	var bestPlayerWeight = -1
	var bestEnemyPosition = Position{0, 0}
	var bestEnemyWeight = -1

	for a := 0; a < 2; a ++ {
		for x := range WeightGameBoard[a] {
			for y := range WeightGameBoard[a][x] {
				WeightGameBoard[a][x][y] = 0
			}
		}
	}
	calcBestPositionAndWeight(&bestPlayerPosition, &bestPlayerWeight, 0, 1)

	if bestPlayerWeight == weightAlarm  {
		return bestPlayerPosition
	}

	for a := 0; a < 2; a ++ {
		for x := range WeightGameBoard[a] {
			for y := range WeightGameBoard[a][x] {
				WeightGameBoard[a][x][y] = 0
			}
		}
	}
	calcBestPositionAndWeight(&bestEnemyPosition, &bestEnemyWeight,1, 2)

	if bestEnemyWeight == weightAlarm {
		return bestEnemyPosition
	} else if bestPlayerWeight == weightWarning {
		return bestPlayerPosition
	} else if bestEnemyWeight == weightWarning {
		return bestEnemyPosition
	}

	return bestPositionInWeightBoard()
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
