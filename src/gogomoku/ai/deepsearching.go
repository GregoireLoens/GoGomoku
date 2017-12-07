package ai

func isWinningPoint(origin Position) {

}

func computeMapWeight() int {

}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 2
	}
}

func computeBestPosition(pos Position, deep int, max bool) int {
	if deep == 0 {
		return computeMapWeight()
	} else {
		GameBoard[pos.X][pos.Y] = boolToInt(max)
		var weights []int
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
		GameBoard[pos.X][pos.Y] = 0
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
