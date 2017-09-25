package util

import (
	"math/rand"
	"sort"
	"time"
)

func Remove(l []int, i int) []int {
	if i == len(l) - 1 {
		return l[:i]
	}
	return append(l[:i], l[i+1:]...)
}

func RandomSort(start, end int) []int {
	numbers := make([]int, 0)
	for i := 0; i < end-start+1; i++ {
		numbers = append(numbers, i+start)
	}

	for i := range numbers {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(numbers)-i) + i

		num := numbers[index]
		numbers[index] = numbers[i]
		numbers[i] = num
	}

	return numbers
}

func AverageSelector(numbers []int, count int) []int {
	replica := make([]int, len(numbers))
	copy(replica, numbers)
	collection := make([]int, 0)

	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		index := r.Intn(len(replica) - len(collection))
		collection = append(collection, replica[index])
		replica = Remove(replica, index)

		if len(collection) == count {
			break
		}
	}

	sort.Ints(collection)

	return collection
}

func Split(numbers []int, num int) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	for i := range numbers {
		if numbers[i] < num {
			left = append(left, numbers[i])
		} else {
			right = append(right, numbers[i])
		}
	}

	return left, right
}
