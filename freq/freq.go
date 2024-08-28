package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer file.Close()

	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(w)
	// mapDemo()

}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

/* Will run before main
func init() {
}
*/

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

func mapDemo() {
	var stocks map[string]float64 // symbol -> price
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	// stocks = make(map[string]float64)
	// stocks[sym] = 136.73
	stocks = map[string]float64{
		sym:    136.73,
		"AAPL": 172.35,
	}

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	for k := range stocks {
		fmt.Println(k)
	}
	for k, v := range stocks {
		fmt.Println(k, "->", v)
	}
	for _, v := range stocks {
		fmt.Println(v)
	}
	delete(stocks, "AAPL")
	fmt.Println(stocks)
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}
	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}
	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}

	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}
