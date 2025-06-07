package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Константы для размера слайса и количества частей
const (
	SIZE   = 100_000_000 // Размер генерируемого слайса
	CHUNKS = 8           // Количество частей для параллельной обработки
)

// generateRandomElements генерирует случайные положительные числа и записывает их в слайс.
// Если размер слайса равен 0 или отрицательный, возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	// Создаем слайс нужного размера и заполняем его случайными числами
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) // Генерируем числа от 0 до 999999
	}
	return data
}

// maximum находит максимальное число в слайсе.
// Если слайс пустой, возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Начинаем с первого элемента и сравниваем с остальными
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks находит максимальное число в слайсе, разделяя его на 8 частей
// и ища максимум в каждой части параллельно.
func maxChunks(data []int) int {
	// Обработка краевых случаев
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	// Вычисляем размер каждой части и создаем слайс для хранения максимумов
	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	// Запускаем горутины для поиска максимума в каждой части
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data) // Для последнего чанка берем все оставшиеся элементы
		}

		// Запускаем горутину для поиска максимума в текущей части
		go func(chunkIndex, start, end int) {
			defer wg.Done()
			chunkMax := data[start]
			for j := start + 1; j < end; j++ {
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}
			maxValues[chunkIndex] = chunkMax
		}(i, start, end)
	}

	// Ждем завершения всех горутин
	wg.Wait()
	// Находим максимум среди максимумов частей
	return maximum(maxValues)
}

func main() {
	// Инициализируем генератор случайных чисел текущим временем
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	// Измеряем время поиска максимума в один поток
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	// Измеряем время поиска максимума в 8 потоков
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
