package main

/*
Now we have to make sure that Merge(left,right) is executed after we get return values from both the recursive calls, i.e.
	both the left and right must be updated before Merge(left,right) can be executable.
	Hence, we introduce a channel of type bool on line 18 and send true on it as soon as
	left = MergeSort(data[:mid]) is executed (line 21).

The <-done operation blocks the code on line 25 before the statement Merge(left,right) so that
	it does not proceed until our goroutine has finished. After the goroutine has finished and
	we receive true on the done channel, the code proceeds forward to Merge(left,right) statement
	on line 29.
*/
func mergeSortUsingChannel(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	//Write here channel & go-routine
	done := make(chan bool)
	var left []int
	go func() {
		left = mergeSortUsingChannel(data[:mid])
		done <- true
	}()
	right := mergeSortUsingChannel(data[mid:])
	<-done
	return doMerge(left, right)
}

func doMerge(left, right []int) []int {
	merges := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merges, right...)
		} else if len(right) == 0 {
			return append(merges, left...)
		} else if left[0] < right[0] {
			merges = append(merges, left[0])
			left = left[1:]
		} else {
			merges = append(merges, right[0])
			right = right[1:]
		}
	}
	return merges
}
