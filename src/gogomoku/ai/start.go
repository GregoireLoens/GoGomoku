package ai

import (
	"strconv"
	"sync"
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
var warningStack[] Position

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
			if emptyCaseA + playerCaseA > 4 {
				for i:= 1; i < playerCaseA; i++{
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
			if emptyCaseB + playerCaseB > 4 {
				for i:= 1; i < playerCaseB; i++{
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

func calcWeightOfCase(origin Position, player int) [4]int {
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
	return weight
}

func calcAllWeightOfCase(bestPosition *Position, bestWeight *int, origin Position, tab int, player int) {
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
	wg := new(sync.WaitGroup)

	wg.Add(8)
	for i := 0; i < 8; i ++ {
		go func(index int) {
			defer wg.Done()
			var weight [4]int
			if posIsAvailable(pos[index]) && WeightGameBoard[tab][pos[index].X][pos[index].Y] == 0 {
				weight = calcWeightOfCase(pos[index], player)
				for j := 0; j < 4; j++ {
					if weight[j] > *bestWeight {
						*bestWeight = weight[j]
						*bestPosition = pos[index]
						warningStack = append(warningStack, pos[index])
					}
				}
				WeightGameBoard[tab][pos[index].X][pos[index].Y] = weight[0] + weight[1] + weight[2] + weight[3]
			}
		}(i)
	}
	wg.Wait()
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

	for x := range WeightGameBoard[0] {
		for y := range WeightGameBoard[0][x] {
			if WeightGameBoard[0][x][y] + WeightGameBoard[1][x][y] > bestWeight {
				bestWeight = WeightGameBoard[0][x][y] + WeightGameBoard[1][x][y]
				bestPosition = Position{x, y}
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
