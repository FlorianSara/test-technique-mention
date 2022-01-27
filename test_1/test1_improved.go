package main

import (
	"bufio"
	"fmt"
	"os"
)

// changing array to map to improve search time (O(1) instead of O(n))
type whitelist map[string]bool

func readWords(path string) whitelist {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	wl := make(map[string]bool)

	for scanner.Scan() {
		wl[scanner.Text()] = true
	}

	return wl
}

func (w whitelist) contains(needle string) bool {
	return w[needle]
}

func main() {

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		panic("missing argument")
	}

	validWords := readWords(path)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		if validWords.contains(w) {
			fmt.Fprintf(os.Stdout, "%s\n", w)
		} else {
			fmt.Fprintf(os.Stdout, "<%s>\n", w)
		}
	}
}
