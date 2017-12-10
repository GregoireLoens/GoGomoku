package com

import (
	"bufio"
)

type Game struct {
	start bool
}

type ComStruct struct {
	game    Game
	reader  *bufio.Reader
}
