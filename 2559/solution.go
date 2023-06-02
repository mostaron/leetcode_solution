package main

import "fmt"

func main() {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
}

func vowelStrings(words []string, queries [][]int) []int {
	hash := make([]int, len(words)+1)

	for i, word := range words {
		hash[i+1] = hash[i]
		if isVowel(word) {
			hash[i+1]++
		}
	}
	result := make([]int, len(queries))
	for i, query := range queries {
		result[i] = hash[query[1]+1] - hash[query[0]]
	}
	return result
}

func isVowel(word string) bool {
	start, end := word[0], word[len(word)-1]
	return (start == 'a' || start == 'e' || start == 'i' || start == 'o' || start == 'u') &&
		(end == 'a' || end == 'e' || end == 'i' || end == 'o' || end == 'u')
}
