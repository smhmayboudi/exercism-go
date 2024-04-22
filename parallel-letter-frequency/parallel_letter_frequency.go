package letter

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
	ch := make(chan FreqMap)
	for _, value := range texts {
		go func(v string) {
			ch <- Frequency(v)
		}(value)
	}
	freqMap := FreqMap{}
	for range texts {
		for letter, count := range <-ch {
			freqMap[letter] += count
		}
	}

	return freqMap
}
