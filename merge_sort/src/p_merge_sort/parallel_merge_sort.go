package p_merge_sort

import "sync"

const MAXDEPTH = 4

func BinarySearch(x int64,left int64, right int64, T []int64)int64{
	low, high := left, left
	if right + 1 > left{
		high = right + 1
	}

	for low < high{
		mid := (low + high) / 2
		if x <= T[mid]{
			high = mid
		}else{
			low = mid + 1
		}
	}

	return high
}

func PSequentialMerge(T []int64, leftIndex1 int64, rightIndex1 int64, leftIndex2 int64, rightIndex2 int64, A []int64, leftIndex3 int64){
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
		PSequentialMerge(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3)
		PSequentialMerge(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1)
	}
}
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

		c := make(chan bool)
		go func(c chan bool){
			PMerge(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3)
			c <- true
		}(c)
		PMerge(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1)
		<-c
	}
}

func PMergev1(T []int64, leftIndex1 int64, rightIndex1 int64, leftIndex2 int64, rightIndex2 int64, A []int64, leftIndex3 int64, depth int){
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

		if depth < MAXDEPTH{
			c := make(chan bool)
			go func(c chan bool){
				PMergev1(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3, depth+1)
				c <- true
			}(c)
			PMergev1(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1, depth+1)
			<-c
			close(c)
		}else{
			PSequentialMerge(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3)
			PSequentialMerge(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1)
		}
	}
}

func PMergev2(wg *sync.WaitGroup, c chan bool,T []int64, leftIndex1 int64, rightIndex1 int64, leftIndex2 int64, rightIndex2 int64, A []int64, leftIndex3 int64, depth int){
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

		if depth < MAXDEPTH{
			go PMergev2(wg,c,T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3, depth+1)
			wg.Done()
			go PMergev2(wg,c,T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1, depth+1)
			wg.Done()
		}else{
			PSequentialMerge(T,leftIndex1, mid1 - 1, leftIndex2, mid2 - 1, A, leftIndex3)
			PSequentialMerge(T,mid1 + 1, rightIndex1, mid2, rightIndex2, A, mid3 + 1)
		}
	}
}

func PSequentialMergeSort(A []int64, leftIndex int64, rightIndex int64, B []int64, s int64){
	n := rightIndex - leftIndex + 1
	if n == 1{
		B[s] = A[leftIndex]
	}else{
		T := make([]int64, n+1)
		mid1 := (leftIndex + rightIndex) / 2
		mid2 := mid1 - leftIndex + 1
		PSequentialMergeSort(A, leftIndex, mid1, T, 1)
		PSequentialMergeSort(A, mid1 + 1, rightIndex, T, mid2 + 1)
		PSequentialMerge(T, 1, mid2, mid2 + 1, n, B, s)
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
	
		c := make(chan bool)
		go func(c chan bool){
			PMergeSort(A, leftIndex, mid1, T, 1)
			c <- true
		}(c)
		PMergeSort(A,mid1 + 1, rightIndex, T, mid2 + 1)
		<-c
		close(c)
		PMerge(T, 1, mid2, mid2 + 1, n, B, s)
	}
}

func PMergeSortv1(A []int64, leftIndex int64, rightIndex int64, B []int64, s int64, depth int){
	n := rightIndex - leftIndex + 1
	if n == 1{
		B[s] = A[leftIndex]
	}else{
		T := make([]int64, n+1)
		mid1 := (leftIndex + rightIndex) / 2
		mid2 := mid1 - leftIndex + 1
	
		if depth < MAXDEPTH{
			c := make(chan bool)
			go func(c chan bool){
				PMergeSortv1(A, leftIndex, mid1, T, 1, depth+1)
				c <- true
			}(c)
			PMergeSortv1(A,mid1 + 1, rightIndex, T, mid2 + 1, depth+1)
			<-c
			close(c)
			PMergev1(T, 1, mid2, mid2 + 1, n, B, s, depth + 1)
		}else{
			PSequentialMergeSort(A, leftIndex, mid1, T, 1)
			PSequentialMergeSort(A, mid1 + 1, rightIndex, T, mid2 + 1)
			PSequentialMerge(T, 1, mid2, mid2 + 1, n, B, s)
		}
	}
}

func PMergeSortv2(wg *sync.WaitGroup, c chan bool,A []int64, leftIndex int64, rightIndex int64, B []int64, s int64, depth int){
	n := rightIndex - leftIndex + 1
	if n == 1{
		B[s] = A[leftIndex]
	}else{
		T := make([]int64, n+1)
		mid1 := (leftIndex + rightIndex) / 2
		mid2 := mid1 - leftIndex + 1

		if depth < MAXDEPTH{
			go PMergeSortv2(wg,c, A, leftIndex, mid1, T, 1, depth+1)
			wg.Done()
			go PMergeSortv2(wg, c, A, mid1 + 1, rightIndex, T, mid2 + 1, depth+1)
			wg.Done()
			go PMergev1(T, 1, mid2, mid2 + 1, n, B, s, depth + 1)
		}else{
			PSequentialMergeSort(A, leftIndex, mid1, T, 1)
			PSequentialMergeSort(A, mid1 + 1, rightIndex, T, mid2 + 1)
			PSequentialMerge(T, 1, mid2, mid2 + 1, n, B, s)
		}		
	}
}