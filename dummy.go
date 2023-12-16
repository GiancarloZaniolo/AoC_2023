package dummy

// package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("./text.txt")

	defer f.Close()

	s := bufio.NewScanner(f)

	var acc []string
	for s.Scan() {
		data := s.Text()
		acc = append(acc, data)
	}
}
