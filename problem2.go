package main

import (
	"bufio"
	"fmt"
	"os"
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

	for _, lnum := range numbers {
		for _, unum := range numbers {
			for _, tnum := range numbers {
				if lnum+unum+tnum == 2020 {
					fmt.Println("Found result: ")
					fmt.Println(lnum)
					fmt.Println(unum)
					fmt.Println(tnum)
					fmt.Println(lnum * unum * tnum)
					break
				}
			}
		}
	}
}
