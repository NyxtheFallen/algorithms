//Package algorithms exposes go implementations of popular algorithms. For simplicity and efficiency, they operate only on integers.
package algorithms

//MergeSort performs a merge sort on a slice of integers. This algorithm is extremely efficient as n becomes large.
func MergeSort(slice []int) ([]int, error) {
	length := len(slice)
	if length == 2 && slice[0] > slice[1] {
		ret := []int{slice[1], slice[0]}
		return ret, nil
	} else if length <= 2 {
		ret := make([]int, length)
		copy(ret, slice)
		return ret, nil
	} else {
		halfLength := length / 2
		left := slice[:halfLength]
		right := slice[halfLength:]
		leftChan := make(chan []int, 2)
		rightChan := make(chan []int, 2)

		func(splitSlice chan []int) {
			val, _ := MergeSort(left)
			splitSlice <- val
		}(leftChan)

		func(splitSlice chan []int) {
			val, _ := MergeSort(right)
			splitSlice <- val
		}(rightChan)

		left = <-leftChan
		right = <-rightChan
		return merge(left, right), nil
	}
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left)+len(right))
	leftOverflow := false
	rightOverflow := false
	for i, j, k := 0, 0, 0; k < len(merged); k++ {
		leftOverflow = i >= len(left)
		rightOverflow = j >= len(right)
		if !leftOverflow && !rightOverflow {
			if left[i] < right[j] {
				merged[k] = left[i]
				i++
			} else {
				merged[k] = right[j]
				j++
			}
		} else if leftOverflow {
			merged[k] = right[j]
			j++
		} else {
			merged[k] = left[i]
			i++
		}
	}
	return merged
}
