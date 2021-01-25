package sortcharbyfrequency

import "sort"

type Pair struct {
	Key byte
	Value int
}

func frequencySort(str string) string {
	m := buildPairMap(str)
	s := buildPairSlice(m)
	sort.Slice(s, func(i, j int) bool {
		return s[i].Value > s[j].Value
	})
	i := 0
	r := make([]byte, len(str))
	for _, p := range s {
		for j := 0; j < p.Value; j++ {
			r[i] = p.Key
			i++
		}
	}
	return string(r)
}

func buildPairMap(s string) map[byte]Pair {
	m := make(map[byte]Pair)
	for i := range s {
		char := s[i]
		pair, ok := m[char]
		if !ok {
			pair.Key = char
			pair.Value = 1
		} else {
			pair.Value++
		}
		m[char] = pair
	}
	return m
}

func buildPairSlice(m map[byte]Pair) []Pair {
	i := 0
	s := make([]Pair, len(m))
	for _, pair := range m {
		s[i] = pair
		i++
	}
	return s
}
