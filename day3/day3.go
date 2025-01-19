package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Puz1() {
	fmt.Printf("Puzzle-1:  	")
	path := "day3/input/3.1.txt"
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`) // get all mul ops
	var sum int
	for scanner.Scan() {
		input := scanner.Text()
		ops := r.FindAllString(input, -1)
		for _, op := range ops {
			sum += getMul(op)
		}
	}
	fmt.Println(sum)
}

func getMul(op string) int {
	r, _ := regexp.Compile("[0-9]+") // get operands from single mul op
	nums := r.FindAllString(op, -1)
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1 * n2
}

func Puz2() {
	fmt.Printf("Puzzle-2:  	")
	path := "day3/input/3.2.txt"
	file, err := os.Open(path)
	if err != nil { // if error is not nil, print error and exit
		fmt.Printf("Cannot read input file: %v\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// get all enabled mul operation lines
	// i.e. any mul instruction between do() and don't()
	r, _ := regexp.Compile(`do\(\)|don't\(\)|mul\([0-9]+,[0-9]+\)`)
	var sum int
	for scanner.Scan() {
		input := scanner.Text()
		ops := r.FindAllString(input, -1)
		// fmt.Println(ops)
		processOp := true
		for _, op := range ops {
			switch {
			case op == "do()":
				processOp = true
			case op == "don't()":
				processOp = false
			case op[:3] == "mul":
				if processOp {
					sum += getMul(op)
				}
			default:
				panic(fmt.Errorf("unknown op: %v", op))
			}
		}
	}
	fmt.Println(sum)
}
