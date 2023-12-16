package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcp1(s [][]string) bool {
	var red_count, blue_count, green_count int
	for _, v := range s {
		for i := 1; i < len(v); i += 2 {
			temp, _ := strconv.Atoi(v[i])
			// fmt.Printf("%s  %d\n", v[i+1], temp)
			switch v[i+1] {
			case "blue":
				blue_count = max(blue_count, temp)
			case "red":
				red_count = max(red_count, temp)
			case "green":
				green_count = max(green_count, temp)
			default:
				fmt.Printf("string \"%s\" did not match with any options!", v[2])
			}
		}
	}
	// fmt.Printf("red %d  green %d  blue %d\n", red_count, green_count, blue_count)
	if red_count > 12 || green_count > 13 || blue_count > 14 {
		return false
	}
	return true
}

func calcp2(s [][]string) int {
	var red_count, blue_count, green_count int
	for _, v := range s {
		for i := 1; i < len(v); i += 2 {
			temp, _ := strconv.Atoi(v[i])
			// fmt.Printf("%s  %d\n", v[i+1], temp)
			switch v[i+1] {
			case "blue":
				blue_count = max(blue_count, temp)
			case "red":
				red_count = max(red_count, temp)
			case "green":
				green_count = max(green_count, temp)
			default:
				fmt.Printf("string \"%s\" did not match with any options!", v[2])
			}
		}
	}
	// fmt.Printf("red %d  green %d  blue %d\n", red_count, green_count, blue_count)
	return red_count * blue_count * green_count
}

func main() {
	f, _ := os.Open("./text.txt")

	defer f.Close()

	s := bufio.NewScanner(f)

	var acc []string
	for s.Scan() {
		data := s.Text()
		acc = append(acc, data)
	}

	// Additional linear processing step
	var acc2 [][][]string
	for _, v := range acc {
		temp := strings.ReplaceAll(v, ",", "")
		games := strings.Split(temp, ";")
		var counts [][]string
		for _, v := range games {
			counts = append(counts, strings.Split(v, " "))
		}
		counts[0] = counts[0][1:]
		acc2 = append(acc2, counts)
	}

	var ret int
	var ret2 int
	for i, v := range acc2 {
		if calcp1(v) {
			ret += i + 1
		}
		ret2 += calcp2(v)
	}

	fmt.Println(ret)
	fmt.Println(ret2)
}

//sum of ids of games for which it is possible
