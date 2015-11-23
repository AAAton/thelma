package main

import "math/rand"

//QuickSort is function
func QuickSort(slice []string) []string {
	length := len(slice)

	if length <= 1 {
		sliceCopy := make([]string, length)
		copy(sliceCopy, slice)
		return sliceCopy
	}

	m := slice[rand.Intn(length)]

	less := make([]string, 0, length)
	middle := make([]string, 0, length)
	more := make([]string, 0, length)

	for _, item := range slice {
		switch {
		case len(item) > len(m):
			less = append(less, item)
		case len(item) == len(m):
			middle = append(middle, item)
		case len(item) < len(m):
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
