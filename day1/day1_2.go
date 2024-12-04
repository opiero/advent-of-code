package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		elemList1, _ := strconv.Atoi(fields[0])
		elemList2, _ := strconv.Atoi(fields[1])

		list1 = append(list1, elemList1)
		list2 = append(list2, elemList2)
	}

	return list1, list2
}

func main() {
	list1, list2 := readInput()

	list2Counter := make(map[int]int)
	for _, num := range list2 {
		list2Counter[num]++
	}

	similarityScore := 0
	for _, num := range list1 {
		similarityScore += num * list2Counter[num]
	}

	fmt.Println(similarityScore)
}
