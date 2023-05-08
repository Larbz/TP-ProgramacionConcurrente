package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func quicksort(arr []int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	if len(arr) <= 1 {
		<-sem // liberar el semáforo
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

	// obtener el semáforo
	sem <- struct{}{}
	go quicksort(arr[:right], wg, sem)

	// obtener el semáforo
	sem <- struct{}{}
	go quicksort(arr[right+1:], wg, sem)
}

func main() {
	i := 0
	var prom time.Duration

	for ; i < 10; i++ {
		start := time.Now()
		rand.Seed(time.Now().UnixNano())
		n := 1000000
		dataset := rand.Perm(n)
		//fmt.Println("Lista desordenada:", dataset)
		var wg sync.WaitGroup
		sem := make(chan struct{}, n) // crear un canal con capacidad de 10

		wg.Add(1)

		// obtener el semáforo
		sem <- struct{}{}
		go quicksort(dataset, &wg, sem)

		wg.Wait()
		//fmt.Println("Lista ordenada:", dataset)
		elapsed := time.Since(start)
		prom += elapsed
	}

	fmt.Printf("El tiempo de ejecución fue de %v\n", prom/10)
}
