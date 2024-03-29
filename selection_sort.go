package main

import (
	"time"
)

var (
	comparissonsSS = 0
	swapsSS        = 0
	evalsSS        = 0
	totalTimeSS    time.Duration
)

func graphSelectionSort(randList []int, updater chan []int) {
	pairsChannel := make(chan []int)
	go selectionsort(randList, pairsChannel)
	for pair := range pairsChannel {
		updater <- pair
	}
	close(updater)

}

func selectionsort(items []int, canales chan []int) {
	startTimeSS := time.Now()
	var n = len(items)
	//condicion del ciclo principal
	for i := 0; i < n; i++ {
		evalsSS++
		var minIdx = i
		//asigna el indice y busca el menor para colocarlo al inicio
		for j := i; j < n; j++ {
			evalsSS++
			comparissonsSS++
			if items[j] < items[minIdx] {
				evalsSS++
				minIdx = j
			}
		}
		//intercambia valores
		items[i], items[minIdx] = items[minIdx], items[i]
		totalTimeSS = time.Since(startTimeSS)
		//envia informacion al canal
		canales <- []int{i, minIdx}
		swapsSS++
	}
	close(canales)
	totalTimeSS = time.Since(startTimeSS)
}
