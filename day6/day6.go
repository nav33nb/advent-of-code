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

var rot = map[rune]rune{
	'^': '>',
	'v': '<',
	'<': '^',
	'>': 'v',
}

func Puz() {
	// input := getInput("day6/input/sample1.txt")
	input := getInput("day6/input/sample2.txt")

	// _,_,_,_,_ := processInput(input)
	currPos, direct, _, _ := processInput(input)

	// newPos := currPos
	for (currPos[0]+dirs[direct][0] < len(input)) && (currPos[1]+dirs[direct][1] < len(input[0])) {
		i, j := currPos[0]+dirs[direct][0], currPos[1]+dirs[direct][1]
		if 0 < i && i < len(input) && 0 < j && j < len(input[0]) && input[i][j] != '#' {
			currPos[0], currPos[1] = i, j
			continue
		}
		direct = rot[direct]
		fmt.Printf("Going %v at %v\n", string(direct), currPos)
	}

	fmt.Println(currPos)
}

// func newPos(currPos [2]int, direct [2]int, input []string, mapRow map[int][]int, mapCol map[int][]int) [2]int {
// 	if direct == '^' {
// 		for hashpos :=  range reverse(mapCol[direct[1]]) {
// 			if hashpos < currPos[i] {

// 			}
// 		}
// 	}
// }

func processInput(input []string) ([2]int, rune, map[int][]int, map[int][]int) {
	sPos := [2]int{}
	dir := [2]int{}
	mapRow := map[int][]int{}
	mapCol := map[int][]int{}

	isFound := false
	for i, row := range input {
		for j, c := range row {
			if c == '.' {
				continue
			}
			d, ok := dirs[c]
			if !isFound && ok {
				dir[0], dir[1] = d[0], d[1]
				sPos[0], sPos[1] = i, j
				isFound = true
				continue
			}
			mapRow[i] = append(mapRow[i], j)
			mapCol[j] = append(mapCol[j], i)
		}
	}

	fmt.Println(sPos, dir)
	fmt.Println(mapRow)
	fmt.Println(mapCol)

	return sPos, '^', mapRow, mapCol
}

func getInput(path string) []string {
	var input []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("cannot open input file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
