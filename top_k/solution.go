package top_k

import "log"

func topKFrequent(nums []int, k int) []int {
	// step1 put all eleemnt into hashMap
	freq := make(map[int]int)
	bucket := make([][]int, len(nums))
	for _, val := range nums {
		freq[val] += 1
	}
	log.Println(bucket)
	// step2 透過 priority queue 來實作 sort
	for key, count := range freq {
		bucket[count-1] = append(bucket[count-1], key)
	}

	var result []int
	for i := len(bucket) - 1; i >= 0; i-- {
		list := bucket[i]
		if len(list) > 0 {
			for idx := 0; idx <= len(list)-1; idx++ {
				result = append(result, list[idx])
				if len(result) == k {
					return result
				}
			}
		}
	}
	return result
}
