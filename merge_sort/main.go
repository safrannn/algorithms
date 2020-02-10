package main

import (
    "fmt"
    "math/rand"
	"time"
	"./merge_sort"
	"./p_merge_sort"
)

func main() {
	fmt.Println("start===============")
	times = runner(10000)
}

func runner(limit int64) []int64{
	times := make([]int64, 5)
	data := make([]int64, limit)
	for i := 0; i < cap(data); i++{
		data[i] = rand.Int63();
	}
	
	// sequential merge sort
	tempData := make([]int64, limit)
	copy(tempData,data)
	start := time.Now()
	merge_sort.SequentialSort(tempData)
	times[0] = time.Since(start). Microseconds()
	fmt.Println("sequential merge sort sorted, time = ",times[0])

	// multithread merge sort
	copy(tempData,data)
	result := make(chan []int64)
	start = time.Now()
	go merge_sort.MergeSort(tempData, result)
	times[1] = time.Since(start). Microseconds()
	<- result
	close(result)
	fmt.Println("multithread merge sort sorted, time = ",times[1])

	// single thread parallel merge sort(?)
	copy(tempData,data)
	C := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PSequentialMergeSort(tempData, 0,limit-1, C, 0)
	times[2] = time.Since(start). Microseconds()
	fmt.Println("sequential parallel merge sort sorted, time = ",times[2])

	// parallel merge sort
	copy(tempData,data)
	D := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PMergeSort(tempData, 0,limit-1, D, 0)
	times[3] = time.Since(start). Microseconds()
	fmt.Println("multithread parallel merge sort sorted, time = ",times[3])

	// parallel merge sort(?)
	copy(tempData,data)
	E := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PMergeSortv1(tempData, 0,limit-1, E, 0, 0)
	times[4] = time.Since(start). Microseconds()
	fmt.Println("multithread parallel merge sort sorted (v1), time = ",times[4])

	return times
}