package ai

import (
	"strconv"
)

var GameBoard [][]int

type Position struct {
	X int
	Y int
}

var LastEnemyPosition = Position{-1, -1}
var LastPlayerPosition = Position{-1, -1}

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
	if pos.X > 0 && pos.X < len(GameBoard) && pos.Y > 0 && pos.Y < len(GameBoard) {
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

func calcWeightByStarPart(pos Position, player int, weightToSet *int) bool {
	if posIsValid(pos) {
		if GameBoard[pos.X][pos.Y] == player {
			*weightToSet++
			return true
		} else if GameBoard[pos.X][pos.Y] == otherPlayer(player) {
			return false
		}
		return true
	}
	return false
}

func calcWeightByStar(position Position, turn int) int {
	var weight = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X - i, position.Y}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[0])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X - i, position.Y + i}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[1])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X, position.Y + i}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[2])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X + i, position.Y + i}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[3])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X + i, position.Y}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[4])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X + i, position.Y - i}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[5])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X, position.Y - i}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[6])
	func(weightToSet *int) {
		for i := 1; i < 5; i++ {
			newPos := Position{position.X - i, position.Y - i}
			if calcWeightByStarPart(newPos, turn, weightToSet) == false {
				break
			}
		}
	}(&weight[7])
	return weight[0] ^ 4 + weight[1] ^ 4 + weight[2] ^ 4 + weight[3] ^ 4 +
		weight[4] ^ 4 + weight[5] ^ 4 + weight[6] ^ 4 + weight[7] ^ 4
}

// SI ON CHANGE MAX PROFONDEUR, IL FAUT LA CHANGER ICI
func calcBestPosition(pos Position, profondeur int) int {
	if GameBoard[pos.X][pos.Y] == 0 {
		return calcWeightByStar(pos, 1) ^ profondeur + calcWeightByStar(pos, 2) ^ profondeur + calcAllBestPosition(pos, profondeur-1)
	}
	return 0
}

func calcAllBestPosition(origin Position, profondeur int) int {
	if profondeur == 0 {
		return 0
	}

	var weight = [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
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

	if posIsAvailable(pos[0]) {
		weight[0] = calcBestPosition(pos[0], profondeur)
	}
	if posIsAvailable(pos[1]) {
		weight[1] = calcBestPosition(pos[1], profondeur)
	}
	if posIsAvailable(pos[2]) {
		weight[2] = calcBestPosition(pos[2], profondeur)
	}
	if posIsAvailable(pos[3]) {
		weight[3] = calcBestPosition(pos[3], profondeur)
	}
	if posIsAvailable(pos[4]) {
		weight[4] = calcBestPosition(pos[4], profondeur)
	}
	if posIsAvailable(pos[5]) {
		weight[5] = calcBestPosition(pos[5], profondeur)
	}
	if posIsAvailable(pos[6]) {
		weight[6] = calcBestPosition(pos[6], profondeur)
	}
	if posIsAvailable(pos[7]) {
		weight[7] = calcBestPosition(pos[7], profondeur)
	}

	var max = 0
	for i := 1; i < 8; i++ {
		if weight[i] > weight[max] {
			max = i
		}
	}
	return weight[max]
}

func bestPosition(origin Position, profondeur int) Position {
	var weight = [8]int{-1, -1, -1, -1, -1, -1, -1, -1}
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
	if posIsAvailable(pos[0]) {
		weight[0] = calcAllBestPosition(pos[0], profondeur)
	}
	if posIsAvailable(pos[1]) {
		weight[1] = calcAllBestPosition(pos[1], profondeur)
	}
	if posIsAvailable(pos[2]) {
		weight[2] = calcAllBestPosition(pos[2], profondeur)
	}
	if posIsAvailable(pos[3]) {
		weight[3] = calcAllBestPosition(pos[3], profondeur)
	}
	if posIsAvailable(pos[4]) {
		weight[4] = calcAllBestPosition(pos[4], profondeur)
	}
	if posIsAvailable(pos[5]) {
		weight[5] = calcAllBestPosition(pos[5], profondeur)
	}
	if posIsAvailable(pos[6]) {
		weight[6] = calcAllBestPosition(pos[6], profondeur)
	}
	if posIsAvailable(pos[7]) {
		weight[7] = calcAllBestPosition(pos[7], profondeur)
	}
	var max = 0
	for i := 1; i < 8; i++ {
		if weight[i] > weight[max] {
			max = i
		}
	}
	return pos[max]
}

func firstTurn() Position {
	return Position{len(GameBoard) / 2, len(GameBoard) / 2}
}

func turn() Position {
	var enemyWeight = -1
	var playerWeight = -1
	if hasPlayed(LastPlayerPosition) {
		playerWeight = calcWeightByStar(LastPlayerPosition, 1)
	}
	if hasPlayed(LastEnemyPosition) {
		enemyWeight = calcWeightByStar(LastEnemyPosition, 2)
	}
	if enemyWeight == -1 && playerWeight == -1 {
		return firstTurn()
	}
	if playerWeight > enemyWeight {
		return bestPosition(LastPlayerPosition, 2)
	} else {
		return bestPosition(LastEnemyPosition, 2)
	}
}

func returnChan(comChan chan<- string, x int, y int) {
	GameBoard[x][y] = 2
	sX := strconv.Itoa(x + 1)
	sY := strconv.Itoa(y + 1)
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
