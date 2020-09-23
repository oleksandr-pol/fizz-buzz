package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../../pkg/games"
)

func Init() func() ([]string, bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var attemptsCount = 3
	var play func() ([]string, bool, error)

	fmt.Println("Game instructions:")
	fmt.Println("Please enter integer number for start and end of the game")
	fmt.Println("To exit the game enter 'exit'")
	fmt.Println()

	play = func() ([]string, bool, error) {
		if attemptsCount == 0 {
			return nil, false, errors.New("You entered wrong input too many times.")
		}
		attemptsCount--

		const maxInput = 100

		fmt.Print("Enter starting number for FizzBuzz game: ")
		start, exit, err := getInput(reader)
		if err != nil {
			fmt.Println(err)
			return play()
		}
		if exit {
			fmt.Println("Exiting")
			return nil, true, nil
		}

		fmt.Print("Enter last number for FizzBuzz game: ")
		end, exit, err := getInput(reader)
		if err != nil {
			fmt.Println(err)
			return play()
		}
		if exit {
			fmt.Println("Exiting")
			return nil, true, nil
		}

		if end <= start {
			fmt.Println("Starting number must be bigger then the last one")
			return play()
		}

		if end > maxInput {
			fmt.Println("too big last number")
			return play()
		}

		return getResult(start, end, games.FizzBuzz), false, nil
	}

	return play
}

func getInput(reader *bufio.Reader) (int, bool, error) {
	const exitInput = "exit"

	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Failed to read input")
	}

	if strings.TrimSpace(input) == exitInput {
		return 0, true, nil
	}

	param, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return 0, false, errors.New("Failed to parse input")
	}

	return param, false, nil
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
