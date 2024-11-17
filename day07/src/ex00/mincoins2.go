// Пакет ex00 предоставляет функции для вычисления минимального количества монет,
// необходимых для представления заданной суммы с использованием определенного набора номиналов.
package ex00

import "sort"

// minCoins2 вычисляет минимальное количество монет, необходимое для составления заданной суммы.
// В отличие от оригинальной функции minCoins, minCoins2 обрабатывает случаи, когда
// массив номиналов может содержать дубликаты или быть несортированным.
// Если массив номиналов пустой, функция возвращает пустой срез.
//
// Данная реализация обеспечивает корректный результат благодаря:
// 1. Удалению дубликатов из входного массива номиналов.
// 2. Сортировке массива для упрощения обработки.
//
// Пример использования:
//     coins := []int{1, 3, 4, 7, 13, 15}
//     result := minCoins2(23, coins)
//     fmt.Println(result) // Вывод: [15, 4, 4]
//
// Чтобы сгенерировать документацию для этой функции, используйте следующую команду:
//     godoc -http=:6060
// Откройте ваш веб-браузер и перейдите по адресу http://localhost:6060/pkg/ex00/ для просмотра документации.

func MinCoins2(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}
	uniqueCoins := removeDuplicatesAndSort(coins)
	res := make([]int, 0)
	for i := len(uniqueCoins) - 1; i >= 0; i-- {
		for val >= uniqueCoins[i] {
			val -= uniqueCoins[i]
			res = append(res, uniqueCoins[i])
		}
	}
	return res
}

// removeDuplicatesAndSort принимает срез номиналов монет, удаляет дубликаты и сортирует его в порядке возрастания.
// Это полезно для подготовки данных к функциям, которые требуют уникальных и отсортированных номиналов.
//
// Пример использования:
//
//	coins := []int{5, 1, 3, 3, 7, 1}
//	result := removeDuplicatesAndSort(coins)
//	fmt.Println(result) // Вывод: [1, 3, 5, 7]
//
// Реализация использует карту для удаления дубликатов и встроенную функцию сортировки для упорядочивания данных.
func removeDuplicatesAndSort(coins []int) []int {
	set := make(map[int]struct{})
	for _, coin := range coins {
		set[coin] = struct{}{}
	}
	uniqueCoins := make([]int, 0, len(set))
	for coin := range set {
		uniqueCoins = append(uniqueCoins, coin)
	}
	sort.Ints(uniqueCoins)
	return uniqueCoins
}
