# goroutine-benchmark

Running concurrent Go routines is an effective way to maximize the performance of multi-core processors. If you have undefined amount of tasks, it’s better to limit number of running go routines with a pool mechanism. However, determining this limit is not straightforward. It usually does not equal to the number of processors and can vary between CPU-related and network-related tasks.

This tool allows you to benchmark your current system to identify the optimal number of go routines. It tests both CPU-related and network-related tasks.

# Installation

## From Binary

You can download the pre-built binaries from the [releases](https://github.com/utkusen/goroutine-benchmark/releases/latest) page and run. For example:

`unzip goroutine-benchmark-linux-amd64.zip`

`./goroutine-benchmark --help`

## From Source

1) Install Go on your system

2) Run: `go install github.com/utkusen/goroutine-benchmark@latest`

# Usage

You can run it with default values

`./goroutine-benchmark`

It will produce a result like:

```
Starting CPU tests...
[CPU] Pool size: 2, Time taken: 30.906672859s
[CPU] Pool size: 3, Time taken: 21.174364486s
[CPU] Pool size: 4, Time taken: 14.614268344s
[CPU] Pool size: 5, Time taken: 12.463523168s
[CPU] Pool size: 6, Time taken: 10.923756879s
[CPU] Pool size: 7, Time taken: 9.343834108s
[CPU] Pool size: 8, Time taken: 8.111060596s
[CPU] Pool size: 9, Time taken: 7.547088552s
[CPU] Pool size: 10, Time taken: 6.669233083s
[CPU] Pool size: 11, Time taken: 7.255197785s
[CPU] Sweet spot found at pool size: 10

Starting HTTP tests...
[HTTP] Pool size: 2, Time taken: 46.302199477s
[HTTP] Pool size: 7, Time taken: 13.294319729s
[HTTP] Pool size: 12, Time taken: 7.771916167s
[HTTP] Pool size: 17, Time taken: 5.541366266s
[HTTP] Pool size: 22, Time taken: 4.247880245s
[HTTP] Pool size: 27, Time taken: 3.525955192s
[HTTP] Pool size: 32, Time taken: 2.968827238s
[HTTP] Pool size: 37, Time taken: 2.589811276s
[HTTP] Pool size: 42, Time taken: 2.252026786s
[HTTP] Pool size: 47, Time taken: 2.037507483s
[HTTP] Pool size: 52, Time taken: 1.851666206s
[HTTP] Pool size: 57, Time taken: 1.676372247s
[HTTP] Pool size: 62, Time taken: 1.637649876s
[HTTP] Pool size: 67, Time taken: 1.481247566s
[HTTP] Pool size: 72, Time taken: 1.321769664s
[HTTP] Pool size: 77, Time taken: 1.300674719s
[HTTP] Sweet spot found at pool size: 76
```

You can also fine-tune different parameters:

```
  -cpu-steps int
    	Steps to increase Go routine pool size (default 1)
  -cpu-task int
    	Total CPU tasks (default 50)
  -cpu-threshold float
    	Success threshold for CPU-intensive tasks (percentage) (default 1)
  -fib int
    	Fibonacci number to calculate (default 40)
  -http-steps int
    	Steps to increase Go routine pool size (default 5)
  -http-task int
    	Total HTTP requests for each pool size (default 500)
  -http-threshold float
    	Success threshold for HTTP tasks (percentage) (default 1)
  -url string
    	Target URL to send HTTP requests (default "http://example.com")
```
