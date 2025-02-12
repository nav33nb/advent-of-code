package main

import (
	"advent2024/day1"
	"advent2024/day2"
	"advent2024/day3"
	"advent2024/day4"
	"advent2024/day5"
	"advent2024/day6"
	"advent2024/day7"
	"advent2024/day8"
	"advent2024/day9"
	"fmt"
	"time"
)

func main() {
	fmt.Println(">>>>>>>>>> Day1")
	day1.Puz1()
	day1.Puz2()

	fmt.Println()
	fmt.Println(">>>>>>>>>> Day2")
	day2.Puz1()
	day2.Puz2()

	fmt.Println()
	fmt.Println(">>>>>>>>>> Day3")
	day3.Puz1()
	day3.Puz2()

	fmt.Println()
	fmt.Println(">>>>>>>>>> Day4")
	day4.Puz1()
	day4.Puz2()

	fmt.Println()
	fmt.Println(">>>>>>>>>> Day5")
	day5.Puz()

	fmt.Println()
	fmt.Println(">>>>>>>>>> Day6")
	day6.Puz()
	day6.Puz2()

	t := time.Now()
	fmt.Println()
	fmt.Println(">>>>>>>>>> Day7")
	day7.Puz()
	fmt.Println(time.Since(t))

	t = time.Now()
	fmt.Println()
	fmt.Println(">>>>>>>>>> Day8")
	day8.Puz()
	day8.Puz2()
	fmt.Println(time.Since(t))

	t = time.Now()
	fmt.Println()
	fmt.Println(">>>>>>>>>> Day9")
	day9.Puz()
	day9.Puz2()
	fmt.Println(time.Since(t))

}
