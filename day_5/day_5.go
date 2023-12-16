package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mapping struct {
	start   int
	bound   int
	mapping int
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func get_seeds(s string) []int {
	seeds_strings := (strings.Split(s, " "))[1:]
	var seed_ints []int
	for _, e := range seeds_strings {
		temp, _ := strconv.Atoi((e))
		seed_ints = append(seed_ints, temp)
	}
	return seed_ints
}

func parse_one_map(s string) mapping {
	temp := strings.Split(s, " ")
	var this_map mapping
	this_map.start, _ = strconv.Atoi(temp[1])
	this_map.bound, _ = strconv.Atoi(temp[2])
	this_map.mapping, _ = strconv.Atoi(temp[0])
	return this_map
}

func parse_full_map(s string) func(int) int {
	elems := strings.Split(s, "\n")
	var mapping_list []mapping
	for i := 1; i < len(elems); i++ {
		mapping_list = append(mapping_list, parse_one_map(elems[i]))
	}

	mapping_fun := func(num int) int {
		// fmt.Printf("testing mapping for %d\n", num)
		for _, v := range mapping_list {
			// fmt.Printf("testing start %d bound %d map %d\n", v.start, v.bound, v.mapping)
			if num >= v.start && num < v.start+v.bound {
				return v.mapping + (num - v.start)
			}
		}
		return num
	}
	return mapping_fun
}

func apply_all_maps(start int, map_funs [](func(num int) int)) int {
	// fmt.Printf("doing seed %d\n", start)
	for _, v := range map_funs {
		start = v(start)
		// fmt.Printf("now %d  \n", start)
	}
	// fmt.Println()
	return start
}

func main() {
	data, _ := os.ReadFile("./text.txt")

	maps := strings.Split(string(data), "\n\n")

	seeds := get_seeds(maps[0])

	var mapfun_list [](func(num int) int)
	for _, v := range maps {
		mapfun_list = append(mapfun_list, parse_full_map(v))
	}

	min_loc := 9999999999999
	for _, v := range seeds {
		min_loc = min(apply_all_maps(v, mapfun_list), min_loc)
	}

	fmt.Println(min_loc)

	// Just wait for a long time lmao
	// That being said, it took less time for it to run than it took me to
	//  implement a memoizing version
	min_loc2 := 9999999999999
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			min_loc2 = min(apply_all_maps(j, mapfun_list), min_loc2)
		}
	}
	fmt.Println(min_loc2)
}
