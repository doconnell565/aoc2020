package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("problem3.dat")
	check(err)

	scanner := bufio.NewScanner(file)

	validCount := 0

	fmt.Println(scanLine("1-3 a: abcde"))
	fmt.Println(scanLine("1-3 b: cdefg"))
	fmt.Println(scanLine("2-9 c: ccccccccc"))

	for scanner.Scan() {
		line := scanner.Text()
		if scanLine(line) {
			validCount++
		}
	}

	fmt.Println(validCount)

}

func scanLine(line string) bool {
	r := regexp.MustCompile(`(?P<rangeStart>\d*)-(?P<rangeEnd>\d*)\s*(?P<character>\S):\s* (?P<password>.*)`)
	res := r.FindStringSubmatch(line)

	rangeStart, _ := strconv.Atoi(res[r.SubexpIndex("rangeStart")])
	rangeEnd, _ := strconv.Atoi(res[r.SubexpIndex("rangeEnd")])
	character := res[r.SubexpIndex("character")]
	password := res[r.SubexpIndex("password")]

	if len(password) > rangeStart && len(password) > (rangeEnd-1) {
		firstChar := password[rangeStart-1]
		secondChar := password[rangeEnd-1]
		if firstChar == secondChar {
			return false
		}

		if firstChar == character[0] || secondChar == character[0] {
			// fmt.Println(character)
			// fmt.Println(rangeStart)
			// fmt.Println(rangeEnd)
			// fmt.Println(password)
			// fmt.Println(len(password))
			return true
		}

	}

	return false
}
