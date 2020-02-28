### Introduction
Purpose of this lab is to use map reduce model to find frequency of title-cased words from files. 
Execution time of different size of input data and number of process are tested in this experiment.

### Description
In this experiment, map reduce is composed of a map procedure that reads data and count the title word 
frequency for several chunks of the data, and a reduce method that combines the result. 
The reason for not having a shuffler in between map and reduce and not sending text data from input reader 
to map through channel is that reading the file is quite time consuming comparing the processing of the data.
Number of reducer is 1.

### Experiment
Sequential version of data processing is compared with results of 2, 4, 6 and 8 process on a machine with 8-core CPU. 
Each version is tested with benchmarking of 8, 12, 16, 20 and 24 files, each around 100MB.
Golang 1.13 is used for this experiment. Process limited is set using runtime.GOMAXPROCS(). 
Number of the goroutine for mapper is the same as the process limit for each execution. 
Execution time is reported in seconds.
 
 

# Run
```sh
// src/main.go
$ go run main.go
```
