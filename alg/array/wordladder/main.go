package main

import (
	"log"
	"sort"
)

func main() {
	//beginWord := "hit"
	//endWord := "cog"
	words := []string{"hot", "dot", "log", "dog", "lot", "cog"}
	sort.Strings(words)
	log.Println(words)
}
