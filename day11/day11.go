package day11

import (
	"fmt"
	"strconv"
)

func Puz() {
	input := []int{7568, 155731, 0, 972, 1, 6919238, 80646, 22}
	// input := []int{125, 17}
	ops := 75

	pmap := map[int]int{}
	for _, val := range input {
		pmap[val] += 1
	}
	fmt.Println(pmap)

	for range ops {
		pmap = smartBlink(pmap)
		sum := 0
		for _, v := range pmap {
			sum += v
		}
		fmt.Print(sum, " ")
	}

}

func smartBlink(pmap map[int]int) map[int]int {
	outmap := map[int]int{}
	for k, v := range pmap {
		// fmt.Println("In Loop", outmap)
		str := strconv.Itoa(k)
		switch {
		case k == 0:
			// fmt.Println("zero")
			outmap[1] += v
		case len(str)%2 == 0:
			// fmt.Println("even")
			first, err := strconv.Atoi(str[:len(str)/2])
			if err != nil {
				fmt.Errorf("Error in conversion")
			}
			second, err := strconv.Atoi(str[len(str)/2:])
			if err != nil {
				fmt.Errorf("Error in conversion")
			}
			outmap[first] += v
			outmap[second] += v
		default:
			// fmt.Println("default")
			outmap[k*2024] += v
		}
	}
	return outmap
}

func bruteBlink(input []int) []int {
	var output []int
	for _, val := range input {
		str := strconv.Itoa(val)
		switch {
		case val == 0:
			output = append(output, 1)
		case len(str)%2 == 0:
			first, _ := strconv.Atoi(str[:len(str)/2])
			second, _ := strconv.Atoi(str[len(str)/2:])
			output = append(output, first)
			output = append(output, second)
		default:
			output = append(output, val*2024)
		}
	}
	return output
}

func Puz2() {

}
