# top_k_frequent

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

## 分析

首先是要紀錄每個直出現的次數 --> 透過 hashmap

第二個需要對這些紀錄出來的做排序 --> 題目要求要做到 O(nlogn) 直覺上可以透過 MaxHeap 來做處理

取出前 K 個最大

## 優化

觀察題目所要求的是出現次數最多

輸入 array 如果是 n, 那出現最多次數只有 n

所以可以透過類似 bucket sort 一樣，把出現次數當作 key 對應到所有出現次數一樣的值

舉例來說：

nums:[] int{1, 1, 1, 2, 2, 3, 3, 4}

bucket: [0] -> {}
        [1] -> {4} //出現一次
        [2] -> {2, 3} // 出現兩次
        [3] -> {1} // 出現 3 次
        [4] -> {}
        [5] -> {}
        [6] -> {}
        [7] -> {}
        [8] -> {}

這樣在排序時， 就可以透過從最大出現的次數往後面找

所以這樣搜尋方式是 O(n)

雖然在實作 for loop Nested 深度為 2

## 初步實作

```golang
import "container/heap"

func topKFrequent(nums []int, k int) []int {
    // step1 put all eleemnt into hashMap
    freq := make(map[int]int)
    for _, val := range nums {
      freq[val] += 1
    }
    
    // step2 透過 priority queue 來實作 sort
    q := PriorityQueue{}
    for key, count := range freq {
        heap.Push(&q, &Item{key: key, count: count})        
    }
    
    var result []int
    for len(result) < k {
        item := heap.Pop(&q).(*Item)
        result = append(result, item.key)
    }
    return result
}


type Item struct {
    key, count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item
func (pq PriorityQueue) Len() int {
    return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
   return pq[i].count > pq[j].count
}
func (pq PriorityQueue) Swap(i, j int) {
   pq[i], pq[j] = pq[j], pq[i]
}
// Push define
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    *pq = append(*pq, item)
}
// Pop define
func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    item := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return item
}
```

## 使用 bucket 來做優化

```golang
package top_k

import "log"

func topKFrequent(nums []int, k int) []int {
	// step1 put all element into hashMap
	freq := make(map[int]int)
	bucket := make([][]int, len(nums))
	for _, val := range nums {
		freq[val] += 1
	}
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
```