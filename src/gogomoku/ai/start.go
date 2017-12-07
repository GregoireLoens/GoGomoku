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

const weightNotEmptyPower = 4

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
	//return int(math.Max(float64(weight[0]), math.Max(float64(weight[1]), math.Max(float64(weight[2]), float64(weight[3])))))
}

type bestPair struct {
	weight int
	pos    Position
}

func calcBestPositionAndWeight(player int, deep int) bestPair {
	boardLen := len(GameBoard)
	var weightTab []bestPair
	// LOOP OVER MAP
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
					if posIsAvailable(pos[i]) {
						val := calcWeightOfCase(pos[i], player) + calcWeightOfCase(pos[i], otherPlayer(player))
						weightTab = append(weightTab, bestPair{val, pos[i]})
					}
				}
			}
		}
	}

	if deep == 0 {
		debugMessage(strconv.Itoa(deep) + " : MAX DEEP")
		var best = weightTab[0]
		if player == 1 { // MAXIMUM
			for i := range weightTab {
				if weightTab[i].weight > best.weight {
					best = weightTab[i]
				}
			}
			debugMessage(strconv.Itoa(deep) + " : MAX DEEP RET (MAX) " + strconv.Itoa(best.weight))
		} else { // MINIMUM
			for i := range weightTab {
				if weightTab[i].weight < best.weight {
					best = weightTab[i]
				}
			}
			debugMessage(strconv.Itoa(deep) + " : MAX DEEP RET (MIN) " + strconv.Itoa(best.weight))
		}
		return best
	} else {
		debugMessage(strconv.Itoa(deep) + " : DIVE")
		var bestPairTab []bestPair
		var best = weightTab[0]
		for i := range weightTab {
			GameBoard[weightTab[i].pos.X][weightTab[i].pos.X] = player
			best := calcBestPositionAndWeight(otherPlayer(player), deep-1)
			best.pos = weightTab[i].pos
			bestPairTab = append(bestPairTab, best)
			GameBoard[weightTab[i].pos.X][weightTab[i].pos.X] = 0
		}
		best = bestPairTab[0]
		if player == 1 { // MINIMUM
			for i := range bestPairTab {
				if bestPairTab[i].weight < best.weight {
					debugMessage(strconv.Itoa(deep) + " : DIVE GET " + strconv.Itoa(bestPairTab[i].weight))
					best = bestPairTab[i]
				}
			}
			debugMessage(strconv.Itoa(deep) + " : DIVE RET (MIN) " + strconv.Itoa(best.weight))
		} else { // MAXIMUM
			for i := range bestPairTab {
				debugMessage(strconv.Itoa(deep) + " : DIVE GET " + strconv.Itoa(bestPairTab[i].weight))
				if bestPairTab[i].weight > best.weight {
					best = bestPairTab[i]
				}
			}
			debugMessage(strconv.Itoa(deep) + " : DIVE RET (MAX) " + strconv.Itoa(best.weight))
		}
		return best
	}
}

func returnChan(comChan chan<- string, x int, y int) {
	sX := strconv.Itoa(x)
	sY := strconv.Itoa(y)
	comChan <- sX + "," + sY
	GameBoard[x][y] = 1
}

const maxDeep = 0

func Start(comChan chan<- string) {
	var pos Position
	if !HasPlayed { // If first turn
		pos = Position{len(GameBoard) / 2, len(GameBoard) / 2}
	} else {
		pos = calcBestPositionAndWeight(1, maxDeep).pos
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
