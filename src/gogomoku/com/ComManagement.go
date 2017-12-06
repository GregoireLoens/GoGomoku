package com

import (
	"gogomoku/ai"
	"os"
	"bufio"
	"sync"
	"fmt"
	"time"
	"regexp"
	"strconv"
	"container/list"
)

type ComFunc func(string)

type ComFuncTab struct {
	fun		ComFunc
	reg		string
}

var comFuncTab = [7]ComFuncTab{
	{ fun: restartGame, reg: "RESTART"},
	{ fun: launchAI, reg: "BEGIN" },
	{ fun: startGame, reg: "^START" },
	{ fun: enemyTurn, reg: "TURN" },
	{ fun: endGame, reg: "END" },
	{ fun: aboutAI, reg: "ABOUT"},
	{ fun: board, reg: "BOARD"},
}

var isActive bool = true

func board(com string) {
	var board = new(list.List)
	reader :=  bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	for tmp != "DONE" {
		board.PushBack(tmp)
	}
	fmt.Println(board)
}

func aboutAI(com string) {
	fmt.Println("name=GoGomoku, version=1.0, author=SaltTeam, country=France")
}

func endGame(_ string) {
	isActive = false
}

func restartGame(_ string) {
	for line := range ai.GameBoard {
		for section := range  ai.GameBoard[line] {
			ai.GameBoard[line][section] = 0
		}
	}
	ai.LastEnemyPosition.X = -1
	ai.LastEnemyPosition.Y = -1
	ai.LastPlayerPosition.X = -1
	ai.LastPlayerPosition.Y = -1
	fmt.Println("OK")
}

func enemyTurn(com string)  {
	r, err := regexp.Compile("TURN ([0-9]+),([0-9]+)")
	if err != nil {
		fmt.Println(err)
	}
	s := r.FindStringSubmatch(com)
	x, err := strconv.Atoi(s[1])
	if err != nil {
		fmt.Println(err)
	}
	y, err := strconv.Atoi(s[2])
	if err != nil {
		fmt.Println(err)
	}

	ai.GameBoard[x][y] = 2

	ai.LastEnemyPosition.X = x
	ai.LastEnemyPosition.Y = y

	launchAI(com)
}

func startGame(com string) {
	r, err := regexp.Compile("START ([0-9]+)")
	if err != nil {
		fmt.Println(err)
	}
	size, err := strconv.Atoi(r.FindStringSubmatch(com)[1])
	if err != nil {
		fmt.Println(err)
	}
	ai.GameBoard = make([][]int, size)
	for x := range ai.GameBoard {
		ai.GameBoard[x] = make([]int, size)
		for y := range ai.GameBoard[x] {
			ai.GameBoard[x][y] = 0
		}
	}

	ai.WeightGameBoard = make([][]int, size)
	for x := range ai.WeightGameBoard {
		ai.WeightGameBoard[x] = make([]int, size)
		for y := range ai.WeightGameBoard[x] {
			ai.WeightGameBoard[x][y] = 0
		}
	}

	fmt.Println("OK")
}

func launchAI(_ string) {
	wg := new(sync.WaitGroup)
	comChan := make(chan string, 1)

	wg.Add(2)
	go func() {
		defer wg.Done()

		ai.Start(comChan)
	}()
	time.Sleep(time.Millisecond * 5)
	go func() {
		defer wg.Done()

		select {
		case res := <- comChan:
			fmt.Println(res)
		case <- time.After(time.Second * 1):
			ai.StartRandom(comChan)
			fmt.Println(<- comChan)
		}
	}()
	wg.Wait()
	close(comChan)
}

func parseCom(com string) {
	for _, elem := range comFuncTab {
		match, err := regexp.Match(elem.reg, []byte(com))
		if err != nil {
			fmt.Print(err)
		} else if match {
			elem.fun(com)
		}
	}
}

func ComManagement() {
	com := new(ComStruct)
	com.reader = bufio.NewReader(os.Stdin)

	isActive = true

	for isActive {
		msg, err := com.reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		parseCom(msg)
	}
}
