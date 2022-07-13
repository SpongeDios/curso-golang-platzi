package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f Function
	cache map[int]FunctionResult
	lock sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err error
}

func NewMemory(f Function) *Memory {
	return &Memory{f: f, cache: make(map[int]FunctionResult)}
}

func (memory *Memory) Get(key int) (interface{}, error) {
	memory.lock.Lock()
	results, exists := memory.cache[key]
	memory.lock.Unlock()

	if !exists {
		memory.lock.Lock()
		results.value, results.err = memory.f(key)
		memory.cache[key] = results
		memory.lock.Unlock()
	}
	return results.value, results.err
}

func GetFibonacci(n int) (interface{}, error){
	return Fibonacci(n), nil
}

func main() {
	cache := NewMemory(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)
	}
	wg.Wait()
}