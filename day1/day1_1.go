package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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

	sort.Ints(list1)
	sort.Ints(list2)

	totalDifference := 0
	for i := 0; i < len(list1); i++ {
		totalDifference += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(totalDifference)
}
