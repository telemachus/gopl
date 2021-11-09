package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type wordFreq struct {
	word  string
	count int
}

type freqList []wordFreq

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	freqMap := make(map[string]int)

	for scanner.Scan() {
		freqMap[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "wordfreq:", err)
		os.Exit(1)
	}

	frequencies := make(freqList, len(freqMap))
	i := 0
	for w, c := range freqMap {
		frequencies[i] = wordFreq{word: w, count: c}
		i++
	}
	// sort.Sort(sort.Reverse(frequencies))
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].count > frequencies[j].count
	})

	for _, f := range frequencies {
		fmt.Printf("%8d: %s\n", f.count, f.word)
	}
}
