package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func createSlice(wg *sync.WaitGroup, chanOfNumbers chan int) {
	defer wg.Done()
	sliceOfNumbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		sliceOfNumbers[i] = rand.Intn(100)
	}
	for _, v := range sliceOfNumbers {
		chanOfNumbers <- v
	}
	close(chanOfNumbers)
}

func sqrtNumbers(wg *sync.WaitGroup, chanOfNumbers chan int, chanSqrtNumbers chan int) {
	defer wg.Done()
	for val := range chanOfNumbers {
		chanSqrtNumbers <- val * val
	}
	close(chanSqrtNumbers)

}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	chanOfNumbers := make(chan int, 10)
	chanSqrtNumbers := make(chan int, 10)
	wg.Add(1)
	go createSlice(&wg, chanOfNumbers)
	wg.Add(1)
	go sqrtNumbers(&wg, chanOfNumbers, chanSqrtNumbers)
	for val := range chanSqrtNumbers {
		fmt.Printf("%d ", val)
	}
}
