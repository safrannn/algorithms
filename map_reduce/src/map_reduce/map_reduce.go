package map_reduce

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func inputReader(n_proc int, n_file int, out []chan string) {
	// calculate file number to read for each mapper
	n_file_each := make([]int, n_proc)
	for n := 0; n < n_proc; n++{
		n_file_each[n] += n_file / n_proc
	}
	for n := n_file % n_proc; n > 0; n--{
		n_file_each[n]++
	}
	fmt.Println("n_proc: ",n_proc,"  n_file: ",n_file)
	fmt.Println("n_file_each: ", n_file_each)

	// read files for each mapper
	file_path := 1
	for i := 0; i < len(n_file_each); i++{
		// send first file and last file number through channel to mapper
		file_path_start := strconv.Itoa(file_path)
		file_path_end := strconv.Itoa(file_path + n_file_each[i])
		out[i] <- file_path_start + " " + file_path_end
		file_path += n_file_each[i] + 1
	}
}

func mapper(in chan string, out chan int) {
	file_paths := strings.Split(<-in, " ")
	file_path_start,_ := strconv.Atoi(file_paths[0])
	file_path_end,_ := strconv.Atoi(file_paths[1])

	count := 0
	// read file from path and get frequency of qualified words
	for i := file_path_start; i <= file_path_end; i++{
		file, _ := os.Open("../dat/novel_" + strconv.Itoa(i) + ".txt")
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords) 
		for scanner.Scan(){
			word := scanner.Text()
			if len(word) > 0 && unicode.IsUpper(rune(word[0])){
				count++
			}
		}
		file.Close()
	}
	out <- count
	close(in)
}

func reducer(in []chan int, out chan int) {
	sum := 0
	for _,c := range in {
		sum += <-c
		close(c)
	}
	out <- sum
}

func Run(n_proc int, n_file int)float64{
	runtime.GOMAXPROCS(n_proc)

	// set up text file
	var file_paths []chan string
	for f := 1; f <= n_file; f++ {
   		file_paths = append(file_paths, make(chan string))
	}

	// set up map channels
	var maps []chan int
	for p := 1; p <= n_proc; p++ {    
   		maps = append(maps, make(chan int))
	}

	// set up result channel
	reduce := make(chan int)

	// run map_reduce
	start := time.Now()
	go inputReader(n_proc, n_file, file_paths)

	for i := 0; i < len(maps); i++{
		go mapper(file_paths[i], maps[i])
	}

	go reducer(maps, reduce)

	result := <-reduce
	fmt.Println("result: ", result)
	close(reduce)

	return time.Since(start).Seconds()
}
