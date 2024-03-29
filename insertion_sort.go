package main

import (
	"time"
)

var (
	comparissonsIS = 0
	swapsIS        = 0
	evalsIS        = 0
	totalTimeIS    time.Duration
)

func graphInsertionSort(randList []int, updater chan []int) {
	pairsChannel := make(chan []int)
	go insertionSort(randList, len(randList), pairsChannel)
	for pair := range pairsChannel {
		//fmt.Println(pair)
		updater <- pair
		/*m.Lock()
		//update(pair,*bc)

		swapFloats(&(bsChart.Data[pair[0]]), &(bsChart.Data[pair[1]]))
		fmt.Println(bsChart.Data)
		ui.Render(&bsChart)
		time.Sleep(100 * time.Millisecond)
		m.Unlock()*/
	}
	//fmt.Println(randList)
	close(updater)
}

func insertionSort(arr []int, size int, canales chan []int) {
	startTimeIS := time.Now()
	// condicion del ciclo principal
	for i := 1; i < size; i++ {
		evalsIS++
		key := arr[i]
		j := i - 1
		//si el elemento es mayor al key
		for j >= 0 && arr[j] > key {
			evalsIS++
			comparissonsIS++
			//intercambia elementos
			arr[j+1] = arr[j]
			totalTimeIS = time.Since(startTimeIS)
			//envia a canal
			canales <- []int{j, j + 1}
			swapsIS++
			j--
		}
		//avanza al siguiente elemento
		arr[j+1] = key
	}
	close(canales)
	totalTimeIS = time.Since(startTimeIS)
	//fmt.Println(totalTimeIS)
}
