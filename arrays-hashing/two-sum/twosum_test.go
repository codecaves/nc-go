package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	subtests := []struct {
		nums []int
		target int
		result []int
	}{
		{
			nums: []int{2,7,11,15},
			target: 9,
			result: []int{0,1}, 
		},
		{
			nums: []int{3,2,4},
			target: 6,
			result: []int{1,2},
		},
		{
			nums: []int{3,3},
			target: 6,
			result: []int{0,1},
		},
	}

	for _, st := range subtests {
		if res := twoSum(st.nums, st.target); !reflect.DeepEqual(res, st.result) {
			t.Errorf("For %v and %d wanted (%v), got %v", st.nums, st.target, res, st.result)
		}
	}
}