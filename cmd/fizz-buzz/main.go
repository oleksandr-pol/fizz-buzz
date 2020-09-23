package main

import (
	"fmt"

	"../../internal/game"
)

func main() {
	defer func() {
		r := recover()
		if r == "Failed to read input" {
			fmt.Println(r)
			fmt.Println("Trying to start new game")
			runGame()
		}

		if r != nil {
			fmt.Println(r)
			panic(r)
		}
	}()

	runGame()
}

func runGame() {
	play := game.Init()
	res, exit, err := play()
	if err != nil {
		fmt.Println(err)
		return
	}

	if exit {
		return
	}

	showResult(res...)
}

func showResult(values ...string) {
	for _, val := range values {
		fmt.Println(val)
	}
}
