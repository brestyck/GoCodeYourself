package main

import (
	"fmt"
	"runtime"
	"time"
)

func merge(left, right []int) (res []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	res = append(res, left...)
	res = append(res, right...)

	return
}

func mergeSort(items []int) {
	runtime.GOMAXPROCS(4)
	if len(items) < 2 {
		return
	}
	middle := len(items) / 2
	leftHalf := items[:middle]
	rightHalf := items[middle:]
	channel := make(chan struct{})
	go func() { mergeSort(leftHalf); channel <- struct{}{} }()
	go func() { mergeSort(rightHalf); channel <- struct{}{} }()
	<-channel
	<-channel
	copy(items, merge(leftHalf, rightHalf))
}

func main() {
	a := []int{}
	for {
		var x int
		if n, _ := fmt.Scan(&x); n == 0 {
			break
		}
		a = append(a, x)
	}
	t1 := time.Now()
	mergeSort(a[:])
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Milliseconds(), "ms")
	// fmt.Println(a)
}
