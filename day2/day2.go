package day2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const path = "day2/input/sample.txt"

func Puz1() {
	fmt.Println("Puzzle-1")
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()

	unsafe, safe := 0, 0
	scanner := bufio.NewScanner(file) // file input

	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		reportInt := make([]int, len(report))
		for i, val := range report {
			reportInt[i], _ = strconv.Atoi(val)
		}

		//consider report to be safe
		isSafe := true
		// make a copy and check for sorting
		temp := append([]int(nil), reportInt...)
		asc := slices.IsSorted(temp)
		slices.Reverse(temp)
		desc := slices.IsSorted(temp)
		if !(asc || desc) {
			isSafe = false
			continue
		}

		for i := 0; i < len(reportInt)-1; i++ {
			if reportInt[i] == reportInt[i+1] {
				isSafe = false
				break
			}

			diff := reportInt[i] - reportInt[i+1]
			if (diff < -3) || (3 < diff) {
				isSafe = false
				break
			}
		}

		if isSafe {
			safe++
		} else {
			unsafe++

		}
	}
	fmt.Printf("SAFE	: %v\n", safe)
	fmt.Printf("UNSAFE	: %v\n", unsafe)
}

func Puz2() {
	fmt.Println("Puzzle-2")
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()

	unsafe, safe := 0, 0
	scanner := bufio.NewScanner(file) // file input

	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		reportInt := make([]int, len(report))
		for i, val := range report {
			reportInt[i], _ = strconv.Atoi(val)
		}

		//consider report to be safe
		isSafe := true
		// make a copy and check for sorting
		temp := append([]int(nil), reportInt...)
		asc := slices.IsSorted(temp)
		slices.Reverse(temp)
		desc := slices.IsSorted(temp)
		if !(asc || desc) {
			isSafe = false
			continue
		}

		for i := 0; i < len(reportInt)-1; i++ {
			if reportInt[i] == reportInt[i+1] {
				isSafe = false
				break
			}

			diff := reportInt[i] - reportInt[i+1]
			if (diff < -3) || (3 < diff) {
				isSafe = false
				break
			}
		}

		if isSafe {
			safe++
		} else {
			unsafe++
			if correctionPossible(reportInt) {
				unsafe--
			}
		}
	}
	fmt.Printf("SAFE	: %v\n", safe)
	fmt.Printf("UNSAFE	: %v\n", unsafe)
}

func correctionPossible(reportInt []int) bool {
	panic("unimplemented")
}
