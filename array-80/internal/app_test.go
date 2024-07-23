package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taztingo/go-practice/array-80/internal"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		expectedSlice  []int
		expectedOutput int
	}{
		{
			name:           "nums is empty",
			nums:           []int{},
			expectedSlice:  []int{},
			expectedOutput: 0,
		},
		{
			name:           "nums 1 element",
			nums:           []int{1},
			expectedSlice:  []int{1},
			expectedOutput: 1,
		},
		{
			name:           "nums multi element no duplicate",
			nums:           []int{1, 2},
			expectedSlice:  []int{1, 2},
			expectedOutput: 2,
		},
		{
			name:           "nums multi element duplicate",
			nums:           []int{1, 2, 2, 3},
			expectedSlice:  []int{1, 2, 2, 3},
			expectedOutput: 4,
		},
		{
			name:           "nums multi element duplicate beginning",
			nums:           []int{1, 1, 2, 3},
			expectedSlice:  []int{1, 1, 2, 3},
			expectedOutput: 4,
		},
		{
			name:           "nums multi element triplet",
			nums:           []int{1, 2, 2, 2, 3},
			expectedSlice:  []int{1, 2, 2, 3, 3},
			expectedOutput: 4,
		},
		{
			name:           "nums multi element triplet beginning",
			nums:           []int{1, 1, 1, 2, 3},
			expectedSlice:  []int{1, 1, 2, 3, 3},
			expectedOutput: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			remaining := internal.RemoveDuplicates(tc.nums)
			assert.Equal(t, tc.expectedSlice, tc.nums, "nums must contain the correct updated array")
			assert.Equal(t, tc.expectedOutput, remaining, "output must be number of elements that are not val")
		})
	}
}
