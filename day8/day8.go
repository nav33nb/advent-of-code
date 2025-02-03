package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Puz() {
	path := "day8/input/sample.txt"
	v := getInput(path)
	for _, row := range v {
		for _, c := range row {
			fmt.Print(string(c), " ")
		}
		fmt.Println("")
	}
}

func getInput(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("cannot open file, exiting")
	}

	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		input = append(input, line)
	}
	return input
}
