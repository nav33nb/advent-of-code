package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var input [][]rune

// var antinodes int

func Puz() {
	path := "day8/input/sample.txt"
	// path := "day8/input/input.txt"

	input = getInput(path)
	// for _, row := range input {
	// 	for _, c := range row {
	// 		fmt.Print(string(c), " ")
	// 	}
	// 	fmt.Println("")
	// }
	loc := map[rune][][2]int{}
	for i, row := range input {
		for j, c := range row {
			if c != '.' {
				loc[c] = append(loc[c], [2]int{i, j})
			}
		}
	}

	for k, v := range loc {
		fmt.Println(string(k), "-> ", v)
	}

	antinodes := 0
	for k, v := range loc {
		fmt.Printf("lines for %v: \n", string(k))
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				fmt.Printf("[%v,%v]", v[i], v[j])
				sx, sy := v[j][0]-v[i][0], v[j][1]-v[i][1]
				p1, p2 := v[i][0]-sx, v[i][1]-sy
				q1, q2 := v[j][0]+sx, v[j][1]+sy
				antinodes += setAntinode(p1, p2)
				antinodes += setAntinode(q1, q2)
				fmt.Printf(", antinodes=%v", antinodes)

				fmt.Println("")
			}
		}
		fmt.Println("")
	}

	for i, row := range input {
		for j, c := range row {
			if input[i][j] == '#' {
				// antinodes++
			}
			fmt.Print(string(c), " ")
		}
		fmt.Println("")
	}
	fmt.Println("ANSWER : ", antinodes)

}

func setAntinode(i, j int) int {
	// fmt.Println("setting for ", i, j)
	if i >= 0 && i < len(input) && j >= 0 && j < len(input[0]) {
		if input[i][j] == '.' || input[i][j] != '#' {
			input[i][j] = '#'
			fmt.Print(", set for ", i, j)
			return 1
		}
	}
	fmt.Print(", OOR for ", i, j)
	return 0
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

func Puz2() {
	path := "day8/input/sample.txt"
	// path := "day8/input/input.txt"

	input = getInput(path)
	// for _, row := range input {
	// 	for _, c := range row {
	// 		fmt.Print(string(c), " ")
	// 	}
	// 	fmt.Println("")
	// }
	loc := map[rune][][2]int{}
	for i, row := range input {
		for j, c := range row {
			if c != '.' {
				loc[c] = append(loc[c], [2]int{i, j})
			}
		}
	}

	for k, v := range loc {
		fmt.Println(string(k), "-> ", v)
	}

	antinodes := 0
	for k, v := range loc {
		fmt.Printf("lines for %v: \n", string(k))
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				fmt.Printf("[%v,%v]", v[i], v[j])
				sx, sy := v[j][0]-v[i][0], v[j][1]-v[i][1]
				antinodes += setAllAntinodes(v[i][0], v[i][1], sx, sy)
				fmt.Printf(", antinodes=%v", antinodes)

				fmt.Println("")
			}
		}
		fmt.Println("")
	}

	for i, row := range input {
		for j, c := range row {
			if input[i][j] == '#' {
				// antinodes++
			}
			fmt.Print(string(c), " ")
		}
		fmt.Println("")
	}
	fmt.Println("ANSWER : ", antinodes)

}

func setAllAntinodes(i, j, sx, sy int) int {
	var count int
	for p, q := i, j; p >= 0 && p < len(input) && q >= 0 && q < len(input[0]); p, q = p+sx, q+sy {
		if input[p][q] == '.' || input[p][q] != '#' {
			input[p][q] = '#'
			fmt.Print(", set for ", p, q)
			count++
		}
	}
	for p, q := i, j; p >= 0 && p < len(input) && q >= 0 && q < len(input[0]); p, q = p-sx, q-sy {
		if input[p][q] == '.' || input[p][q] != '#' {
			input[p][q] = '#'
			fmt.Print(", set for ", p, q)
			count++
		}
	}
	return count
}
