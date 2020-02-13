package merge_sort

// import "sync"

func Merge(leftData []int64, rightData []int64) (result []int64) {
    result = make([]int64, len(leftData) + len(rightData))
    leftIndex, rightIndex := 0, 0

    for i:=0;i<cap(result);i++ {
        switch {
            case leftIndex >= len(leftData):
                result[i] = rightData[rightIndex]
                rightIndex++
            case rightIndex >= len(rightData):
                result[i] = leftData[leftIndex]
                leftIndex++
            case leftData[leftIndex] < rightData[rightIndex]:
                result[i] = leftData[leftIndex]
                leftIndex++
            default:
                result[i] = rightData[rightIndex]
                rightIndex++
        }
    }
    return
}

func SequentialSort(data []int64) []int64{
	if len(data) < 2 {
		return data
	}
	mid := (len(data)) / 2
	return Merge(SequentialSort(data[:mid]), SequentialSort(data[mid:]))
}

func MergeSort(data []int64, result chan []int64) {
    if len(data) < 2 {
        result <- data
        return
    }

    leftChan := make(chan []int64)
    rightChan := make(chan []int64)
    middle := len(data)/2

    go MergeSort(data[:middle], leftChan)
    go MergeSort(data[middle:], rightChan)

    leftData := <-leftChan
    rightData := <-rightChan

    close(leftChan)
    close(rightChan)
	result <- Merge(leftData, rightData)
    return
}

