package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func remove_spaces(s []string) []string {
	var ret []string
	for _, v := range s {
		if v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}

func count_ways(time int, distance int) int {
	ways := 0
	for i := 0; i < time; i++ {
		if (time-i)*i > distance {
			ways++
		}
	}
	return ways
}

func best_dist(time int, distance int) int {
	best := 0
	for i := 0; i < time; i++ {
		best = max(best, (time-i)*i)
	}
	return best
}

func main() {
	data, _ := os.ReadFile("./text.txt")

	data2 := strings.Split(string(data), "\n")
	times := remove_spaces((strings.Split(data2[0], " "))[1:])
	distances := remove_spaces((strings.Split(data2[1], " "))[1:])

	ret := 1
	for i := 0; i < len(times); i++ {
		temp1, _ := strconv.Atoi(times[i])
		temp2, _ := strconv.Atoi(distances[i])
		// fmt.Printf("way %d: %d\n", i, count_ways(temp1, temp2))
		ret *= count_ways(temp1, temp2)
	}
	fmt.Println(ret)

	re := regexp.MustCompile("[^0-9]")
	time_s := re.ReplaceAllString(data2[0], "")
	time, _ := strconv.Atoi(time_s)
	distance_s := re.ReplaceAllString(data2[1], "")
	distance, _ := strconv.Atoi(distance_s)

	// fmt.Println(time, distance)
	fmt.Println(count_ways(time, distance))
}

// 61709063 too high
