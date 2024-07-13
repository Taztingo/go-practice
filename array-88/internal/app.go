package internal

func Run() {
	// Get some input
	//Merge()

	// Output the results of Merge
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
