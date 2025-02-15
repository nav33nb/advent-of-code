package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = [][]int{}

func Puz() {
	path := "day10/input/sample.txt"
	input = getInput(path)
	score := 0
	for i, row := range input {
		for j, val := range row {
			if val == 0 {
				visited = make([][]bool, len(input))
				for i := range visited {
					visited[i] = make([]bool, len(input[0]))
				}
				score += trailHead(i, j, -1)
				// fmt.Println("RUNNING score is ", score)
			}
		}
	}
	fmt.Println("Trailhead score is ", score)
}

var visited = [][]bool{}

func trailHead(i, j, prev int) int {
	if i < 0 || len(input) <= i || j < 0 || len(input[0]) <= j || input[i][j] != prev+1 {
		fmt.Print("F")
		return 0
	}
	if input[i][j] == 9 && prev == 8 {
		if visited[i][j] {
			return 0
		}
		fmt.Println(" SUCCESS 9 at", i, j)
		// visited[i][j] = true    /// ONLY CHANGE for PART 2
		return 1
	}
	fmt.Println("")
	var left, right, up, down int
	fmt.Print(" R", input[i][j])
	right = trailHead(i, j+1, input[i][j])
	fmt.Print(" D", input[i][j])
	down = trailHead(i+1, j, input[i][j])
	fmt.Print(" L", input[i][j])
	left = trailHead(i, j-1, input[i][j])
	fmt.Print(" U", input[i][j])
	up = trailHead(i-1, j, input[i][j])
	return left + right + up + down
}

func getInput(path string) [][]int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	nums := [][]int{}
	for scanner.Scan() {
		char := strings.Split(scanner.Text(), "")
		row := []int{}
		for _, c := range char {
			i, _ := strconv.Atoi(c)
			row = append(row, i)
		}
		nums = append(nums, row)
	}
	return nums
}

func Puz2() {

}
