package com

import (
	"gogomoku/ai"
	"os"
	"bufio"
	"sync"
	"fmt"
	"time"
)

func mainRoutine(msg string) {
	wg := new(sync.WaitGroup)
	comChan := make(chan string, 1)

	wg.Add(2)
	go func() {
		defer wg.Done()

		comChan <- msg
		ai.Start(comChan)
	}()
	time.Sleep(time.Millisecond * 5)
	go func() {
		defer wg.Done()

		select {
		case res := <- comChan:
			fmt.Println("On recoit " + res)
		case <- time.After(time.Second * 3):
			fmt.Println("Timeout 3")
		}
	}()
	wg.Wait()
	close(comChan)
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

			mainRoutine(msg)
		}(msg)
	}
	wg.Wait()
}
