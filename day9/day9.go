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

//////////////////////////////////////////////

var space = map[int]int{}
var file = map[int][2]int{}

func Puz2() {
	fmt.Println("Day 9, Puzzle 2")

	input := "2333133121414131402"

	// foo, _ := os.Open("day9/input/input.txt")
	// scanner := bufio.NewScanner(foo)
	// scanner.Scan()
	// input := scanner.Text()

	isFile := true

	diskHeader := 0
	id := 0
	for _, c := range input {
		v, _ := strconv.Atoi(string(c))
		if isFile {
			// fmt.Println("setting file", i, id, v)
			file[diskHeader] = [2]int{id, v}
			id++
		} else {
			// fmt.Println("setting space", i, v)
			space[diskHeader] = v
		}
		diskHeader += v
		isFile = !isFile
	}
	fmt.Println("")

	fmt.Println("Disk size is ", diskHeader)
	fancyPrinter(diskHeader)
	fmt.Println("  Original")

	for bw := diskHeader - 1; 0 <= bw; bw-- {
		if f, ok := file[bw]; ok {
			for fw := 0; fw < bw; fw++ {
				if free, ok := space[fw]; ok && free >= f[1] {
					space[fw+f[1]] = space[fw] - f[1] // set reduced space on an offset
					file[fw] = file[bw]
					delete(space, fw) // delete space on fw head
					delete(file, bw)  // delete file on bw head
					space[bw] = f[1]  // set new empty, on moved file
					fancyPrinter(diskHeader)
					fmt.Println("  Moved", f[0])
					break
				}
			}
		}
	}

	fmt.Println("")
	fmt.Println("ANSWER:", checksumCalc(diskHeader))
}

func fancyPrinter(diskHeader int) {
	for fw := 0; fw <= diskHeader; fw++ {
		if f, ok := file[fw]; ok {
			for range f[1] {
				fmt.Print(f[0])
			}
		} else if s, ok := space[fw]; ok && s != 0 {
			for range s {
				fmt.Print(".")
			}
		}
	}
}

func checksumCalc(diskHeader int) int {
	checksum := 0
	for q := 0; q < diskHeader; q++ {
		if f, ok := file[q]; ok {
			for r := q; r < q+f[1]; r++ {
				checksum += r * f[0]
			}
		}
	}
	return checksum
}
