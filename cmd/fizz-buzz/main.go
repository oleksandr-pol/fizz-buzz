package main

import (
	"bufio"
	"fmt"
	"os"

	"../../internal/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			game.Play(reader)
		}
	}()

	res := game.Play(reader)
	showResult(res...)
}

func showResult(values ...string) {
	for _, val := range values {
		fmt.Println(val)
	}
}
