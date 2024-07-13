package internal

import (
	"strconv"
)

func StringSliceToIntSlice(slice []string, capacity int) ([]int, error) {
	nums := make([]int, capacity)
	for i, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}
