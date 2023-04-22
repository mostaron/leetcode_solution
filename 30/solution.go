package main

import "fmt"

/*
*
"barfoothefoobarman"
["foo","bar"]
"wordgoodgoodgoodbestword"
["word","good","best","word"]
"barfoofoobarthefoobarman"
["bar","foo","the"]
"wordgoodgoodgoodbestword"
["word","good","best","good"]
"ababababab"
["ababa","babab"]
*/
func main() {
	fmt.Println(findSubstring("ababababab", []string{"ababa", "babab"}))
}
func findSubstring(s string, words []string) []int {
	result := []int{}
	for i := 0; i <= len(s)-len(words)*len(words[0]); i++ {
		dict := make(map[string]int)
		for _, word := range words {
			dict[word]++
		}
		isMatch := true
		for j := 0; j < len(words); j++ {
			query := s[i+(j*len(words[0])) : i+((j+1)*len(words[0]))]
			if dict[query] > 0 {
				dict[query]--
			} else {
				isMatch = false
				break
			}
			if dict[query] == 0 {
				delete(dict, query)
			}
			if isMatch && len(dict) == 0 {
				result = append(result, i)
			}
		}

	}
	return result
}
