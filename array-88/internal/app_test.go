package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taztingo/go-practice/array-88/internal"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "both arrays are empty",
			nums1:    []int{},
			m:        0,
			nums2:    []int{},
			n:        0,
			expected: []int{},
		},
		{
			name:     "nums1 is empty",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "nums2 is empty",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "both contain 1 element",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{2},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "nums 1 has multiple elements",
			nums1:    []int{1, 3, 0},
			m:        2,
			nums2:    []int{2},
			n:        1,
			expected: []int{},
		},
		{
			name:     "nums 2 has multiple elements",
			nums1:    []int{3, 0, 0},
			m:        1,
			nums2:    []int{1, 2},
			n:        2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "both have multiple elements",
			nums1:    []int{1, 2, 0, 0},
			m:        2,
			nums2:    []int{3, 4},
			n:        2,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "duplicates are handled",
			nums1:    []int{1, 1, 0, 0, 0},
			m:        2,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 1, 1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			internal.Merge(tc.nums1, tc.m, tc.nums2, tc.n)
			assert.Equal(t, tc.expected, tc.nums1, "nums1 must contain the correct merged array")
		})
	}
}
