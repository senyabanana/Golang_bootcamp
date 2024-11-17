package calculation

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Numbers
	}{
		{
			name:  "valid input",
			input: "10\n20\n30\ngenerate\n",
			want:  Numbers{10, 20, 30},
		},
		{
			name:  "input with empty lines",
			input: "10\n\n20\n   \n30\ngenerate\n",
			want:  Numbers{10, 20, 30},
		},
		{
			name:  "mixed valid and invalid input",
			input: "10\n-100001\n30\n002\ngenerate\n",
			want:  Numbers{10, 30},
		},
		{
			name:  "non-integer input",
			input: "10\nabc\n30\ngenerate\n",
			want:  Numbers{10, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем временный файл с тестовыми данными
			tmpFile, err := os.CreateTemp("", "test_input")
			if err != nil {
				t.Fatalf("Не удалось создать временный файл: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			// Записываем тестовые данные во временный файл
			if _, err := tmpFile.WriteString(tt.input); err != nil {
				t.Fatalf("Не удалось записать данные во временный файл: %v", err)
			}
			if _, err := tmpFile.Seek(0, 0); err != nil {
				t.Fatalf("Не удалось сбросить указатель файла: %v", err)
			}

			// Сохраняем текущий os.Stdin и восстанавливаем его в конце теста
			oldStdin := os.Stdin
			defer func() {
				os.Stdin = oldStdin
			}()

			// Перенаправляем os.Stdin на временный файл
			os.Stdin = tmpFile

			nums := ParseInput()
			assert.Equal(t, tt.want, nums)
		})
	}
}

func TestCalculateMean(t *testing.T) {
	tests := []struct {
		name    string
		numbers Numbers
		want    float64
	}{
		{
			name:    "empty slice",
			numbers: Numbers{},
			want:    0.0,
		},
		{
			name:    "single element",
			numbers: Numbers{10},
			want:    10.0,
		},
		{
			name:    "positive numbers",
			numbers: Numbers{1, 2, 3, 4, 5},
			want:    3.0,
		},
		{
			name:    "negative numbers",
			numbers: Numbers{-1, -2, -3, -4, -5},
			want:    -3.0,
		},
		{
			name:    "mixed numbers",
			numbers: Numbers{1, -1, 2, -2, 3, -3},
			want:    0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.numbers.CalculateMean()
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCalculateMedian(t *testing.T) {
	tests := []struct {
		name    string
		numbers Numbers
		want    float64
	}{
		{
			name:    "empty slice",
			numbers: Numbers{},
			want:    0.0,
		},
		{
			name:    "single element",
			numbers: Numbers{10},
			want:    10.0,
		},
		{
			name:    "odd number of elements",
			numbers: Numbers{1, 3, 2},
			want:    2.0,
		},
		{
			name:    "even number of elements",
			numbers: Numbers{1, 2, 3, 4},
			want:    2.5,
		},
		{
			name:    "mixed numbers",
			numbers: Numbers{1, -1, 2, -2, 3, -3},
			want:    0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.numbers.CalculateMedian()
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCalculateMode(t *testing.T) {
	tests := []struct {
		name    string
		numbers Numbers
		want    int64
	}{
		{
			name:    "single mode",
			numbers: Numbers{1, 2, 2, 3, 4},
			want:    2,
		},
		{
			name:    "multiple modes, smallest selected",
			numbers: Numbers{1, 1, 2, 2, 3},
			want:    1,
		},
		{
			name:    "all elements same",
			numbers: Numbers{5, 5, 5, 5, 5},
			want:    5,
		},
		{
			name:    "no repeated elements",
			numbers: Numbers{1, 2, 3, 4, 5},
			want:    1,
		},
		{
			name:    "negative numbers with mode",
			numbers: Numbers{-1, -1, -2, -2, -3, -3, -3},
			want:    -3,
		},
		{
			name:    "negative and positive numbers",
			numbers: Numbers{-1, -1, 1, 1, 2},
			want:    -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.numbers.CalculateMode())
		})
	}
}

func TestCalculateSD(t *testing.T) {
	tests := []struct {
		name    string
		numbers Numbers
		want    float64
	}{
		{
			name:    "empty slice",
			numbers: Numbers{},
			want:    0.0,
		},
		{
			name:    "single element",
			numbers: Numbers{10},
			want:    0.0,
		},
		{
			name:    "two elements",
			numbers: Numbers{1, 1},
			want:    0.0,
		},
		{
			name:    "three elements",
			numbers: Numbers{1, 2, 3},
			want:    0.82,
		},
		{
			name:    "mixed numbers",
			numbers: Numbers{1, -1, 2, -2, 3, -3},
			want:    2.16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.numbers.CalculateSD()
			assert.Equal(t, tt.want, result)
		})
	}
}
