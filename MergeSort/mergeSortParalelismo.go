package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge(L chan int, R chan int, ch chan int) {
	defer close(ch)
	valL, okL := <-L
	valR, okR := <-R
	for okL && okR {
		if valL < valR {
			ch <- valL
			valL, okL = <-L
		} else {
			ch <- valR
			valR, okR = <-R
		}
	}
	for okL {
		ch <- valL
		valL, okL = <-L
	}
	for okR {
		ch <- valR
		valR, okR = <-R
	}
}

func mergeSort(arr []int, ch chan int) {
	num := len(arr)
	if num > 1 {
		var L, R []int
		L = append(L, arr[0:num/2]...)
		R = append(R, arr[num/2:num]...)
		ch1 := make(chan int)
		ch2 := make(chan int)
		go mergeSort(L, ch1)
		go mergeSort(R, ch2)
		merge(ch1, ch2, ch)
	} else {
		ch <- arr[0]
		defer close(ch)
		return
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	numeros := rand.Perm(1000000)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Arreglo original: ", numeros)
	start := time.Now()
	go func() {
		wg.Done()
		mergeSort(numeros, ch)
	}()
	wg.Wait()
	var s []int
	for v := range ch {
		s = append(s, v)
	}
	duration := time.Since(start)
	fmt.Println(s)
	fmt.Println("Duration since start: ", duration.Milliseconds())
}
