package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	temp := make(map[string][]string)
	for _, str := range strs {
		ordered := getOrdered(str)
		if temp[ordered] == nil {
			temp[ordered] = []string{}
		}
		temp[ordered] = append(temp[ordered], str)
	}
	result := [][]string{}
	for _, v := range temp {
		result = append(result, v)
	}

	return result

}

func getOrdered(str string) string {
	tmp := make([]byte, 26)
	for _, c := range []rune(str) {
		tmp[c-'a']++
	}
	return string(tmp)

}
