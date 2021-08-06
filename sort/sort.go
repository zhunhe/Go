package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := make([]int, 0)
	sli = append(sli, 3)
	sli = append(sli, 2)
	sli = append(sli, 6)
	sli = append(sli, 8)
	sli = append(sli, 10)
	sli = append(sli, 7)
	sli = append(sli, 1)
	sli = append(sli, 9)

	fmt.Println("before", sli)
	// 오름차순 정렬
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	fmt.Println("after", sli)
}
