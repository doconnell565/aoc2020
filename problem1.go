package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("problem1.dat")
	check(err)
	numbers := []int64{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 64)
		check(err)
		numbers = append(numbers, number)
	}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	var halfway int64 = 2020 / 2
	var upper []int64
	var lower []int64

	for i, num := range numbers {
		if num > halfway {
			upper = numbers[i+1 : len(numbers)]
			lower = numbers[0:i]
			break
		}
	}

	for _, lnum := range lower {
		for _, unum := range upper {
			if lnum+unum == 2020 {
				fmt.Println("Found result: ")
				fmt.Println(lnum * unum)
				break
			}
		}
	}
}
