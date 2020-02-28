package map_reduce

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"time"
)

func Run_sequential(n_file int)float64{
	fmt.Println("n_proc: 0","  n_file: ",n_file)
	fmt.Println("n_file_each: ", n_file)

	start := time.Now()

	// read files and calculate frequency
	count := 0
	for file_path := 1; file_path <= n_file; file_path++{
		file, _ := os.Open("../dat/novel_" + strconv.Itoa(file_path) + ".txt")
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
	fmt.Println("result: ", count)

	return time.Since(start).Seconds()
}