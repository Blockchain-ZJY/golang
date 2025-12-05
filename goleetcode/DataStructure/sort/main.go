package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min_i := i
		minv := 100
		for j := i; j < n; j++ { //找到从i到n的最小值
			if arr[j] < minv {
				minv = arr[j]
				min_i = j
			}
		}
		arr[i], arr[min_i] = arr[min_i], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i] // 要插入的数据
		j := i
		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j -= 1
		}
		arr[j] = temp
	}

	return arr
}

// 合并两个数组
func merge(left []int, right []int) []int {
	// 1. 初始化长度为 0，容量为两个数组之和，避免扩容开销
	newarr := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			newarr = append(newarr, left[i])
			i++
		} else {
			newarr = append(newarr, right[j])
			j++
		}
	}

	// 2. & 3. 直接追加剩余部分
	// 如果 i 还没走到头，说明 left 还有剩余
	newarr = append(newarr, left[i:]...)
	// 如果 j 还没走到头，说明 right 还有剩余
	newarr = append(newarr, right[j:]...)

	return newarr
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	// l 从 1 开始，因为 0 是 pivot
	l, r := 1, len(arr)-1
	for l <= r {
		if arr[l] <= pivot {
			l++
		} else {
			arr[l], arr[r] = arr[r], arr[l]
			r--
		}
	}
	// 将 pivot 放到正确的位置 (r 的位置)
	arr[0], arr[r] = arr[r], arr[0]

	quickSort(arr[:r])
	quickSort(arr[r+1:])

	return arr
}

func search(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left+1)/2
		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if arr[left] == target {
		return left
	}
	return -1
}

//
func main() {
	// 各种sort
	fmt.Println(bubbleSort([]int{5, 9, 1, 6, 8}))
	fmt.Println(selectionSort([]int{5, 9, 1, 6, 8}))
	fmt.Println(insertionSort([]int{5, 9, 1, 6, 8}))
	fmt.Println(mergeSort([]int{5, 9, 1, 6, 8}))
	fmt.Println(quickSort([]int{5, 9, 1, 6, 8}))
	fmt.Println(search([]int{5, 9, 1, 6, 8}, 6))
}
