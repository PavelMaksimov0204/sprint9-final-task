// Пишите тесты в этом файле

package main

import (
	"testing"
)

// TestGenerateRandomElements проверяет функцию генерации случайных чисел
// на различных входных данных и краевых случаях
func TestGenerateRandomElements(t *testing.T) {
	// Тестовые случаи для проверки различных размеров слайса
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "Нулевой размер",
			size:     0,
			expected: 0,
		},
		{
			name:     "Размер 1",
			size:     1,
			expected: 1,
		},
		{
			name:     "Размер 1000",
			size:     1000,
			expected: 1000,
		},
		{
			name:     "Отрицательный размер",
			size:     -1,
			expected: 0,
		},
	}

	// Проверяем каждый тестовый случай
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Получаем результат генерации
			result := generateRandomElements(tt.size)
			// Проверяем размер полученного слайса
			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = %d элементов; ожидалось %d",
					tt.size, len(result), tt.expected)
			}
			// Проверяем, что все сгенерированные числа положительные
			for i, v := range result {
				if v < 0 {
					t.Errorf("generateRandomElements(%d) содержит отрицательное число %d на позиции %d",
						tt.size, v, i)
				}
			}
		})
	}
}

// TestMaximum проверяет функцию поиска максимума
// на различных входных данных и краевых случаях
func TestMaximum(t *testing.T) {
	// Тестовые случаи для проверки различных слайсов
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Несколько элементов",
			input:    []int{1, 5, 3, 8, 2},
			expected: 8,
		},
		{
			name:     "Одинаковые элементы",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "Отрицательные числа",
			input:    []int{-1, -5, -3, -8},
			expected: -1,
		},
	}

	// Проверяем каждый тестовый случай
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d; ожидалось %d",
					tt.input, result, tt.expected)
			}
		})
	}
}

// TestMaxChunks проверяет функцию поиска максимума в чанках
// на различных входных данных и краевых случаях
func TestMaxChunks(t *testing.T) {
	// Тестовые случаи для проверки различных размеров слайса
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Меньше 8 элементов",
			input:    []int{1, 5, 3, 8, 2},
			expected: 8,
		},
		{
			name:     "Ровно 8 элементов",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: 8,
		},
		{
			name:     "Больше 8 элементов",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 10,
		},
	}

	// Проверяем каждый тестовый случай
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d; ожидалось %d",
					tt.input, result, tt.expected)
			}
		})
	}
}
