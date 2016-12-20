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

	waitLevel := 0

	for keepAtIt {
		t := time.NewTimer(300 * time.Millisecond)
		select {
		case s := <-in.Lines:
			fmt.Printf("%s\n", s)
			waitLevel = 0
		case <-t.C:
			waitLevel++
			if waitLevel < 3 {
				fmt.Printf("\n\n")
			}
		case <-in.Done:
			keepAtIt = false
		}
		t.Stop()
	}
}
