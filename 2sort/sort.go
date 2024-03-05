package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(f([]int{3, 2, 1}))
}

func f(input []int) string {

	sort.Ints(input)

	inputStr := ""
	for index, v := range input {
		fmt.Println(index, v, len(input))
		if index != len(input)-1 {
			inputStr += strconv.Itoa(v) + " "
		} else {
			inputStr += strconv.Itoa(v)
		}
	}

	return inputStr
}
