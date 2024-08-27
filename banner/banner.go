package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☻", 6)

	s := "G☻"
	fmt.Println("len: ", len(s))
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r) // rune (int32)
		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b) // byte (uint8)

	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y) // use #v in debug/log

	fmt.Printf("%20s!\n", s)

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("g☻g", isPalindrome("g☻g"))
}

func isPalindrome(s string) bool {
	// for i := 0; i < len(s)/2; i++ {
	// 	if s[i] != s[len(s)-i-1] {
	// 		return false
	// 	}
	// }
	rs := []rune(s) // to handle unicode characters
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}

	return true
}

func banner(text string, width int) {
	// padding := (width - len(text)) / 2 // BUG: len is in bytes
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
