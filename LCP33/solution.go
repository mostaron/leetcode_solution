package main

import (
	"fmt"
)

func main() {
	fmt.Println(storeWater([]int{1, 3}, []int{6, 8}))
	fmt.Println(storeWater([]int{9, 0, 1}, []int{0, 2, 2}))
	fmt.Println(storeWater([]int{3, 2, 5}, []int{0, 0, 0}))
	fmt.Println(storeWater([]int{0}, []int{1}))
	fmt.Println(storeWater([]int{16, 29, 42, 70, 42, 9}, []int{89, 44, 50, 90, 94, 91}))
}

func storeWater(bucket []int, vat []int) int {
	dict := make(map[int]map[int]interface{})
	incr := 0
	max := 0
	for i := 0; i < len(bucket); i++ {
		if vat[i] == 0 {
			continue
		}
		if bucket[i] == 0 {
			bucket[i] = 1
			incr++
		}
		times := vat[i] / bucket[i]
		if vat[i]%bucket[i] != 0 {
			times++
		}
		if len(dict[times]) == 0 {
			dict[times] = make(map[int]interface{})
		}
		dict[times][i] = nil
		if times > max {
			max = times
		}
	}
	init := incr + max

	for init >= incr+max && max > 1 {
		init = incr + max
		for i := range dict[max] {
			bucket[i]++
			incr++
			times := vat[i] / bucket[i]
			if vat[i]%bucket[i] != 0 {
				times++
			}
			delete(dict[max], i)
			if len(dict[times]) == 0 {
				dict[times] = make(map[int]interface{})
			}
			dict[times][i] = nil
		}
		if len(dict[max]) == 0 {
			delete(dict, max)
		}
		max = 0
		for k, _ := range dict {
			if k > max {
				max = k
			}
		}
	}
	return init
}
