package game

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"../../pkg/games"
)

func Play(reader *bufio.Reader) []string {
	const maxInput = 100

	fmt.Print("Enter stating number for FizzBuzz game: ")
	start, err := getInput(reader)
	if err != nil {
		fmt.Println(err)
		return Play(reader)
	}

	fmt.Print("Enter last number for FizzBuzz game: ")
	end, err := getInput(reader)
	if err != nil {
		fmt.Println(err)
		return Play(reader)
	}

	if end <= start {
		fmt.Print("Starting number must be bigger then the last one \n")
		return Play(reader)
	}

	if end > maxInput {
		panic("too big last number")
	}

	return getResult(start, end, games.FizzBuzz)
}

func getInput(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	param, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, err
	}

	return param, nil
}

func getResult(start int, end int, game func(int) string) []string {
	res := make([]string, end-start+1)
	j := start
	for i := 0; i <= end-start; i++ {
		res[i] = game(j)
		j++
	}

	return res
}
