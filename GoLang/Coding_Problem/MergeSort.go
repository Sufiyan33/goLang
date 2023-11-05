package main

/*
Firstly, in merge sort, we keep dividing our array recursively into the right side and the left side
	and call the MergeSort function on both sides from line 14 to line 16.
	Now in next section we will add go-routines and channels for the same merge sort
*/
func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2

	left := mergeSort(data[:mid])
	right := mergeSort(data[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	merging := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merging, right...)
		} else if len(right) == 0 {
			return append(merging, left...)
		} else if left[0] < right[0] {
			merging = append(merging, left[0])
			left = left[1:]
		} else {
			merging = append(merging, right[0])
			right = right[1:]
		}
	}
	return merging
}
