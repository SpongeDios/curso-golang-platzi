package main

import (
	"fmt"
)

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40, 45}
	numberWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < numberWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d, started fib with %d\n", id, job)
		fibonacci := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fibonacci)
		results <- fibonacci
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}
