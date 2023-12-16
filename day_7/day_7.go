package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Just make a comparison function that follows all the rules

type hand struct {
	cards string
	bet   int
}

type ByCards []hand

func (a ByCards) Len() int           { return len(a) }
func (a ByCards) Less(i, j int) bool { return !compare_hands(a[i], a[j]) }
func (a ByCards) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func kind_count(h hand) (int, int) {
	m := make(map[byte]int)
	for i := 0; i < len(h.cards); i++ {
		val, exists := m[h.cards[i]]
		if !exists {
			val = 0
		}
		m[h.cards[i]] = val + 1
	}
	max_count := 0
	for _, v := range m {
		max_count = max(max_count, v)
	}
	return len(m), max_count
}

func value_card(b1 byte) int {
	switch b1 {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return -1
	}
}

func high_card(h1, h2 hand) bool {
	for i := 0; i < len(h1.cards); i++ {
		cmp := value_card(h1.cards[i]) - value_card(h2.cards[i])
		switch {
		case cmp < 0:
			return false
		case cmp > 0:
			return true
		}
	}
	return true
}

func compare_hands(hand1 hand, hand2 hand) bool {
	kinds1, count1 := kind_count(hand1)
	kinds2, count2 := kind_count(hand2)

	// Five of a kind
	if kinds1 == 1 {
		if kinds2 == 1 {
			return high_card(hand1, hand2)
		} else {
			return true
		}
	} else if kinds2 == 1 {
		return false
	}

	// Four of a kind
	if count1 == 4 {
		if count2 == 4 {
			return high_card(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 4 {
		return false
	}

	// Full house
	if count1 == 3 && kinds1 == 2 {
		if count2 == 3 && kinds2 == 2 {
			return high_card(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 3 && kinds2 == 2 {
		return false
	}

	// Three of a kind
	if count1 == 3 {
		if count2 == 3 {
			return high_card(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 3 {
		return false
	}

	// Two pair
	if kinds1 == 3 && count1 == 2 {
		if kinds2 == 3 && count2 == 2 {
			return high_card(hand1, hand2)
		} else {
			return true
		}
	} else if kinds2 == 3 && count2 == 2 {
		return false
	}

	// One pair
	if count1 == 2 {
		if count2 == 2 {
			return high_card(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 2 {
		return false
	}

	return high_card(hand1, hand2)
}

func parse_line(s string) hand {
	// fmt.Println("parsing ", s)
	elems := strings.Split(s, " ")
	var one_hand hand
	one_hand.cards = elems[0]
	one_hand.bet, _ = strconv.Atoi(elems[1])
	return one_hand
}

// =============================================================================

type ByCards2 []hand

func (a ByCards2) Len() int           { return len(a) }
func (a ByCards2) Less(i, j int) bool { return !compare_hands2(a[i], a[j]) }
func (a ByCards2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func find_j_replacement(s string) byte {
	// Create frequency map
	m := make(map[byte]int)
	for i, _ := range s {
		val, exists := m[s[i]]
		if !exists {
			val = 0
		}
		m[s[i]] = val + 1
	}

	// Create map containing only highest frequency
	best := 0
	var m2 map[byte]bool
	for k, v := range m {
		if k != 'J' {
			if v > best {
				m2 = make(map[byte]bool)
				m2[k] = true
				best = v
			} else if v == best {
				m2[k] = true
			}
		}

	}

	// Find highest character of map
	var best_key byte
	best_key = byte(0)
	for k, _ := range m2 {
		if value_card(k) > value_card(byte(best_key)) {
			best_key = k
		}
	}
	if best_key == 0 {
		return 'A'
	}
	return byte(best_key)
}

func do_j_replacement(s string) string {
	replacement := string(find_j_replacement(s))
	// fmt.Println("replacement is ", replacement)
	return strings.Replace(s, "J", replacement, -1)
}

func value_card2(b1 byte) int {
	switch b1 {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 1
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return -1
	}
}

func high_card2(h1, h2 hand) bool {
	for i := 0; i < len(h1.cards); i++ {
		cmp := value_card2(h1.cards[i]) - value_card2(h2.cards[i])
		switch {
		case cmp < 0:
			return false
		case cmp > 0:
			return true
		}
	}
	return true
}

func compare_hands2(hand1 hand, hand2 hand) bool {
	hand1b := hand1
	hand1b.cards = do_j_replacement(hand1.cards)
	hand2b := hand2
	hand2b.cards = do_j_replacement(hand2.cards)
	kinds1, count1 := kind_count(hand1b)
	kinds2, count2 := kind_count(hand2b)

	// Five of a kind
	if kinds1 == 1 {
		if kinds2 == 1 {
			return high_card2(hand1, hand2)
		} else {
			return true
		}
	} else if kinds2 == 1 {
		return false
	}

	// Four of a kind
	if count1 == 4 {
		if count2 == 4 {
			return high_card2(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 4 {
		return false
	}

	// Full house
	if count1 == 3 && kinds1 == 2 {
		if count2 == 3 && kinds2 == 2 {
			return high_card2(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 3 && kinds2 == 2 {
		return false
	}

	// Three of a kind
	if count1 == 3 {
		if count2 == 3 {
			return high_card2(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 3 {
		return false
	}

	// Two pair
	if kinds1 == 3 && count1 == 2 {
		if kinds2 == 3 && count2 == 2 {
			return high_card2(hand1, hand2)
		} else {
			return true
		}
	} else if kinds2 == 3 && count2 == 2 {
		return false
	}

	// One pair
	if count1 == 2 {
		if count2 == 2 {
			return high_card2(hand1, hand2)
		} else {
			return true
		}
	} else if count2 == 2 {
		return false
	}

	return high_card2(hand1, hand2)
}

func main() {
	data, _ := os.ReadFile("./text.txt")

	data_arr := strings.Split(string(data), "\n")
	data_arr = data_arr[:len(data_arr)-1]

	var hand_list []hand
	for _, v := range data_arr {
		hand_list = append(hand_list, parse_line(v))
	}

	sort.Sort(ByCards(hand_list))

	// fmt.Println(hand_list)

	ret1 := 0
	for i := 0; i < len(hand_list); i++ {
		ret1 += hand_list[i].bet * (i + 1)
	}

	fmt.Println(ret1)

	// Do j replacement

	// for i := 0; i < len(hand_list); i++ {
	// 	hand_list[i].cards = do_j_replacement(hand_list[i].cards)
	// }

	// fmt.Println(hand_list)

	sort.Sort(ByCards2(hand_list))

	// fmt.Println(hand_list)

	ret2 := 0
	for i := 0; i < len(hand_list); i++ {
		ret2 += hand_list[i].bet * (i + 1)
	}

	fmt.Println(ret2)

}

// 241455695 too low
// 241700715 too low
// 241514165 must also be too low
// 242473900 too low
// 242236747 must also be too low
// 242903680 "not the right answer"
// 243101568

// Tody was a bit of a travesty because I did not read the instructions
