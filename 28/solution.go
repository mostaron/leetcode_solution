package main

func main() {

}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		matched := true
		for j := 0; j < len(needle); j++ {
			if needle[j] == haystack[i+j] {
				continue
			}
			matched = false
			break
		}
		if matched {
			return i
		}
	}
	return -1
}
