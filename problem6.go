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

	treeMap := []string{}
	for scanner.Scan() {
		treeMap = append(treeMap, scanner.Text())
	}

	x := calculateTrees(treeMap, 1, 1)
	y := calculateTrees(treeMap, 3, 1)
	z := calculateTrees(treeMap, 5, 1)
	w := calculateTrees(treeMap, 7, 1)
	v := calculateTrees(treeMap, 1, 2)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(v)

	fmt.Println(w * x * y * z * v)
}

func calculateTrees(treeMap []string, columnStep int, lineStep int) int {

	lineLength := len(treeMap[0])
	trees := 0
	column := 0
	line := 0

	for line < len(treeMap) {
		calculatedColumn := column % lineLength
		if string(treeMap[line][calculatedColumn]) == "#" {
			trees++
		}
		column = column + columnStep
		line = line + lineStep
	}

	return trees

}
