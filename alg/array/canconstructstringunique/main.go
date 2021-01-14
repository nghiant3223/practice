package canconstructstringunique

import "strings"

// Given an target string and a word bank, check whether the target word can be constructed from words in the word bank

var visited map[string]struct{}

func canConstructUnique(target string, banks []string) bool {
	if target == "" {
		return true
	}
	for _, word := range banks {
		if _, ok := visited[word]; ok {
			continue
		}
		if !strings.HasPrefix(target, word) {
			continue
		}

		visited[word] = struct{}{}
		constructable := canConstructUnique(target[len(word):], banks)
		delete(visited, word)

		if constructable {
			return true
		}
	}
	return false
}
