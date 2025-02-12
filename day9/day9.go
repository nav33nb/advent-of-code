package day9

import (
	"fmt"
	"strconv"
)

func Puz() {
	fmt.Println("Day 9, Puzzle 1")

	input := "2333133121414131402"

	// file, _ := os.Open("day9/input/input.txt")
	// scanner := bufio.NewScanner(file)
	// scanner.Scan()
	// input := scanner.Text()

	isFile := true
	var id int

	disk := []int{}
	for _, v := range input {
		val, _ := strconv.Atoi(string(v))
		if isFile {
			for i := 0; i < val; i++ {
				disk = append(disk, id)
			}
			id++
		} else {
			for i := 0; i < val; i++ {
				disk = append(disk, -1)
			}
		}
		isFile = !isFile
	}
	// fmt.Println(disk)

	// i := 0
	for i, j := 0, len(disk)-1; i < j; {
		if disk[i] != -1 {
			i++
			continue
		}
		for disk[j] == -1 {
			j--
			continue
		}
		disk[i], disk[j] = disk[j], -1
		// fmt.Println(disk, j, i)
		// fmt.Println(j, i)
	}

	checksum := 0
	for i, v := range disk {
		if v == -1 {
			continue
		}
		checksum += i * v
	}
	fmt.Println(checksum)
}

func Puz2() {
	fmt.Println("Day 9, Puzzle 2")

	input := "2333133121414131402"

	// file, _ := os.Open("day9/input/input.txt")
	// scanner := bufio.NewScanner(file)
	// scanner.Scan()
	// input := scanner.Text()

	isFile := true
	space := map[int]int{}
	file := map[int][2]int{}

	i := 0
	id := 0
	for _, c := range input {
		v, _ := strconv.Atoi(string(c))
		if isFile {
			// fmt.Println("setting file", i, id, v)
			file[i] = [2]int{id, v}
			id++
		} else {
			// fmt.Println("setting space", i, v)
			space[i] = v
		}
		i += v
		isFile = !isFile
	}
	fmt.Println("")

	for z := 0; z <= i; z++ {
		if q, ok := file[z]; ok {
			fmt.Println("file", z, "->", q)
		}
	}

	for z := 0; z <= i; z++ {
		if q, ok := space[z]; ok {
			fmt.Println("space", z, "->", q)
		}
	}

	fmt.Println("")

	//logic here

	for k, v := range file {
		fmt.Println(k, " -> ", v)
	}
	for k, v := range space {
		fmt.Println(k, " -> ", v)
	}

	for q := 0; q < i; q++ {
		if f, ok := file[q]; ok {
			for j := 0; j < f[1]; j++ {
				fmt.Print(f[0])
				q++
			}
		}
		if s, ok := space[q]; ok {
			for j := 0; j < s; j++ {
				fmt.Print(".")
				q++
			}
		}
	}

	checksum := 0
	for k, v := range file {
		checksum += k * v[0]
	}
	fmt.Println("ANSWER:", checksum)
}
