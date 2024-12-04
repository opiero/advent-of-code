package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func readInput() [][]int {
	fp, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var reports [][]int
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		var levels []int
		for _, value := range values {
			intValue, _ := strconv.Atoi(value)
			levels = append(levels, intValue)
		}
		reports = append(reports, levels)

	}
	return reports
}

func isAscending(arr []int) bool {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	slices.Sort(sortedArr)
	return reflect.DeepEqual(arr, sortedArr)
}

func isDescending(arr []int) bool {
	reverseSortedArr := make([]int, len(arr))
	copy(reverseSortedArr, arr)

	sort.Slice(
		reverseSortedArr,
		func(a, b int) bool { return reverseSortedArr[b] < reverseSortedArr[a] },
	)
	return reflect.DeepEqual(arr, reverseSortedArr)
}

func isSafe(arr []int) bool {
	if !(isAscending(arr) || isDescending(arr)) {
		return false
	}

	for i := 0; i < len(arr)-1; i++ {
		difference := int(math.Abs(float64(arr[i] - arr[i+1])))

		if !(difference >= 1 && difference <= 3) {
			return false
		}
	}
	return true
}

func main() {
	reports := readInput()
	safeCounter := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCounter++
		}
	}

	fmt.Println(safeCounter)
}
