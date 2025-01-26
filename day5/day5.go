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

func Puz() {
	rules, pageList := getInput("day5/input/sample1.txt")
	// rules, pages := getInput("day5/input/5.1.txt")
	// fmt.Println(rules)
	// fmt.Println(pages)

	afterMap := map[string][]string{}
	for _, rule := range rules {
		afterMap[rule[0]] = append(afterMap[rule[0]], rule[1])
	}
	// fmt.Println(afterMap)

	sum := 0
	correctedSum := 0

	for _, pages := range pageList {
		pass := checkPageOrder(pages, afterMap)
		if pass {
			fmt.Printf("PASS %v\n", pages)
			num, _ := strconv.Atoi(pages[len(pages)/2])
			sum += num
		} else {
			fmt.Printf("FAIL %v\n", pages)
			mid, err := getCorrectedMid(pages, afterMap)
			if err != nil {
				log.Fatal(err)
			}
			correctedSum += mid
		}
	}
	fmt.Println(sum)
	fmt.Println(correctedSum)

}

func checkPageOrder(pages []string, afterMap map[string][]string) bool {
	allowedValues := []string{}
	// fmt.Printf("allowed vals for %v after %v is %v\n", pages[i][0], "''", allowedValues)
	allowedValues = append(allowedValues, afterMap[pages[0]]...)
	pass := true
	for _, page := range pages[1:] {
		// fmt.Printf("allowed vals for %v is %v\n", page, allowedValues)
		if !slices.Contains(allowedValues, page) {
			fmt.Printf("Failed for %v\n", page)
			pass = false
			break
		}
		allowedValues = getIntersection(allowedValues, afterMap[page])
	}
	return pass
}

func getCorrectedMid(pages []string, m map[string][]string) (int, error) {
	possibilities := map[string][]string{}
	fmt.Printf("  TOTAL: %v elements\n", len(pages))
	for _, page := range pages {
		intersection := getIntersection(m[page], pages)
		possibilities[page] = intersection
		fmt.Printf("  correct index for %v is i=%v\n", page, len(intersection))
		if len(intersection) == len(pages)/2 {
			fmt.Println("    middle ele is", page)
			return strconv.Atoi(page)
		}
	}
	return 0, fmt.Errorf("no middle ele found")
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
