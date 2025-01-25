package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Puz1() {
	// rules, pages := getInput("day5/input/sample1.txt")
	rules, pages := getInput("day5/input/5.1.txt")
	// fmt.Println(rules)
	// fmt.Println(pages)

	// prevMap := map[string][]string{}
	afterMap := map[string][]string{}

	for _, rule := range rules {
		afterMap[rule[0]] = append(afterMap[rule[0]], rule[1])
		// prevMap[rule[1]] = append(afterMap[rule[1]], rule[0])
	}
	fmt.Println(afterMap)
	// fmt.Println(prevMap)

	sum := 0
	for i := 0; i < len(pages); i++ {
		allowedValues := append([]string{}, afterMap[pages[i][0]]...)
		// fmt.Printf("allowed vals after %v is %v\n", pages[i][0], allowedValues)
		pass := true
		// fmt.Println()
		for j := 1; j < len(pages[i]); j++ {
			fmt.Printf("allowed vals after %v is %v\n", pages[i][j], allowedValues)
			if !slices.Contains(allowedValues, pages[i][j]) {
				// fmt.Printf("Failed for %v\n", pages[i])
				pass = false
				break
			}
			// fmt.Printf("%v can come after %v\n", pages[i][j], pages[i][j-1])
			allowedValues = getIntersection(allowedValues, afterMap[pages[i][j]])
		}
		if pass {
			fmt.Printf("PASS %v\n", pages[i])
			num, _ := strconv.Atoi(pages[i][len(pages[i])/2])
			sum += num
		} else {
			fmt.Printf("FAIL %v\n", pages[i])
		}
	}
	fmt.Println(sum)
}

func getIntersection(s1 []string, s2 []string) []string {
	intersection := []string{}
	for _, s := range s1 {
		if slices.Contains(s2, s) {
			intersection = append(intersection, s)
		}
	}
	// fmt.Printf("%v intersect %v is %v\n", s2, s1, intersection)
	return intersection
}

func getInput(path string) ([][]string, [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot open sample file: %v, %v", path, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	isRulesOver := false
	rules, pages := [][]string{}, [][]string{}
	for scanner.Scan() {
		text := scanner.Text()
		if !isRulesOver && text == "" {
			isRulesOver = true
			continue
		}
		if isRulesOver {
			vals := strings.Split(text, ",")
			pages = append(pages, vals)
		} else {
			vals := strings.Split(text, "|")
			rules = append(rules, vals)
		}
	}
	return rules, pages
}
