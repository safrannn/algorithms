package p_merge_sort

import (
    // "fmt"
	// "sync"
	// "time"
)

func PMerge(T []int64, leftIndex1 int64, rightIndex1 int64, leftIndex2 int64, rightIndex2 int64, A []int64, leftIndex3 int64){
	n1, n2 := rightIndex1 - leftIndex1 + 1, rightIndex2 - leftIndex2 + 1
	if n1 < n2{
		leftIndex1, rightIndex1, leftIndex2, rightIndex2, n1, n2 = leftIndex2,rightIndex2, leftIndex1, rightIndex1, n2, n1
	}

	if n1 == 0{
		return 
	}else{
		mid1 := (leftIndex1 + rightIndex1) / 2
		mid2 := BinarySearch(T[mid1],leftIndex2, rightIndex2,T)
		mid3 := leftIndex3 + (mid1 - leftIndex1) + (mid2 - leftIndex2)
		A[mid3] = T[mid1]

		// var wgm sync.WaitGroup
		// wgm.Add(2)
		// go func() {
		// 	PMerge(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3)
		// 	wgm.Done()
		// }()
		// go func() {
		// 	PMerge(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1)
		// 	wgm.Done()
		// }()
		// wgm.Wait()
		c1 := make(chan struct{})
		c2 := make(chan struct{})
		go func() {
			PMerge(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3)
			c1 <- struct{}{}
		}()
		go func() {
			PMerge(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1)
			c2 <- struct{}{}
		}()
		<-c1
		<-c2

	}
}


func PMergeSort(A []int64, leftIndex int64, rightIndex int64, B []int64, s int64){
	n := rightIndex - leftIndex + 1
	if n == 1{
		B[s] = A[leftIndex]
	}else{
		T := make([]int64, n+1)
		mid1 := (leftIndex + rightIndex) / 2
		mid2 := mid1 - leftIndex + 1

		c1 := make(chan struct{})
		c2 := make(chan struct{})
		// var wg sync.WaitGroup
		// wg.Add(2)
		go func() {
			PMergeSort(A, leftIndex, mid1, T, 1)
			// wg.Done()
			c1 <- struct{}{}
		}()
		go func() {
			PMergeSort(A, mid1 + 1, rightIndex, T, mid2 + 1)
			// wg.Done()
			c2 <- struct{}{}
		}()
		<-c1
		<-c2
		// wg.Wait()
		PMerge(T, 1, mid2, mid2 + 1, n, B, s)
	}
}