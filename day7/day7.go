package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var target []int
var operands [][]int

func Puz() {
	path := "day7/input/sample.txt"
	target, operands = getInput(path)
	// fmt.Println(target, operands)
	// fmt.Println(target)
	// fmt.Println(operands)

	answer := 0
	for i := 0; i < len(target); i++ {
		// fmt.Println("====================")
		if val, ok := checkFor(target[i], operands[i][1:], operands[i][0]); ok {
			answer += val
		}
	}
	fmt.Println("Valid sequence sum1:", answer)
	answer = 0
	for i := 0; i < len(target); i++ {
		// fmt.Println("====================")
		if val, ok := checkFor2(target[i], operands[i][1:], operands[i][0]); ok {
			answer += val
		}
	}
	fmt.Println("Valid sequence sum2:", answer)
}

func checkFor(target int, operands []int, answer int) (int, bool) {
	// fmt.Println("checkfor target =", target, "on operands =", operands, "with answer =", answer)
	if answer > target {
		// fmt.Println("FAILED, overshot the answer WANT", target, "HAVE (", answer, ")")
		return answer, false
	}
	if len(operands) == 1 {
		if answer+operands[0] == target || answer*operands[0] == target {
			fmt.Println("SUCCESS, on any (+ or *) WANT (", target, ")\tHAVE (", answer+operands[0], "or", answer*operands[0], ")")
			return target, true
		} else {
			// fmt.Println("FAILED, on either adding or multiply WANT", target, "HAVE (", answer+operands[0], "or", answer*operands[0], ")")
			return 0, false
		}
	} else {
		// fmt.Println("Answer =", answer+operands[0], "=", answer, "+", operands[0])
		if t, ok := checkFor(target, operands[1:], answer+operands[0]); ok {
			return t, ok
		}
		// fmt.Println("Answer =", answer*operands[0], "=", answer, "*", operands[0])
		if t, ok := checkFor(target, operands[1:], answer*operands[0]); ok {
			return t, ok
		}
	}
	return answer, false
}

func checkFor2(target int, operands []int, answer int) (int, bool) {
	// fmt.Println("checkfor target =", target, "on operands =", operands, "with answer =", answer)
	val, _ := strconv.Atoi(fmt.Sprintf("%v%v", answer, operands[0]))
	if answer > target {
		// fmt.Println("FAILED, overshot WANT (", target, ") HAVE (", answer, ")")
		return answer, false
	}
	if len(operands) == 1 {
		if answer+operands[0] == target || answer*operands[0] == target || val == target {
			fmt.Println("SUCCESS, on any (+ or * or ||) WANT (", target, ") \tHAVE (", answer+operands[0], "or", answer*operands[0], "or", val, ")")
			return target, true
		} else {
			// fmt.Println("FAILED, on all WANT (", target, ") HAVE (", answer+operands[0], "or", answer*operands[0], "or", val, ")")
			return 0, false
		}
	} else {
		// fmt.Println("Answer =", answer+operands[0], "=", answer, "+", operands[0], "| remaining", operands[1:])
		if t, ok := checkFor2(target, operands[1:], answer+operands[0]); ok {
			return t, ok
		}
		// fmt.Println("Answer =", answer*operands[0], "=", answer, "*", operands[0], "| remaining", operands[1:])
		if t, ok := checkFor2(target, operands[1:], answer*operands[0]); ok {
			return t, ok
		}
		// fmt.Println("Answer =", val, "=", answer, "||", operands[0], "| remaining", operands[1:])
		if t, ok := checkFor2(target, operands[1:], val); ok {
			return t, ok
		}
	}
	return answer, false
}

func getInput(path string) ([]int, [][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("cannot open file, exiting")
	}

	var target []string
	var operands [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ":")
		target = append(target, l[0])
		operands = append(operands, strings.Split(strings.TrimSpace(l[1]), " "))
	}
	return convToInt(target, operands)
}

func convToInt(target []string, operands [][]string) ([]int, [][]int) {
	var iTarget []int
	var iOperands [][]int

	for _, v := range target {
		i, _ := strconv.Atoi(v)
		iTarget = append(iTarget, i)
	}
	for _, ops := range operands {
		var iops []int
		for _, op := range ops {
			i, _ := strconv.Atoi(op)
			iops = append(iops, i)
		}
		iOperands = append(iOperands, iops)
	}
	return iTarget, iOperands
}
