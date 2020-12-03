package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile(`(?P<rangeStart>\d*)-(?P<rangeEnd>\d*)\s*(?P<character>\S):\s* (?P<password>.*)`)
		res := r.FindStringSubmatch(line)

		rangeStart, _ := strconv.Atoi(res[r.SubexpIndex("rangeStart")])
		rangeEnd, _ := strconv.Atoi(res[r.SubexpIndex("rangeEnd")])
		character := res[r.SubexpIndex("character")]
		password := res[r.SubexpIndex("password")]

		repCount := strings.Count(password, character)
		if repCount >= rangeStart && repCount <= rangeEnd {
			validCount++
		}
	}

	fmt.Println(validCount)

}
