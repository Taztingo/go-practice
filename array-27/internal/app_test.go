package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taztingo/go-practice/array-27/internal"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name           string
		nums           []int
		val            int
		expectedSlice  []int
		expectedOutput int
	}{
		{
			name:           "nums is empty",
			nums:           []int{},
			val:            1,
			expectedSlice:  []int{},
			expectedOutput: 0,
		},
		{
			name:           "val doesn't exist",
			nums:           []int{1},
			val:            2,
			expectedSlice:  []int{1},
			expectedOutput: 1,
		},
		{
			name:           "val removes one element",
			nums:           []int{1},
			val:            1,
			expectedSlice:  []int{1},
			expectedOutput: 0,
		},
		{
			name:           "val removes multiple elements",
			nums:           []int{1, 1},
			val:            1,
			expectedSlice:  []int{1, 1},
			expectedOutput: 0,
		},
		{
			name:           "val ignores other elements",
			nums:           []int{1, 2, 1},
			val:            1,
			expectedSlice:  []int{2, 2, 1},
			expectedOutput: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			remaining := internal.RemoveElement(tc.nums, tc.val)
			assert.Equal(t, tc.expectedSlice, tc.nums, "nums must contain the correct updated array")
			assert.Equal(t, tc.expectedOutput, remaining, "output must be number of elements that are not val")
		})
	}
}
