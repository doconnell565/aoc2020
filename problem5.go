package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("problem5.dat")
	check(err)

	scanner := bufio.NewScanner(file)

	trees := 0
	column := 0
	line := 0
	treeMap := []string{}
	for scanner.Scan() {
		treeMap = append(treeMap, scanner.Text())
	}
	lineLength := len(treeMap[0])
	for line < len(treeMap) {
		calculatedColumn := column % lineLength
		if string(treeMap[line][calculatedColumn]) == "#" {
			trees++
		}
		column = column + 3
		line++
	}

	fmt.Println(trees)

}
