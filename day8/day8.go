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
	loc := map[rune][][2]int{}
	for i, row := range v {
		for j, c := range row {
			if c != '.' {
				loc[c] = append(loc[c], [2]int{i, j})
			}
		}
	}

	// for k, v := range loc {
	// 	fmt.Println(string(k), "-> ", v)
	// }

	for k, v := range loc {
		for i, _ := range v {
			for j=i+1; j<len(v) {

			}
		}
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
