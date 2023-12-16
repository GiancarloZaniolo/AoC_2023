package main

import (
	"bufio"
	"fmt"
	"os"
)

func byte_is_num(b byte) bool {
	return b >= '0' && b <= '9'
}

func is_valid_num(arr [][]byte, row int, col int) bool {
	var sec_len int
	for ; col+sec_len < len(arr[0]) && byte_is_num(arr[row][col+sec_len]); sec_len++ {
	}
	if row != 0 {
		// Scan above for chars
		for i := 0; i < sec_len; i++ {
			if arr[row-1][col+i] != '.' {
				return true
			}
		}
	}
	if row != len(arr)-1 {
		// scan below for chars
		for i := 0; i < sec_len; i++ {
			if arr[row+1][col+i] != '.' {
				return true
			}
		}
	}
	if col != 0 {
		if (row != 0 && arr[row-1][col-1] != '.') ||
			(arr[row][col-1] != '.') ||
			(row != len(arr)-1 && arr[row+1][col-1] != '.') {
			return true
		}
	}
	// fmt.Printf("row %d  col %d  sec_len %d\n", row, col, sec_len)
	if col+sec_len < len(arr[0]) {
		if (row != 0 && arr[row-1][col+sec_len] != '.') ||
			(arr[row][col+sec_len] != '.') ||
			(row < len(arr)-1 && arr[row+1][col+sec_len] != '.') {
			return true
		}
	}
	// fmt.Println("after last checks")
	return false
}

func retrieve_num(str []byte, start int) int {
	acc := 0
	for i := 0; start+i < len(str) && byte_is_num(str[start+i]); i++ {
		acc *= 10
		acc += int(str[start+i] - '0')
	}
	return acc
}

type coord struct {
	row int
	col int
}

func add_map_coord(num int, row int, col int, m map[coord]([]int)) {
	key := coord{row, col}
	m[key] = append(m[key], num)
}

func add_gear_num(arr [][]byte, row int, col int, m map[coord]([]int)) bool {
	curr_num := retrieve_num(arr[row], col)
	var sec_len int
	for ; col+sec_len < len(arr[0]) && byte_is_num(arr[row][col+sec_len]); sec_len++ {
	}
	if row != 0 {
		// Scan above for chars
		for i := 0; i < sec_len; i++ {
			if arr[row-1][col+i] == '*' {
				add_map_coord(curr_num, row-1, col+i, m)
			}
		}
	}
	if row != len(arr)-1 {
		// scan below for chars
		for i := 0; i < sec_len; i++ {
			if arr[row+1][col+i] == '*' {
				add_map_coord(curr_num, row+1, col+i, m)
			}
		}
	}
	if col != 0 {
		if row != 0 && arr[row-1][col-1] == '*' {
			add_map_coord(curr_num, row-1, col-1, m)
		}
		if arr[row][col-1] == '*' {
			add_map_coord(curr_num, row, col-1, m)
		}
		if row != len(arr)-1 && arr[row+1][col-1] == '*' {
			add_map_coord(curr_num, row+1, col-1, m)
		}
	}
	if col+sec_len < len(arr[0]) {
		if row != 0 && arr[row-1][col+sec_len] == '*' {
			add_map_coord(curr_num, row-1, col+sec_len, m)
		}
		if arr[row][col+sec_len] == '*' {
			add_map_coord(curr_num, row, col+sec_len, m)
		}
		if row < len(arr)-1 && arr[row+1][col+sec_len] == '*' {
			add_map_coord(curr_num, row+1, col+sec_len, m)
		}
	}
	// fmt.Println("after last checks")
	return false
}

func main() {
	f, _ := os.Open("./text.txt")

	defer f.Close()

	s := bufio.NewScanner(f)

	var acc [][]byte
	for s.Scan() {
		data := s.Text()
		acc = append(acc, []byte(data))
	}

	// fmt.Printf("got %t\n", is_valid_num(acc, 0, 5))

	var ret int

	for row, v := range acc {
		i := 0
		is_num := false
		for ; i < len(v); i++ {
			if byte_is_num(v[i]) && is_num == false {
				if is_valid_num(acc, row, i) {
					// fmt.Printf("adding %d\n", retrieve_num(v, i))
					ret += retrieve_num(v, i)
				}
				is_num = true
			} else if is_num && !byte_is_num(v[i]) {
				is_num = false
			}
			// fmt.Printf("current state row %d col %d  char %c end %t\n", row, i, v[i], is_num)
		}
	}
	fmt.Println(ret)

	m := make(map[coord]([]int))

	for row, v := range acc {
		i := 0
		is_num := false
		for ; i < len(v); i++ {
			if byte_is_num(v[i]) && is_num == false {
				add_gear_num(acc, row, i, m)
				is_num = true
			} else if is_num && !byte_is_num(v[i]) {
				is_num = false
			}
			// fmt.Printf("current state row %d col %d  char %c end %t\n", row, i, v[i], is_num)
		}
	}
	var ret2 int
	for _, value := range m {
		// fmt.Printf("key row:%d  col:%d  arr: ", key.row, key.col)
		// fmt.Println(value)
		if len(value) == 2 {
			ret2 += value[0] * value[1]
		}
	}
	fmt.Println(ret2)

}

// add itself to every gear it is adjacent to
