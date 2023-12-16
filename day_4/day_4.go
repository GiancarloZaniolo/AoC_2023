package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	nums    []int
	winners []int
}

func construct_nums(s string) []int {
	temp := strings.Split(s, ":")
	return construct_winners(temp[1])
}

func construct_winners(s string) []int {
	s = strings.ReplaceAll(s, "  ", " ")
	s = s[1:]
	temp := strings.Split(s, " ")
	var acc []int
	for _, v := range temp {
		temp, _ := strconv.Atoi(v)
		acc = append(acc, temp)
	}
	return acc
}

func construct_card(s string) card {
	var card card
	temp := strings.Split(s, "|")
	card.nums = construct_nums(temp[0])
	card.winners = construct_winners(temp[1])
	return card
}

func print_card(c card) {
	fmt.Println("nums ", c.nums, "  winners ", c.winners)
}

func calc_score(c card) int {
	this_winnner := 0
	for _, num := range c.nums {
		for _, winner := range c.winners {
			if num == winner {
				if this_winnner == 0 {
					this_winnner = 1
				} else {
					this_winnner *= 2
				}
			}
		}
	}
	return this_winnner
}

func num_winners(c card) int {
	win_count := 0
	for _, num := range c.nums {
		for _, winner := range c.winners {
			if num == winner {
				win_count++
			}
		}
	}
	return win_count
}

func cards_iteration(unproc_cards []int, winner_count []int) int {
	cards := len(winner_count)
	for i := 0; i < len(winner_count); i++ {
		for j := 1; i+j < len(winner_count) && j <= winner_count[i]; j++ {
			// fmt.Printf("Card %d gets %d more\n", i+j+1, unproc_cards[i])
			unproc_cards[i+j] += unproc_cards[i]
			cards += unproc_cards[i]
		}
	}
	return cards
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

	var cards []card
	for _, v := range acc {
		cards = append(cards, construct_card(v))
	}

	var ret int
	for _, v := range cards {
		ret += calc_score(v)
	}
	fmt.Println(ret)

	var winner_count []int
	var unproc_cards []int
	for _, v := range cards {
		winner_count = append(winner_count, num_winners(v))
		unproc_cards = append(unproc_cards, 1)
	}
	// var ret2 int
	// for cards_iteration(&unproc_cards, winner_count, &ret2) {
	// }
	fmt.Println(cards_iteration(unproc_cards, winner_count))

}
