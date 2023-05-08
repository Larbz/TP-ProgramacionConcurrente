package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
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



func mergeSort(arr []int, wg *sync.WaitGroup) {
	num := len(arr)
	defer wg.Done()
	if num > 1 {
		var L, R []int
		L = append(L, arr[0:num/2]...)
		R = append(R, arr[num/2:num]...)
		var wgLeft, wgRight sync.WaitGroup
		wgLeft.Add(1)
		wgRight.Add(1)
		go mergeSort(L, &wgLeft)
		go mergeSort(R, &wgRight)
		wgLeft.Wait()
		wgRight.Wait()
		merge(L, R, arr)
	}

}

func analyseTime(times []int){
	sort.Ints(times);
	times=times[5:len(times)-5]
	suma:=0
	for i := 0; i < len(times); i++ {
		suma+=times[i]
	}
	fmt.Println("El tiempo promedio de ejecucion para un arreglo de 1 millon de datos es igual a ",suma/1000," segundos")
}

func main() {
	var times []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 20; i++ {
		numeros := rand.Perm(1000000)
		var wg sync.WaitGroup
		wg.Add(1)
		start := time.Now()
		go mergeSort(numeros, &wg)
		wg.Wait()
		duration := time.Since(start)
		times=append(times,int(duration.Milliseconds()))
	}
	analyseTime(times)
	
}
