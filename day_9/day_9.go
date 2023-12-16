package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_line(s string) []int {
	nums := strings.Split(s, " ")
	var acc_nums []int
	for _, v := range nums {
		temp, _ := strconv.Atoi(v)
		acc_nums = append(acc_nums, temp)
	}
	return acc_nums
}

func slice_all_zeros(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func extrapolate(a []int) int {
	// base case all zeroes
	if slice_all_zeros(a) {
		return 0
	}

	// recursive case, construct next, extrapolate for next, add value to own last

	var construct []int
	for i := 0; i < len(a)-1; i++ {
		construct = append(construct, a[i+1]-a[i])
	}

	return a[len(a)-1] + extrapolate(construct)
}

func extrapolate_backwards(a []int) int {
	// base case all zeros
	if slice_all_zeros(a) {
		return 0
	}

	// recursive case

	var construct []int
	for i := 0; i < len(a)-1; i++ {
		construct = append(construct, a[i+1]-a[i])
	}

	return a[0] - extrapolate_backwards(construct)
}

// func extrapolate(a []int) {
// 	var extrapolation [][]int
// 	extrapolation = append(extrapolation, a)
// 	for i := 0; i < len(a); i++ {
// 		// Loop thru possible powers
// 		var construct []int
// 		for j := 0; j < len(a)-i-1; j++ {
// 			construct = append(construct, extrapolation[i][j+1]-extrapolation[i][j])
// 		}
// 		extrapolation = append(extrapolation, construct)
// 		if slice_all_zeros(construct) {
// 			break
// 		}
// 	}

// 	for
// }

func main() {
	data, _ := os.ReadFile("./text.txt")

	data_lines := strings.Split(string(data), "\n")

	var nums_parsed [][]int

	for _, v := range data_lines {
		nums_parsed = append(nums_parsed, parse_line(v))
	}

	ret := 0
	for _, v := range nums_parsed {
		ret += extrapolate(v)
		// fmt.Println(extrapolate(v))
	}
	fmt.Println(ret)

	ret2 := 0
	for _, v := range nums_parsed {
		ret2 += extrapolate_backwards(v)
	}
	fmt.Println(ret2)
}

// seems easy enough...
