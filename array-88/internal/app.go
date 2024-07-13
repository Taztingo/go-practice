package internal

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	if len(os.Args) != 3 {
		fmt.Printf("Unexpected number of args. Expected %d, but received %d\n", 3, len(os.Args))
		fmt.Println("./app 1,2,3 7,8,9")
		os.Exit(1)
	}

	nums1Str := strings.Split(os.Args[1], ",")
	nums2Str := strings.Split(os.Args[2], ",")

	nums1, err := StringSliceToIntSlice(nums1Str, len(nums1Str)+len(nums2Str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nums2, err := StringSliceToIntSlice(nums2Str, len(nums2Str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Merge(nums1, len(nums1Str), nums2, len(nums2Str))

	fmt.Printf("Output: %v\n", nums1)
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	// Create temporary destination and pointers
	nums3 := make([]int, len(nums1))
	i := 0
	j := 0

	// Select the smaller from the 2 arrays and add to nums3
	for i, j = 0, 0; i < m && j < n; {
		if nums1[i] <= nums2[j] {
			nums3[i+j] = nums1[i]
			i += 1
		} else {
			nums3[i+j] = nums2[j]
			j += 1
		}
	}

	// Copy the remainders
	for ; i < m; i++ {
		nums3[i+j] = nums1[i]
	}
	for ; j < n; j++ {
		nums3[i+j] = nums2[j]
	}

	// Copy into nums1
	copy(nums1, nums3)
}
