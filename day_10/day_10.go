package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	north = iota
	south
	east
	west
)

func find_s(s []string) (int, int) {
	for i, v := range s {
		for j, v2 := range v {
			if v2 == 'S' {
				return i, j
			}
		}
	}
	// panic
	return -1, -1
}

func find_highest(s [][]int) int {
	highest := 0
	for _, v := range s {
		for _, v2 := range v {
			if v2 > highest {
				highest = v2
			}
		}
	}
	return highest
}

func count_pipes(s [][]int) int {
	count := 0
	for _, v := range s {
		for _, v2 := range v {
			if v2 != -1 {
				// fmt.Println("pipe at", i, j)
				count++
			}
		}
	}
	return count
}

func has_north(b byte) bool {
	return b == '|' || b == 'L' || b == 'J'
}

func has_south(b byte) bool {
	return b == '|' || b == '7' || b == 'F'
}

func has_east(b byte) bool {
	return b == '-' || b == 'L' || b == 'F'
}

func has_west(b byte) bool {
	return b == '-' || b == 'J' || b == '7'
}

func trigger_mini_bfs(visited [][]int, row, col int) int {
	// gives starter row and col
	if visited[row][col] != -1 {
		return 0
	}
	count := 0
	var queue [][]int
	queue = append(queue, []int{row, col})
	visited[row][col] = -2
	count++

	for len(queue) != 0 {
		row := queue[0][0]
		col := queue[0][1]
		queue = queue[1:]
		// fmt.Println(row, col)

		if row != 0 && visited[row-1][col] == -1 {
			queue = append(queue, []int{row - 1, col})
			visited[row-1][col] = -2
			count++
		}
		if row != len(visited)-1 && visited[row+1][col] == -1 {
			queue = append(queue, []int{row + 1, col})
			visited[row+1][col] = -2
			count++
		}
		if col != 0 && visited[row][col-1] == -1 {
			queue = append(queue, []int{row, col - 1})
			visited[row][col-1] = -2
			count++
		}
		if col != len(visited[row])-1 && visited[row][col+1] == -1 {
			queue = append(queue, []int{row, col + 1})
			visited[row][col+1] = -2
			count++
		}
	}

	//return the count
	return count
}

func main() {
	data, _ := os.ReadFile("./text.txt")

	data_lines := strings.Split(string(data), "\n")
	data_lines = data_lines[:len(data_lines)-1]

	visited := make([][]int, len(data_lines))
	for i := 0; i < len(data_lines); i++ {
		visited[i] = make([]int, len(data_lines[i]))
		for j := 0; j < len(data_lines[i]); j++ {
			visited[i][j] = -1
		}
	}

	var queue [][]int

	t1, t2 := find_s(data_lines)
	queue = append(queue, []int{t1, t2})
	visited[t1][t2] = 0

	for len(queue) != 0 {
		row := queue[0][0]
		col := queue[0][1]
		// fmt.Println(row, col, len(queue))
		queue = queue[1:]
		switch data_lines[row][col] {
		case '|':
			if row != 0 && visited[row-1][col] == -1 && has_south(data_lines[row-1][col]) {
				queue = append(queue, []int{row - 1, col})
				visited[row-1][col] = visited[row][col] + 1
			}
			if row != len(visited)-1 && visited[row+1][col] == -1 && has_north(data_lines[row+1][col]) {
				queue = append(queue, []int{row + 1, col})
				visited[row+1][col] = visited[row][col] + 1
			}
		case '-':
			if col != 0 && visited[row][col-1] == -1 && has_east(data_lines[row][col-1]) {
				queue = append(queue, []int{row, col - 1})
				visited[row][col-1] = visited[row][col] + 1
			}
			if col != len(visited[row])-1 && visited[row][col+1] == -1 && has_west(data_lines[row][col+1]) {
				queue = append(queue, []int{row, col + 1})
				visited[row][col+1] = visited[row][col] + 1
			}
		case 'L':
			if row != 0 && visited[row-1][col] == -1 && has_south(data_lines[row-1][col]) {
				queue = append(queue, []int{row - 1, col})
				visited[row-1][col] = visited[row][col] + 1
			}
			if col != len(visited[row])-1 && visited[row][col+1] == -1 && has_west(data_lines[row][col+1]) {
				queue = append(queue, []int{row, col + 1})
				visited[row][col+1] = visited[row][col] + 1
			}
		case 'J':
			if row != 0 && visited[row-1][col] == -1 && has_south(data_lines[row-1][col]) {
				queue = append(queue, []int{row - 1, col})
				visited[row-1][col] = visited[row][col] + 1
			}
			if col != 0 && visited[row][col-1] == -1 && has_east(data_lines[row][col-1]) {
				queue = append(queue, []int{row, col - 1})
				visited[row][col-1] = visited[row][col] + 1
			}
		case '7':
			if row != len(visited)-1 && visited[row+1][col] == -1 && has_north(data_lines[row+1][col]) {
				queue = append(queue, []int{row + 1, col})
				visited[row+1][col] = visited[row][col] + 1
			}
			if col != 0 && visited[row][col-1] == -1 && has_east(data_lines[row][col-1]) {
				queue = append(queue, []int{row, col - 1})
				visited[row][col-1] = visited[row][col] + 1
			}
		case 'F':
			if row != len(visited)-1 && visited[row+1][col] == -1 && has_north(data_lines[row+1][col]) {
				queue = append(queue, []int{row + 1, col})
				visited[row+1][col] = visited[row][col] + 1
			}
			if col != len(visited[row])-1 && visited[row][col+1] == -1 && has_west(data_lines[row][col+1]) {
				queue = append(queue, []int{row, col + 1})
				visited[row][col+1] = visited[row][col] + 1
			}
		case '.':
			// panic we should never be here
			panic(". in search case")
		case 'S':
			if row != 0 && visited[row-1][col] == -1 && has_south(data_lines[row-1][col]) {
				queue = append(queue, []int{row - 1, col})
				visited[row-1][col] = visited[row][col] + 1
			}
			if row != len(visited)-1 && visited[row+1][col] == -1 && has_north(data_lines[row+1][col]) {
				queue = append(queue, []int{row + 1, col})
				visited[row+1][col] = visited[row][col] + 1
			}
			if col != 0 && visited[row][col-1] == -1 && has_east(data_lines[row][col-1]) {
				queue = append(queue, []int{row, col - 1})
				visited[row][col-1] = visited[row][col] + 1
			}
			if col != len(visited[row])-1 && visited[row][col+1] == -1 && has_west(data_lines[row][col+1]) {
				queue = append(queue, []int{row, col + 1})
				visited[row][col+1] = visited[row][col] + 1
			}
		default:
			// panic
			panic("default in search case")
		}
	}

	fmt.Println(find_highest(visited))

	// Try to traverse loop instead
	// First try, pretend "inside" is to our right
	// row+col = t1,t2
	count := 0
	direction := north
	row := t1 - 1
	col := t2
	// direction := south
	// row := t1 + 1
	// col := t2
	for row != t1 || col != t2 {
		// fmt.Println(row, col, direction)

		// check besides
		switch direction {
		case north:
			if col != len(visited[row])-1 {
				count += trigger_mini_bfs(visited, row, col+1)
			}
			if row != 0 && data_lines[row][col] == '7' {
				count += trigger_mini_bfs(visited, row-1, col+1)
				count += trigger_mini_bfs(visited, row-1, col)
			}
		case south:
			if col != 0 {
				count += trigger_mini_bfs(visited, row, col-1)
			}
			if row != len(data_lines)-1 && data_lines[row][col] == 'L' {
				count += trigger_mini_bfs(visited, row+1, col-1)
				count += trigger_mini_bfs(visited, row+1, col)
			}
		case east:
			if row != len(visited)-1 {
				count += trigger_mini_bfs(visited, row+1, col)
			}
			if col != len(data_lines[row+1])-1 && data_lines[row][col] == 'J' {
				count += trigger_mini_bfs(visited, row+1, col+1)
				count += trigger_mini_bfs(visited, row, col+1)
			}
		case west:
			if row != 0 {
				count += trigger_mini_bfs(visited, row-1, col)
			}
			if col != 0 && data_lines[row][col] == 'F' {
				count += trigger_mini_bfs(visited, row-1, col-1)
				count += trigger_mini_bfs(visited, row, col-1)
			}
		}

		// find next thing
		switch direction {
		case north:
			switch data_lines[row][col] {
			case 'F':
				col++
				direction = east
			case '7':
				col--
				direction = west
			case '|':
				row--
			default:
				panic("invalid token when northward")
			}
		case south:
			switch data_lines[row][col] {
			case 'J':
				col--
				direction = west
			case 'L':
				col++
				direction = east
			case '|':
				row++
			default:
				panic("invalid token when southward")
			}
		case east:
			switch data_lines[row][col] {
			case '7':
				row++
				direction = south
			case 'J':
				row--
				direction = north
			case '-':
				col++
			default:
				panic("invalid token when eastbound")
			}
		case west:
			switch data_lines[row][col] {
			case 'F':
				row++
				direction = south
			case 'L':
				row--
				direction = north
			case '-':
				col--
			default:
				panic("invalid token when westbound")
			}
		}
	}
	fmt.Println(count)
}

// for each thing, can case on directions related to it

// "north" -> |, L or J
// "south" -> |, 7 or F
// "east" -> -, L or F
// "west" -> -, J or 7

// 671 too high
// 361 too low
// 5433 too high
// 367
//  pain
