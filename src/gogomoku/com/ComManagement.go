package com

import (
	"gogomoku/ai"
	"os"
	"bufio"
	"sync"
	"fmt"
	"time"
	"regexp"
)

type ComFunc func(string)

type ComFuncTab struct {
	fun		ComFunc
	reg		string
}

var comFuncTab = [1]ComFuncTab{
	{ fun: launchAI, reg: "BEGIN" },
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
		case <- time.After(time.Second * 3):
			fmt.Println("Timeout 3")
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
	wg := new(sync.WaitGroup)
	com := new(ComStruct)
	com.reader = bufio.NewReader(os.Stdin)

	for true {
		msg, err := com.reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()

			parseCom(msg)
		}(msg)
	}
	wg.Wait()
}
