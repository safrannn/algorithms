package main

import (
    "fmt"
    "math/rand"
	"time"
	"./merge_sort"
	"./p_merge_sort"
)

func main() {
	runner(1000)

}

func runner(limit int64){
	times := make([]int64, 4) // ms
	data := make([]int64, limit)
	for i := 0; i < cap(data); i++{
		data[i] =  rand.Int63();
	}
	// fmt.Println(data)
	
	// sequential merge sort
	tempData := make([]int64, limit)
	copy(tempData,data)
	start := time.Now()
	returnData := merge_sort.SequentialSort(tempData)
	times[0] = time.Since(start). Microseconds()
	fmt.Println("// sequential merge sort sorted",len(returnData))

	// multithread merge sort
	copy(tempData,data)
	result := make(chan []int64)
	start = time.Now()
	go merge_sort.MergeSort(tempData, result)
	times[1] = time.Since(start). Microseconds()
	temp := <- result
	close(result)
	fmt.Println("// multithread merge sort sorted", len(temp))

	// single thread parallel merge sort(?)
	copy(tempData,data)
	C := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PSequentialSort(tempData, 0,limit-1, C, 0)
	times[2] = time.Since(start). Microseconds()
	fmt.Println("// sequential parallel merge sort sorted")

	// parallel merge sort(?)
	copy(tempData,data)
	D := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PMergeSort(tempData, 0,limit-1, D, 0)
	times[3] = time.Since(start). Microseconds()
	fmt.Println("// multithread parallel merge sort sorted")

	fmt.Println(times[0])
	fmt.Println(times[1])
	fmt.Println(times[2])
	fmt.Println(times[3])
}