package main

import(
	"encoding/csv"
	"./map_reduce"
	"os"
	"runtime"
	"strconv"
)

func printResults(results [][]float64){
	file, _ := os.Create("../out/time.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
    defer writer.Flush()

	for _, result := range results {
		resultString := make([]string, len(result))
		for i, v := range result{
			resultString[i] = strconv.FormatFloat(v, 'f', 2, 64)
		}
        writer.Write(resultString)
    }
}

func run_all(){
	max_proc := runtime.NumCPU()
	results := [][]float64{}

	// run benchmarks to test result
	for f := 8; f <= 24; f += 4{
		result := []float64{}

		// sequential version running time
		run_time := map_reduce.Run_sequential(f)
		result = append(result, run_time)

		// map reduce versions running times
		for p := 2; p <= max_proc; p += 2{
			run_time = map_reduce.Run(p, f)
			result = append(result, run_time)
		}

		results = append(results, result)
	}

	// output result
	printResults(results)
}

func main() {
	run_all()
}