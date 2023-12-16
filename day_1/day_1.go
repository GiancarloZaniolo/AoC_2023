package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func find_val(s string) int {
	var first, last int
	for i, v := range s {
		if unicode.IsDigit(v) {
			first = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			last = i
			break
		}
	}
	final_s := string(s[first]) + string(s[last])

	ret, _ := strconv.Atoi(final_s)

	// fmt.Printf("For string %s got %d\n", s, ret)

	return ret
}

var nums = [10]string{"zero", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine"}

func find_val2(s string) int {
	var first, last int
	for i, v := range s {
		if unicode.IsDigit(v) {
			first = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			last = i
			break
		}
	}
	// fmt.Printf("first round %d  %d ", int(s[first])-48, int(s[last])-48)
	//use first and last to narrow down length of search
	var first_idx, first_val, last_idx, last_val int
	first_idx = 100000
	last_idx = -1
	for i, v := range nums {
		curr_idx := strings.Index(s, v)
		if curr_idx != -1 && curr_idx < first_idx {
			first_idx = curr_idx
			first_val = i
		}

		curr_end_idx := strings.LastIndex(s, v)
		if curr_end_idx != -1 && curr_end_idx > last_idx {
			last_idx = curr_end_idx
			last_val = i
		}
	}
	// fmt.Printf("last round %d  %d\n", first_val, last_val)

	var final_first, final_last int
	if first < first_idx {
		final_first = int(s[first]) - 48
	} else {
		final_first = first_val
	}

	if last > last_idx {
		final_last = int(s[last]) - 48
	} else {
		final_last = last_val
	}

	// fmt.Printf("final %d %d\n", final_first, final_last)
	return final_first*10 + final_last
}

func main() {
	// dat, _ := os.ReadFile("./text2.txt")
	// Could check but I'm not gunna
	// check(err)

	// fmt.Print(string(dat))

	f, _ := os.Open("./text.txt")

	defer f.Close()

	s := bufio.NewScanner(f)

	var acc []string
	for s.Scan() {
		data := s.Text()
		acc = append(acc, data)
	}

	ret := 0
	ret2 := 0
	for _, v := range acc {
		// For part 1
		ret += find_val(v)
		ret2 += find_val2(v)
	}
	fmt.Println(ret)
	fmt.Println(ret2)
	// 54589 too high
}
