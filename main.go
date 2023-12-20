package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/panjf2000/ants/v2"
)

// sendRequest sends an HTTP GET request to the specified URL.
func sendRequest(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP Request Error: %s\n", err)
	}
}

// fibonacci calculates the n-th Fibonacci number.
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// isDifferenceSignificant checks if the time difference between two durations is significant.
func isDifferenceSignificant(current, previous time.Duration, threshold float64) bool {
	if previous == 0 {
		return true
	}
	diff := previous - current
	percentageDiff := (float64(diff) / float64(previous)) * 100
	return percentageDiff >= threshold
}

// testPoolSize tests the pool size for a specific task type.
func testPoolSize(taskType string, threshold float64, totalTasks int, steps int, taskFunc func()) {
	var previousDuration time.Duration

	for poolSize := 2; ; poolSize += steps {
		var wg sync.WaitGroup
		p, _ := ants.NewPool(poolSize)
		defer p.Release()

		start := time.Now()
		for i := 0; i < totalTasks; i++ {
			wg.Add(1)
			_ = p.Submit(func() {
				defer wg.Done()
				taskFunc()
			})
		}
		wg.Wait()
		duration := time.Since(start)

		fmt.Printf("[%s] Pool size: %d, Time taken: %s\n", taskType, poolSize, duration)

		if !isDifferenceSignificant(duration, previousDuration, threshold) {
			green := color.New(color.FgGreen).SprintfFunc()
			fmt.Println(green("[%s] Sweet spot found at pool size: %d\n", taskType, poolSize-1))
			break
		}
		previousDuration = duration
	}
}

func main() {
	var (
		fibonacciNumber = flag.Int("fib", 40, "Fibonacci number to calculate")
		url             = flag.String("url", "http://example.com", "Target URL to send HTTP requests")
		cpuThreshold    = flag.Float64("cpu-threshold", 1.0, "Success threshold for CPU-intensive tasks (percentage)")
		cpuTotalTasks   = flag.Int("cpu-task", 50, "Total CPU tasks")
		httpThreshold   = flag.Float64("http-threshold", 1.0, "Success threshold for HTTP tasks (percentage)")
		httpTotalTasks  = flag.Int("http-task", 500, "Total HTTP requests for each pool size")
		cpuSteps        = flag.Int("cpu-steps", 1, "Steps to increase Go routine pool size")
		httpSteps       = flag.Int("http-steps", 5, "Steps to increase Go routine pool size")
	)

	flag.Parse()

	fmt.Println("Starting CPU tests...")
	// Testing with CPU-intensive tasks (Fibonacci calculation)
	testPoolSize("CPU", *cpuThreshold, *cpuTotalTasks, *cpuSteps, func() {
		_ = fibonacci(*fibonacciNumber)
	})

	fmt.Println("Starting HTTP tests...")
	// Testing with I/O-bound tasks (HTTP requests)
	testPoolSize("HTTP", *httpThreshold, *httpTotalTasks, *httpSteps, func() {
		sendRequest(*url)
	})
}
