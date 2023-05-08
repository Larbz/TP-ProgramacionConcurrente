package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	i := 0
	var prom time.Duration

	for ; i < 10; i++ {
		start := time.Now()
		rand.Seed(time.Now().UnixNano())
		n := 1000000
		dataset := rand.Perm(n)

		quickSortStart(dataset)
		//fmt.Println(dataset)
		elapsed := time.Since(start)
		prom += elapsed
	}

	fmt.Printf("El tiempo de ejecución fue de %v\n", prom/10)
}
