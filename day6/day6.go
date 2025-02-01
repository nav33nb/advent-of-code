package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var dirs = map[rune][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

var rotate = map[rune]rune{
	'^': '>',
	'v': '<',
	'<': '^',
	'>': 'v',
}

func Puz() {
	// input := getInput("day6/input/sample.txt")
	input := getInput("day6/input/input.txt")

	currPos, d := processInput(input)
	positions := 0
	for {
		i, j := currPos[0]+dirs[d][0], currPos[1]+dirs[d][1]
		if 0 > i || i >= len(input) || 0 > j || j >= len(input[0]) {
			fmt.Println("Got out at", i, j)
			break
		}
		if input[i][j] != '#' {
			currPos[0], currPos[1] = i, j
			if input[i][j] != 'X' {
				input[i][j] = 'X'
				positions++
			}
			continue
		}
		d = rotate[d]
		fmt.Printf("Going %v at %v\n", string(d), currPos)
	}
	// for _, row := range input {
	// 	fmt.Println(string(row))
	// }
	// time.Sleep(500 * time.Millisecond)
	fmt.Println(currPos, positions)

}

func Puz2() {
	input := getInput("day6/input/sample2.txt")
	// input := getInput("day6/input/input.txt")

	currPos, d, rowMap, colMap := processInput2(input)
	visitedRowMap := map[int][]bool{}
	visitedColMap := map[int][]bool{}

	positions := 0
	i, j := currPos[0], currPos[1]
	for {
		if 0 > i || i >= len(input) || 0 > j || j >= len(input[0]) {
			fmt.Println("Got out at", i, j)
			break
		}
		// var ki, kj int
		var loop int
		if input[i][j] != '#' {
			switch d {
			case '^':
				for q := j; q < len(input[i]); q++ {
					if visitedRowMap[i][q] {
						loop++
					}
				}
			case '>':
				// check colmap j after i
				for p := i; p < len(input); p++ {
					if visitedRowMap[p][j] {
						loop++
					}
				}
			case 'v':
				//check rowMap i before j
				for q := j; q > 0; q-- {
					if visitedRowMap[i][q] {
						loop++
					}
				}
			case '<':
				// check colmap j before i
				for p := i; p > 0; p-- {
					if visitedRowMap[p][j] {
						loop++
					}
				}
			}
			continue
		}
		d = rotate[d]
		fmt.Printf("Going %v at %v\n", string(d), currPos)
		i, j = currPos[0]+dirs[d][0], currPos[1]+dirs[d][1]
	}
	// for _, row := range input {
	// 	fmt.Println(string(row))
	// }
	// time.Sleep(500 * time.Millisecond)
	fmt.Println(currPos, positions)

}

func processInput(input [][]rune) ([2]int, rune) {
	for i, row := range input {
		for j, c := range row {
			if _, ok := dirs[c]; ok {
				return [2]int{i, j}, c
			}
		}
	}
	// invalid return
	return [2]int{}, ' '
}

func processInput2(input [][]rune) ([2]int, rune, map[int][]int, map[int][]int) {
	var currPos [2]int
	var direction rune
	rowMap, colMap := map[int][]int{}, map[int][]int{}
	for i, row := range input {
		for j, c := range row {
			if c == '.' {
				continue
			} else if c == '#' {
				rowMap[i] = append(rowMap[i], j)
				rowMap[j] = append(rowMap[j], i)
			} else { // will only execute once, else is fine
				currPos = [2]int{i, j}
				direction = c
			}
		}
	}
	// invalid return
	return currPos, direction, rowMap, colMap
}

func getInput(path string) [][]rune {
	var input [][]rune
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("cannot open input file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}

	return input
}
