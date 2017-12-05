package ai

import (
	"strconv"
)

var GameBoard [][]int
var WeightGameBoard [][]int

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
		var emptyCase = 1
		var playerCase = 0
		var continue1 = true
		var continue2 = true
		for i := 1; i < 5; i++ {
			if !continue1 && !continue2 {
				break
			}
			if continue1 {
				continue1 = calcWeightOfCase_(Position{origin.X - i, origin.Y}, player, &emptyCase, &playerCase)
			}
			if continue2 {
				continue2 = calcWeightOfCase_(Position{origin.X + i, origin.Y}, player, &emptyCase, &playerCase)
			}
		}
		var res = 3
		if emptyCase + playerCase > 4 {
			for i:= 1; i < playerCase; i++{
				res *= 3
			}
			weight[0] = 5 - playerCase + res
		}
	}()
	func () {
		var emptyCase = 1
		var playerCase = 0
		var continue1 = true
		var continue2 = true
		for i := 1; i < 5; i++ {
			if !continue1 && !continue2 {
				break
			}
			if continue1 {
				continue1 = calcWeightOfCase_(Position{origin.X - i, origin.Y + i}, player, &emptyCase, &playerCase)
			}
			if continue2 {
				continue2 = calcWeightOfCase_(Position{origin.X + i, origin.Y - i}, player, &emptyCase, &playerCase)
			}
		}
		var res = 3
		if emptyCase + playerCase > 4 {
			for i:= 1; i < playerCase; i++{
				res *= 3
			}
			weight[1] = 5 - playerCase + res
		}
	}()
	func () {
		var emptyCase = 1
		var playerCase = 0
		var continue1 = true
		var continue2 = true
		for i := 1; i < 5; i++ {
			if !continue1 && !continue2 {
				break
			}
			if continue1 {
				continue1 = calcWeightOfCase_(Position{origin.X, origin.Y + i}, player, &emptyCase, &playerCase)
			}
			if continue2 {
				continue2 = calcWeightOfCase_(Position{origin.X, origin.Y - i}, player, &emptyCase, &playerCase)
			}
		}
		var res = 3
		if emptyCase + playerCase > 4 {
			for i:= 1; i < playerCase; i++{
				res *= 3
			}
			weight[2] = 5 - playerCase + res
		}
	}()
	func () {
		var emptyCase = 1
		var playerCase = 0
		var continue1 = true
		var continue2 = true
		for i := 1; i < 5; i++ {
			if !continue1 && !continue2 {
				break
			}
			if continue1 {
				continue1 = calcWeightOfCase_(Position{origin.X + i, origin.Y + i}, player, &emptyCase, &playerCase)
			}
			if continue2 {
				continue2 = calcWeightOfCase_(Position{origin.X - i, origin.Y - i}, player, &emptyCase, &playerCase)
			}
		}
		var res = 3
		if emptyCase + playerCase > 4 {
			for i:= 1; i < playerCase; i++{
				res *= 3
			}
			weight[3] = 5 - playerCase + res
		}
	}()
	return weight
}

func calcAllWeightOfCase(bestPosition *Position, bestWeight *[4]int, origin Position, player int) {
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
			if posIsAvailable(pos[i]) && WeightGameBoard[pos[i].X][pos[i].Y] == 0 {
				weight[i] = calcWeightOfCase(pos[i], player)
			}
		}
		for i := range weight {
			if weight[i][0] > (*bestWeight)[0] ||
				weight[i][1] > (*bestWeight)[1] || weight[i][2] > (*bestWeight)[2] || weight[i][3] > (*bestWeight)[3] {
				*bestWeight = weight[i]
				*bestPosition = pos[i]
			}
			if posIsAvailable(pos[i]) {
				WeightGameBoard[pos[i].X][pos[i].Y] = 1
			}
		}
		}

func calcBestPositionAndWeight(bestPosition *Position, bestWeight *[4]int, player int) {
	for x := range GameBoard {
		for y := range GameBoard[x] {
			if GameBoard[x][y] == player {
				calcAllWeightOfCase(bestPosition, bestWeight, Position{x, y}, player)
			}
		}
	}
}

func turn() Position {
	if !hasPlayed(LastPlayerPosition) && !hasPlayed(LastEnemyPosition) { // If first turn
		return Position{len(GameBoard) / 2, len(GameBoard) / 2}
	}

	var bestPlayerPosition = Position{0, 0}
	var bestPlayerWeight = [4]int{-1, -1, -1, -1}
	var bestEnemyPosition = Position{0, 0}
	var bestEnemyWeight = [4]int{-1, -1, -1, -1}

	for x := range WeightGameBoard {
		for y := range WeightGameBoard[x] {
			WeightGameBoard[x][y] = 0
		}
	}
	calcBestPositionAndWeight(&bestPlayerPosition, &bestPlayerWeight, 1)
	/*t := strconv.Itoa(bestPlayerWeight[0])
	debugMessage(t + " pos 0")
	u := strconv.Itoa(bestPlayerWeight[1])
	debugMessage(u + " pos 1")
	v := strconv.Itoa(bestPlayerWeight[2])
	debugMessage(v + " pos 2")
	w := strconv.Itoa(bestPlayerWeight[3])
	debugMessage(w + " pos 3")

	weight:= strconv.Itoa(weightAlarm)
	debugMessage("Weight alarm is " + weight)*/
	if bestPlayerWeight[0] == weightAlarm || bestPlayerWeight[1] == weightAlarm || bestPlayerWeight[2] == weightAlarm || bestPlayerWeight[3] == weightAlarm {
		return bestPlayerPosition
	}

	for x := range WeightGameBoard {
		for y := range WeightGameBoard[x] {
			WeightGameBoard[x][y] = 0
		}
	}
	calcBestPositionAndWeight(&bestEnemyPosition, &bestEnemyWeight, 2)

	/*a := strconv.Itoa(bestEnemyWeight[0])
	debugMessage(a + " ennemy pos 0")
	b := strconv.Itoa(bestEnemyWeight[1])
	debugMessage(b + " ennemy pos 1")
	c := strconv.Itoa(bestEnemyWeight[2])
	debugMessage(c + " ennemy pos 2")
	d := strconv.Itoa(bestEnemyWeight[3])
	debugMessage(d + " ennemy pos 3")*/

	if bestEnemyWeight[0] == weightAlarm || bestEnemyWeight[1] == weightAlarm || bestEnemyWeight[2] == weightAlarm || bestEnemyWeight[3] == weightAlarm {
		return bestEnemyPosition
	} else if bestPlayerWeight[0] == weightWarning || bestPlayerWeight[1] == weightWarning || bestPlayerWeight[2] == weightWarning || bestPlayerWeight[3] == weightWarning {
		return bestPlayerPosition
	} else if bestEnemyWeight[0] == weightWarning || bestEnemyWeight[1] == weightWarning || bestEnemyWeight[2] == weightWarning || bestEnemyWeight[3] == weightWarning {
		return bestEnemyPosition
	}

	var bestSumPlayerWeight = bestPlayerWeight[0] + bestPlayerWeight[1] + bestPlayerWeight[2] + bestPlayerWeight[3]
	var bestSumEnemyWeight = bestEnemyWeight[0] + bestEnemyWeight[1] + bestEnemyWeight[2] + bestEnemyWeight[3]

	/*g := strconv.Itoa(bestSumPlayerWeight)
	h := strconv.Itoa(bestSumEnemyWeight)
	debugMessage(g + " contre " + h)*/
	if bestSumEnemyWeight > bestSumPlayerWeight {
		return bestEnemyPosition
	}
	return bestPlayerPosition
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
