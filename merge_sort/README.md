# Merge Sort
Two algorithms are tested using golang.
First one is the merge sort, a sequential version and a multithreaded version is tested.
Second one is the p merge sort, a sequential one, a multithreaded one and two optimized version is tested.
First optimized version uses depth to limit the number of goroutine. If function is executed when depth is larger than the maximum depth, sequential merge and sequential sort will be called.
Second optimized version is based upon the first optimized version and creates a pool of channels beforehand to reduce the time of channel creation during execution.

# Run
```sh
// src/main.go
$ go run main.go
```
