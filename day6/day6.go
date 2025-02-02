package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var input [][]rune

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
	input = getInput("day6/input/sample2.txt")
	// input := getInput("day6/input/input.txt")

	currPos, d := processInput(input)
	positions := 0
	for {
		i, j := currPos[0]+dirs[d][0], currPos[1]+dirs[d][1]
		if 0 > i || i >= len(input) || 0 > j || j >= len(input[0]) {
			// fmt.Println("Got out at", i, j)
			// fmt.Println(currPos, positions)
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
		// fmt.Printf("Going %v at %v\n", string(d), currPos)
	}
	fmt.Println("[INFO] ANSWER", positions)
	for _, row := range input {
		fmt.Println(string(row))
	}
}

func checkLoop(currPos [2]int, d rune) (bool, [][2]int) {
	// fmt.Println("Starting from", currPos, "in", d)
	var positions [][2]int
	visited := map[string][]rune{}
	i, j := currPos[0], currPos[1]
	// fmt.Printf("Going %v at %v,%v: ", string(d), i, j)
	for {
		i, j = i+dirs[d][0], j+dirs[d][1]
		key := fmt.Sprintf("%v,%v", i, j)
		if 0 > i || i >= len(input) || 0 > j || j >= len(input[0]) {
			// fmt.Println("Exit found:", i, j)
			// fmt.Println(currPos, positions)
			return false, positions
		}
		if input[i][j] != '#' {
			positions = append(positions, [2]int{i, j})
			continue
		}
		if !slices.Contains(visited[key], d) {
			visited[key] = append(visited[key], d)
		} else {
			return true, positions
		}
		i, j = i-dirs[d][0], j-dirs[d][1]
		d = rotate[d]
		// fmt.Printf("  Going %v at %v,%v\n", string(d), i, j)
	}

}

func Puz2() {
	input = getInput("day6/input/sample2.txt")
	// input = getInput("day6/input/input.txt")
	currPos, d := processInput2()
	var loop int
	_, positions := checkLoop(currPos, d)
	// fmt.Println(isOut, loop)
	// fmt.Println("====================================")
	loopPos := [][2]int{}

	loopExists := map[string]bool{}
	// fmt.Println(len(positions))
	for _, pos := range positions[1:] {
		// fmt.Print("new O at ", pos[0], ",", pos[1], " : ")
		input[pos[0]][pos[1]] = '#'
		key := fmt.Sprintf("%v,%v", pos[0], pos[1])

		if loopExists[key] {
			// fmt.Println("setting with pos[0]=", pos[0], "pos[1]", pos[1])
			input[pos[0]][pos[1]] = '.'
			continue
		}
		if ok, _ := checkLoop(currPos, d); ok {
			// fmt.Println("Loop found: ", pos[0], pos[1])
			loop++
			loopPos = append(loopPos, [2]int{pos[0], pos[1]})
			loopExists[key] = true
		}

		// fmt.Println("setting with pos[0]=", pos[0], "pos[1]", pos[1])
		input[pos[0]][pos[1]] = '.'
		// fmt.Printf("%v", string(input[pos[0]][pos[1]]))
		// fmt.Println(loop)
	}

	for _, v := range loopPos {
		input[v[0]][v[1]] = 'o'
	}

	fmt.Println()
	fmt.Println("[INFO] ANSWER ", loop)
	for _, row := range input {
		fmt.Println(string(row))
	}
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

func processInput2() ([2]int, rune) {
	var currPos [2]int
	var direction rune
	// rowMap, colMap := map[int][]int{}, map[int][]int{}
	for i, row := range input {
		for j, c := range row {
			if c == '.' {
				continue
			} else if c == '#' {
				// rowMap[i] = append(rowMap[i], j)
				// colMap[j] = append(colMap[j], i)
			} else { // will only execute once
				currPos = [2]int{i, j}
				direction = c
				input[i][j] = '.'
			}
		}
	}
	// invalid return
	return currPos, direction //, rowMap, colMap
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
