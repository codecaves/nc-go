package duplicate

import "testing"

func TestMultiContainsDuplicate(t *testing.T) {
	subtests := []struct {
		nums []int
		result bool
	}{
		{
			nums: []int {1,2,3,1},
			result: true,
		},
		{
			nums: []int {1,2,3,4},
			result: false,
		},
		{
			nums: []int {1,1,1,3,3,4,3,2,4,2},
			result: true,
		},
		{
			nums: []int {},
			result: false,
		},
		{
			nums: []int {0},
			result: false,
		},
	}

	for _, st := range subtests {
		if res := containsDuplicate(st.nums); res != st.result {
			t.Errorf("For %v wanted (%t), got %t", st.nums,
				res, st.result)
		}
	}
}