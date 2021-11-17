package main

import (
	"time"
)

var (
	comparissonsHS = 0
	swapsHS        = 0
	totalTimeHS    time.Duration
)

func graphHeapSort(randList []int, updater chan []int) {
	pairsChannel := make(chan []int)
	go heapSort(randList, len(randList), pairsChannel)
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

func buildMaxHeap(arr []int, n int, canales chan []int) {
	for i := 1; i < n; i++ {
		//if child is bigger than parent
		comparissonsHS++
		if arr[i] > arr[(i-1)/2] {
			j := i
			// swap child and parent until
			// parent is smaller
			for arr[j] > arr[(j-1)/2] {
				swap(&arr[j], &arr[(j-1)/2])
				swapsHS++
				canales <- []int{j, (j - 1) / 2}
				j = (j - 1) / 2
			}
		}
	}
}

func heapSort(arr []int, n int, canales chan []int) {
	startTimeHS := time.Now()
	buildMaxHeap(arr, n, canales)
	//fmt.Println(arr)
	for i := n - 1; i > 0; i-- {
		// swap value of first indexed
		// with last indexed
		swap(&arr[0], &arr[i])
		canales <- []int{0, i}
		swapsHS++
		// maintaining heap property
		// after each swapping
		j := 0
		index := 0

		for {
			index = 2*j + 1
			if index > i {
				break
			}
			// if left child is smaller than
			// right child point index variable
			// to right child
			if arr[index] < arr[index+1] && index < (i-1) {
				index++
			}
			// if parent is smaller than child
			// then swapping parent with child
			// having higher value
			comparissonsHS++
			if arr[j] < arr[index] && index < i {
				swap(&arr[j], &arr[index])
				swapsHS++
				canales <- []int{j, index}
			}
			j = index
		}
	}
	close(canales)
	totalTimeHS = time.Since(startTimeHS)
	//fmt.Println(totalTimeHS)
}

/*
// Driver Code to test above
int main()
{
int arr[] = {10, 20, 15, 17, 9, 21};
int n = sizeof(arr) / sizeof(arr[0]);

printf("Given array: ");
for (int i = 0; i < n; i++)
printf("%d ", arr[i]);

printf("\n\n");

heapSort(arr, n);

// print array after sorting
printf("Sorted array: ");
for (int i = 0; i < n; i++)
printf("%d ", arr[i]);

return 0;
}*/
