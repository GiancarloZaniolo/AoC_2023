package main

import (
	"fmt"
	"os"
	"strings"
)

type pair struct {
	left  string
	right string
}

func parse_line(s string, m map[string]pair) {
	s_arr := strings.Split(s, " ")

	// fmt.Println(s_arr)

	var p pair

	p.left = s_arr[2][1:4]
	p.right = s_arr[3][0:3]

	m[s_arr[0]] = p
}

func assemble_map(elems []string) map[string]pair {

	m := make(map[string]pair)

	for _, v := range elems {
		parse_line(v, m)
	}

	return m
}

func verify_incomplete(curr_vals []string, count int) bool {
	res := true
	for i, v := range curr_vals {
		if v[2] != 'Z' {
			res = false
		} else {
			fmt.Println("Thing num ", i, " at count ", count)
		}
	}
	return res
}

func main() {
	data, _ := os.ReadFile("./text.txt")

	data_lines := strings.Split(string(data), "\n")
	data_lines = data_lines[:len(data_lines)-1]
	directions := data_lines[0]

	m := assemble_map(data_lines[2:])

	for key, val := range m {
		fmt.Println("key ", key, "  val ", val)
	}

	curr_val := "AAA"
	steps := 0
	for ; curr_val != "ZZZ"; steps++ {
		direction := directions[steps%len(directions)]
		if direction == 'L' {
			curr_val = m[curr_val].left
		} else {
			curr_val = m[curr_val].right
		}
	}
	fmt.Println(steps)

	// // create starter list
	// var curr_val_list []string

	// for key, _ := range m {
	// 	if key[2] == 'A' {
	// 		curr_val_list = append(curr_val_list, key)
	// 	}
	// }

	// // fmt.Println(curr_val_list)
	// steps := 0
	// for ; !verify_incomplete(curr_val_list, steps); steps++ {
	// 	direction := directions[steps%len(directions)]
	// 	for i, v := range curr_val_list {
	// 		if direction == 'L' {
	// 			// fmt.Printf("%s becomes %s\n", v, m[v].left)
	// 			curr_val_list[i] = m[v].left
	// 		} else {
	// 			// fmt.Printf("%s becomes %s\n", v, m[v].right)
	// 			curr_val_list[i] = m[v].right
	// 		}
	// 	}
	// }
	// fmt.Println(steps)

	fmt.Println(21083806112641)
}

/*

1: 12599 25198 37797 50396
   12599 12599 -> always?
0: 19631 39262 58893
   19631 19631
2: 17873
3: 20803
4: 21389
5: 23147

brooooo...

very special case, but its the lcm of all of these numbers :/

*/
