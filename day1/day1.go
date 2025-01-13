package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const path = "day1/input/1.txt"

func Puz1() {
	fmt.Printf("Puzzle-1: ")
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // file input
	var list1, list2 []int
	for scanner.Scan() {
		valuePair := strings.Split(scanner.Text(), "   ") // split by 3 spaces
		v1, _ := strconv.Atoi(valuePair[0])
		v2, _ := strconv.Atoi(valuePair[1])

		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}
	slices.Sort(list1)
	slices.Sort(list2)

	answer := 0
	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		answer += diff
	}
	fmt.Println("Distance is", answer)
}

func Puz2() {
	fmt.Printf("Puzzle-2: ")
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // file input
	var list1, list2 []string
	for scanner.Scan() {
		valuePair := strings.Split(scanner.Text(), "   ") // split by 3 spaces

		list1 = append(list1, valuePair[0])
		list2 = append(list2, valuePair[1])
	}

	var numCount = make(map[string]int)
	for i := 0; i < len(list2); i++ {
		numCount[list2[i]]++
	}
	sum := 0
	for i := 0; i < len(list1); i++ {
		val, _ := strconv.Atoi(list1[i])
		sum += val * numCount[list1[i]]
	}

	fmt.Println("Similarity is", sum)
}
