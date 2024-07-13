package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taztingo/go-practice/array-27/internal"
)

func TestStringSliceToIntSlice(t *testing.T) {
	tests := []struct {
		name     string
		str      []string
		cap      int
		expected []int
		err      string
	}{
		{
			name:     "empty array",
			str:      []string{},
			cap:      0,
			expected: []int{},
		},
		{
			name:     "single digit",
			str:      []string{"1"},
			cap:      1,
			expected: []int{1},
		},
		{
			name:     "multiple digit",
			str:      []string{"12"},
			cap:      2,
			expected: []int{12, 0},
		},
		{
			name: "no digit",
			str:  []string{"abc"},
			err:  "strconv.Atoi: parsing \"abc\": invalid syntax",
		},
		{
			name: "mix digit",
			str:  []string{"1a"},
			err:  "strconv.Atoi: parsing \"1a\": invalid syntax",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice, err := internal.StringSliceToIntSlice(tc.str, tc.cap)
			if err == nil {
				assert.Equal(t, tc.expected, slice, "must have correctly converted elements in slice")
			} else {
				assert.EqualErrorf(t, err, tc.err, "must output correct error for invalid input")
			}
		})
	}
}
