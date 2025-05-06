package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// Генерация случайных чисел
func generateRandomElements(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	return data
}

// Поиск максимума в одном потоке
func maximum(data []int) int {
	maxVal := data[0]
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// Поиск максимума в нескольких потоках
func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	maxChan := make(chan int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(d []int) {
			defer wg.Done()
			localMax := d[0]
			for _, val := range d {
				if val > localMax {
					localMax = val
				}
			}
			maxChan <- localMax
		}(data[start:end])
	}

	wg.Wait()
	close(maxChan)

	finalMax := <-maxChan
	for val := range maxChan {
		if val > finalMax {
			finalMax = val
		}
	}
	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	start := time.Now()
	data := generateRandomElements(SIZE)
	fmt.Printf("Генерация завершена за %d ms\n", time.Since(start).Milliseconds())

	fmt.Println("Ищем максимальное значение в один поток")
	start = time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxParallel := maxChunks(data)
	elapsedParallel := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxParallel, elapsedParallel)
}
