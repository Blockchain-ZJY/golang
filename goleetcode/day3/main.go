package main

import (
	"fmt"
)

// 6. Z-Transfer
func convert(s string, numRows int) string {
	index := make([]int, len(s))
	counter := 1
	flip := true
	for j := 0; j < len(s); j++ {

		if numRows == 1 {

			return s
		}
		if flip {
			index[j] = counter
			counter++
		} else {
			index[j] = counter
			counter--
		}
		if counter == numRows || counter == 1 {
			flip = !flip
		}
	}
	fmt.Println(index)
	bans := make([]byte, len(s))
	index111 := 0
	for counter := 1; counter <= numRows; counter++ {
		for i := 0; i < len(s); i++ {
			if index[i] == counter {
				bans[index111] = s[i]
				index111++
			}
		}
	}
	return string(bans)
}

func main() {
	fmt.Println(convert("AB", 1))
}
