package calculation

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Numbers []int64

// ParseInput принимает последовательность чисел, а так же валидирует ввод.
func ParseInput() Numbers {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите числа, разделенные новыми строками (для завершения введите 'generate'):")

	var nums Numbers

	for scanner.Scan() {
		line := scanner.Text()
		if line == "generate" {
			break
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		if len(line) > 1 && line[0] == '0' {
			fmt.Printf("Ошибка ввода: ведущие нули %s\n", line)
			continue
		}
		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("Ошибка ввода %s\n", err)
			continue
		}
		if number > 100000 || number < -100000 {
			fmt.Printf("Число %d выходит за границы [-100000, 100000]\n", number)
			continue
		}
		nums = append(nums, number)
	}
	return nums
}

// CalculateMean вычисляет среднее значение из переданных чисел.
func (n Numbers) CalculateMean() float64 {
	if len(n) == 0 {
		return 0.0
	}
	sum := int64(0)
	for _, num := range n {
		sum += num
	}
	mean := float64(sum) / float64(len(n))
	return math.Round(mean*100) / 100
}

// CalculateMedian находит среднее число отсортированной последовательности, если ее размер нечетный,
// или среднее значение между двумя средними числами в отсортированной последовательности.
func (n Numbers) CalculateMedian() float64 {
	lenNums := len(n)
	if lenNums == 0 {
		return 0.0
	}
	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})
	if lenNums%2 == 1 {
		return float64(n[lenNums/2])
	} else {
		middleOne := float64(n[lenNums/2-1])
		middleTwo := float64(n[lenNums/2])
		median := (middleOne + middleTwo) / 2
		return math.Round(median*100) / 100
	}
}

// CalculateMode вычисляет значение, которое наиболее часто встречается в последовательности
// В случае нескольких значений с одинаковой максимальной частотой - возвращается наименьшее из них.
func (n Numbers) CalculateMode() int64 {
	frequency := make(map[int64]int)
	for _, num := range n {
		frequency[num]++
	}
	maxFreq := 0
	mode := int64(0)
	for num, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
			mode = num
		} else if freq == maxFreq && num < mode {
			mode = num
		}
	}
	return mode
}

// CalculateSD находит среднее отклонение относительно среднего значения последовательности
func (n Numbers) CalculateSD() float64 {
	if len(n) == 0 {
		return 0.0
	}
	mean := n.CalculateMean()
	varianceSum := float64(0)
	for _, num := range n {
		varianceSum += math.Pow(float64(num)-mean, 2)
	}
	variance := varianceSum / float64(len(n))
	sd := math.Sqrt(variance)
	return math.Round(sd*100) / 100
}
