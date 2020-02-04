package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Input struct {
	Lines chan string
	Done  chan bool
}

func NewInput() *Input {
	return &Input{
		Lines: make(chan string),
		Done:  make(chan bool),
	}
}

func (i *Input) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i.Lines <- scanner.Text()
	}
	i.Done <- true
}

func main() {
	keepAtIt := true
	in := NewInput()
	go in.Start()
	lastPrint := time.Now()

	for keepAtIt {
		select {
		case s := <-in.Lines:
			since := time.Since(lastPrint)
			switch {
			case since > 600*time.Millisecond:
				fmt.Printf("\n\n\n\n")
			case since > 300*time.Millisecond:
				fmt.Printf("\n\n")
			default:
			}
			fmt.Printf("%s\n", s)
			lastPrint = time.Now()
		case <-in.Done:
			keepAtIt = false
		}
	}
}
