package twosum

func twoSum(nums []int, target int) []int {
	solution := []int{}

	Exit:
		for i, num_i := range nums {
			for j, num_j := range nums[i+1:] {
				if (num_i + num_j == target) {
					solution = append(solution, i)
					solution = append(solution, i + j + 1)
					break Exit
				}
			}
		}

	return solution
}