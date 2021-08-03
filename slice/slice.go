package main

import "fmt"

func slice_num() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("<<slice_num()>>")
	fmt.Println()

	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	fmt.Println()
}

func slice_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)

	fmt.Println("<<slice_append()>>")
	fmt.Println()

	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	fmt.Println()
}

func slice_cap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("<<slice_cap()>>")
	fmt.Println()

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
}

func slice_len_cap() {
	nums := make([]int, 3, 5) // len: 3, cap: 5

	/*
		nums := make([]int, 5)	// cap: 5
		nums = nums[:3]			// len: 3
	*/
}

func main() {
	slice_num()
	slice_append()
	slice_cap()
}
