package internal

func Run() {
	/*
		if len(os.Args) != 3 {
			fmt.Printf("Unexpected number of args. Expected %d, but received %d\n", 3, len(os.Args))
			fmt.Println("./app 1,2,3 7,8,9")
			os.Exit(1)
		}

		nums1Str := strings.Split(os.Args[1], ",")
		nums2Str := strings.Split(os.Args[2], ",")

		if nums1Str[0] == "" {
			nums1Str = []string{}
		}
		if nums2Str[0] == "" {
			nums2Str = []string{}
		}

		fmt.Printf("num1str=%v\n", nums1Str)

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
	*/
}

func RemoveElement(nums []int, val int) int {
	freeSlot := 0

	for i := range nums {
		if nums[i] != val {
			nums[freeSlot] = nums[i]
			freeSlot++
		}
	}

	return freeSlot
}
