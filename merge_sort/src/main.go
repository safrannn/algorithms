package main

import (
    "math/rand"
	"time"
	"sync"
	"strconv"
	"os"
    "encoding/csv"
	"./merge_sort"
	"./p_merge_sort"
)

func main() {
	experiment := 5
	results := [][]int64{}
	
	var limit int64
	limit_number := []int64{}
	// for limit = 100; limit <= 10000000; limit *= 10{
	for limit = 100; limit <= 100; limit *= 10{
		results = append(results,runnerMultiple(experiment,limit))
		limit_number = append(limit_number,limit)
	}

	printResult(results,limit_number)
}

func runnerMultiple(experiment int, limit int64) []int64{
	result := make([]int64,6)
	for k := 0; k < experiment; k++{
		times := runnerSingle(limit)
		for i := 0; i < 6; i++{
			result[i] += times[i]
		}
	}
	for i := 0; i < 6; i++{
		result[i] /= int64(experiment)
	} 
	return result
}

func runnerSingle(limit int64) []int64{
	times := make([]int64, 6)
	data := make([]int64, limit)
	for i := 0; i < cap(data); i++{
		data[i] = rand.Int63();
	}
	tempData := make([]int64, limit)

	// sequential merge sort
	copy(tempData,data)
	start := time.Now()
	merge_sort.SequentialSort(tempData)
	times[0] = time.Since(start). Microseconds()

	// multithread merge sort
	copy(tempData,data)
	result := make(chan []int64)
	start = time.Now()
	go merge_sort.MergeSort(tempData, result)
	times[1] = time.Since(start). Microseconds()
	<- result
	close(result)
	// fmt.Println("multithread merge sort sorted, time = ",times[1])

	// single thread parallel merge sort(?)
	copy(tempData,data)
	C := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PSequentialMergeSort(tempData, 0,limit-1, C, 0)
	times[2] = time.Since(start). Microseconds()
	// fmt.Println("sequential parallel merge sort sorted, time = ",times[2])

	// multithread parallel merge sort
	copy(tempData,data)
	D := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PMergeSort(tempData, 0,limit-1, D, 0)
	times[3] = time.Since(start). Microseconds()
	// fmt.Println("multithread parallel merge sort sorted, time = ",times[3])

	// multithread parallel merge sort v1
	copy(tempData,data)
	E := make([]int64,limit,limit)
	start = time.Now()
	p_merge_sort.PMergeSortv1(tempData, 0,limit-1, E, 0, 0)
	times[4] = time.Since(start). Microseconds()
	// fmt.Println("multithread parallel merge sort sorted (v1), time = ",times[4])

	// // parallel merge sort v2
	copy(tempData,data)
	F := make([]int64,limit,limit)
	var wg sync.WaitGroup
	wg.Add(30)
	c := make(chan bool)
	start = time.Now()
	go p_merge_sort.PMergeSortv2(&wg,c,tempData, 0,limit-1, F, 0, 0)
	wg.Wait()
	times[5] = time.Since(start). Microseconds()
	// fmt.Println("multithread parallel merge sort sorted (v2), time = ",times[5])
	
	return times
}

func printResult(results [][]int64, limit_number []int64){
	file,_ := os.Create("../out/results.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
    defer writer.Flush()

	head := []string{"population","merge sort(µs)","mt merge sort(µs)","p merge sort(µs)","mt p merge sort(µs)","mt p merge sort v1(µs)","mt p merge sort v2(µs)"}
	writer.Write(head)
	for i, result := range results {
		resultString := make([]string,len(result)+1)
		resultString[0] = strconv.FormatInt(limit_number[i],10)
		for i,v := range result{
			resultString[i+1] = strconv.FormatInt(v,10)
		}
        writer.Write(resultString)
    }
}