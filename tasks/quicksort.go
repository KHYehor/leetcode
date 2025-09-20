package tasks

import (
	"sync"

	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Integer | constraints.Float](arr []T) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort[T](arr[:left])
	QuickSort[T](arr[left+1:])
}

func QuickSortMT[T constraints.Integer | constraints.Float](arr []T) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		QuickSort[T](arr[:left])
		wg.Done()
	}()

	go func() {
		QuickSort[T](arr[left+1:])
		wg.Done()
	}()

	wg.Wait()
}
