package main

import (
	"fmt"

	"../../internal/game"
)

func main() {
	defer func() {
		err := recover()
		if e, ok := err.(*game.ReaderError); ok {
			fmt.Println(e)
			fmt.Println("Trying to start new game")
			runGame()
		}

		if err != nil {
			fmt.Println(err)
			panic(err)
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
