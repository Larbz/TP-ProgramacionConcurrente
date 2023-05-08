package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func quicksort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	left := 1
	right := len(arr) - 1
	for left <= right {
		for left <= len(arr)-1 && arr[left] < pivot {
			left++
		}
		for right >= 1 && arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	arr[0], arr[right] = arr[right], arr[0]
	wg.Add(2)
	go quicksort(arr[:right], wg)
	go quicksort(arr[right+1:], wg)
}

func main() {
	i := 0
	var prom time.Duration

	for ; i < 10; i++ {
		start := time.Now()
		var wg sync.WaitGroup
		rand.Seed(time.Now().UnixNano())
		dataset := rand.Perm(100000)
		//fmt.Println("Lista desordenada:", dataset)
		wg.Add(1)
		go quicksort(dataset, &wg)
		wg.Wait()
		//fmt.Println("Lista ordenada:", dataset)
		elapsed := time.Since(start)
		prom += elapsed
	}

	fmt.Printf("El tiempo de ejecución fue de %v\n", prom/10)
}
