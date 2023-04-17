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
		if result := containsDuplicate(st.nums); result != st.result {
			t.Errorf("wanted %v (%t), got %t", st.nums,
				st.result, result)
		}
	}
}