package filter

import "testing"

func TestLeach(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	killed := []int{1, 4}
	t.Logf("%v", Leach(numbers, killed))
}