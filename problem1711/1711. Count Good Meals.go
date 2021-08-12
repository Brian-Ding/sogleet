package problem1711

import "fmt"

func Judge() {
	deliciousness := []int{1, 3, 5, 7, 9}
	result := countPairs(deliciousness)

	fmt.Println(result)
}

func countPairs(deliciousness []int) int {
	count := 0

	for i := 0; i < len(deliciousness); i++ {
		for j := i + 1; j < len(deliciousness); j++ {
			if isGoodMeal(deliciousness[i] + deliciousness[j]) {
				count++
			}
		}
	}

	return count
}

func isGoodMeal(meal int) bool {
	for i := 0; i <= 20; i++ {
		if meal>>i == 1 {
			return true
		}
	}

	return false
}
