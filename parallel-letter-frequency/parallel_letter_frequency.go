package letter

import "maps"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	le := len(texts)
	ch := make(chan FreqMap, le)
	for _, value := range texts {
		go func() {
			ch <- Frequency(value)
		}()
	}
	freqMap := FreqMap{}
	for i := 0; i < le; i++ {
		out := <-ch
		maps.Copy(freqMap, out)
	}
	return freqMap
}
