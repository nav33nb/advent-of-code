package day2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const path = "day2/input/2.txt"

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
		r := make([]int, len(report))
		for i, val := range report {
			r[i], _ = strconv.Atoi(val)
		}

		if checkSafety(r) {
			// fmt.Println(r, "SAFE")
			safe++
		} else {
			// fmt.Println(r, "UNSAFE")
			unsafe++
		}
	}
	fmt.Printf("SAFE	: %v\n", safe)
	fmt.Printf("UNSAFE	: %v\n", unsafe)
}

func checkSafety(r []int) bool {
	// make a copy and check for sorting
	temp := append([]int(nil), r...)
	asc := slices.IsSorted(temp)
	slices.Reverse(temp)
	desc := slices.IsSorted(temp)
	if !(asc || desc) {
		// fmt.Printf("%v is UNSAFE due to SEQUENCE\n", r)
		return false
	}

	for i := 0; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			// fmt.Printf("%v is UNSAFE due to UNCHANGING\n", r)
			return false
		}

		diff := r[i] - r[i+1]
		if (diff < -3) || (3 < diff) {
			// fmt.Printf("%v is UNSAFE due to DIFFERENCE\n", r)
			return false
		}
	}
	// fmt.Printf("%v is SAFE\n", r)
	return true
}

func Puz2() {
	fmt.Println("\nPuzzle-2")
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()

	unsafe, safe := 0, 0
	scanner := bufio.NewScanner(file) // file input

	// trying a different approach for this
	for scanner.Scan() {
		report := strings.Split(scanner.Text(), " ")
		r := make([]int, len(report))
		for i, val := range report {
			r[i], _ = strconv.Atoi(val)
		}

		if checkSafety(r) || checkSafetyWithCorrection(r) {
			// fmt.Println(r, "SAFE\n")
			safe++
		} else {
			// fmt.Println(r, "UNSAFE\n")
			unsafe++
		}
	}
	fmt.Printf("SAFE	: %v\n", safe)
	fmt.Printf("UNSAFE	: %v\n", unsafe)
}

func checkSafetyWithCorrection(r []int) bool {
	for i := 0; i < len(r); i++ {
		var temp []int
		temp = append(temp, r[:i]...)
		temp = append(temp, r[i+1:]...)

		// fmt.Println(temp)
		if checkSafety(temp) {
			// fmt.Printf("%v is SAFE after removing %v\n", r, r[i])
			return true
		}
	}
	// fmt.Printf("%v is NEVER SAFE\n", r)
	return false
}
