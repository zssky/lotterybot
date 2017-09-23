package gen

import (
	"math/rand"
	"sort"
	"time"
)

func random(count, total int) []int {
	numbers := make([]int, 0)

	for i := 0; i < total; i++ {
		numbers = append(numbers, i+1)
	}

	for i := range numbers {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(numbers)-i) + i

		num := numbers[index]
		numbers[index] = numbers[i]
		numbers[i] = num
	}

	collection := make([]int, 0)
	for i := range numbers {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if r.Intn(len(numbers)-i)+1 <= count-len(collection) {
			collection = append(collection, numbers[i])
		}
		if len(collection) == count {
			break
		}
	}

	sort.Ints(collection)

	return collection
}

func Red() []int {
	return random(6, 33)
}

func Blue() []int {
	return random(1, 16)
}
