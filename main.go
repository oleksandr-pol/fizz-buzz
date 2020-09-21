package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			play(reader)
		}
	}()

	res := play(reader)
	showResult(res...)
}

func FizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}

	if n%5 == 0 {
		return "Fizz"
	}

	if n%3 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func play(reader *bufio.Reader) []string {
	const maxInput = 100

	fmt.Print("Enter stating number for FizzBuzz game: ")
	start, err := getInput(reader)
	if err != nil {
		fmt.Println(err)
		return play(reader)
	}

	fmt.Print("Enter last number for FizzBuzz game: ")
	end, err := getInput(reader)
	if err != nil {
		fmt.Println(err)
		return play(reader)
	}

	if end <= start {
		fmt.Print("Starting number must be bigger then the last one \n")
		return play(reader)
	}

	if end > maxInput {
		panic("too big last number")
	}

	return getResult(start, end, FizzBuzz)
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

func showResult(values ...string) {
	for _, val := range values {
		fmt.Println(val)
	}
}
