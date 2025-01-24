package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

var dirs = map[string]([]int){
	"t":  {-1, 0},
	"tr": {-1, 1},
	"r":  {0, 1},
	"dr": {1, 1},
	"d":  {1, 0},
	"dl": {1, -1},
	"l":  {0, -1},
	"tl": {-1, -1},
}

var target = "MAS"

func Puz1() {
	fmt.Println("Puzzle-1:")
	path := "day4/input/4.1.txt"
	input := getInput(path)
	findAllXmas(input)
}

func getInput(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot open sample file: %v, %v", path, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input [][]rune
	for scanner.Scan() {
		text := scanner.Text()
		row := []rune{}
		for _, r := range text {
			row = append(row, r)
		}
		input = append(input, row)
	}
	return input
}

func findAllXmas(input [][]rune) {
	var wg_send sync.WaitGroup
	var wg_rcv sync.WaitGroup
	res := make(chan int)

	// setup the rcvr for totals
	wg_rcv.Add(1)
	go func() {
		defer wg_rcv.Done()
		var sum int
		for val := range res {
			sum += val
		}
		fmt.Println("sum is ", sum)
	}()

	// setup senders for each XMAS
	for i, row := range input {
		for j, char := range row {
			if char == 'X' {
				wg_send.Add(1)
				go checkSpread(&wg_send, res, input, i, j)
			}
		}
	}

	wg_send.Wait()
	close(res)
	wg_rcv.Wait()
}

func checkSpread(wg *sync.WaitGroup, res chan int, input [][]rune, i, j int) {
	defer wg.Done()
	sum := 0
	for _, val := range dirs {
		isXmas := false
		dirY := i
		dirX := j
		for _, c := range target {
			dirY += val[0]
			dirX += val[1]
			if (dirY < 0 || len(input) <= dirY) ||
				(dirX < 0 || len(input[0]) <= dirX) ||
				input[dirY][dirX] != c {
				isXmas = false
				break
			} else if c == 'S' {
				isXmas = true
			}
		}
		if isXmas {
			// fmt.Printf("Found XMAS in %v at %v,%v\n", dir, i+val[0], j+val[1])
			sum++
		}
	}

	res <- sum
}

func Puz2() {
	fmt.Println("Puzzle-2:")
	path := "day4/input/4.1.txt"
	input := getInput(path)
	findAllMas(input)
}

func findAllMas(input [][]rune) {
	// setup senders for each XMAS
	sum := 0
	for i, row := range input {
		for j, char := range row {
			if char == 'A' {
				sum += checkCross(input, i, j)
				// fmt.Println("found A at", i, j, sum)
			}
		}
	}
	fmt.Println("sum is:", sum)
}

func checkCross(input [][]rune, i, j int) int {
	if i-1 < 0 ||
		len(input) <= i+1 ||
		j-1 < 0 ||
		len(input[0]) <= j+1 {
		return 0
	}

	var combs = []string{"MMSS", "MSSM", "SSMM", "SMMS"}
	curr := string([]rune{
		input[i-1][j-1],
		input[i-1][j+1],
		input[i+1][j+1],
		input[i+1][j-1],
	})

	// fmt.Println(curr)
	for _, comb := range combs {
		if comb == curr {
			return 1
		}
	}

	return 0
}
