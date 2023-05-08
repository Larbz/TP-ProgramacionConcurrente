package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(L []int, R []int, arr []int) {
	num := len(arr)
	i, j, k := 0, 0, 0
	for i < num/2 && j < num-num/2 {
		if L[i] < R[j] {
			arr[k] = L[i]
			i++
			k++
		} else {
			arr[k] = R[j]
			j++
			k++
		}
	}
	for i < num/2 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < num-num/2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func mergeSort(arr []int) {
	num := len(arr)
	if num > 1 {
		var L, R []int
		L = append(L, arr[0:num/2]...)
		R = append(R, arr[num/2:num]...)
		mergeSort(L)
		mergeSort(R)
		merge(L, R, arr)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	numeros := rand.Perm(1000000)
	fmt.Println(numeros)
	start := time.Now()
	mergeSort(numeros)
	duration := time.Since(start)
	fmt.Println(numeros)
	fmt.Println(duration.Milliseconds())
}
