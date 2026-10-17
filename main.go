package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	generatedNumbers := make([]int, 0, size)
	for i := 0; i < size; i++ {
		generatedNumbers = append(generatedNumbers, rand.Intn(SIZE))
	}

	return generatedNumbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("cannot find maximum of an empty slice")
	}
	maxNum := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > maxNum {
			maxNum = data[i]
		}
	}
	return maxNum, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	maxOfChunks := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		leftIdx := i * chunkSize
		rightIdx := leftIdx + chunkSize
		if i == CHUNKS-1 {
			rightIdx = len(data)
		}

		go func(chunkIdx int, data []int) {
			defer wg.Done()

			maxOfChunk, err := maximum(data)
			if err != nil {
				log.Printf("failed to find max in chunk: %v", err)
				return
			}

			maxOfChunks[chunkIdx] = maxOfChunk
		}(i, data[leftIdx:rightIdx])
	}
	wg.Wait()

	return maximum(maxOfChunks)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max, err := maximum(data)
	if err != nil {
		log.Printf("failed to find maximum in whole slice: %v", err)
	}
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max, err = maxChunks(data)
	if err != nil {
		log.Printf("failed to find max with chunks: %v", err)
	}
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
