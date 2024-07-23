package internal

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	if len(os.Args) != 2 {
		fmt.Printf("Unexpected number of args. Expected %d, but received %d\n", 2, len(os.Args))
		fmt.Println("./app 1,2,3")
		os.Exit(1)
	}

	numsStr := strings.Split(os.Args[1], ",")

	if numsStr[0] == "" {
		numsStr = []string{}
	}

	nums, err := StringSliceToIntSlice(numsStr, len(numsStr))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	unique := RemoveDuplicates(nums)

	fmt.Printf("Output: unique=%d nums=%v\n", unique, nums)
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slot := 1
	for i := slot; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[slot] = nums[i]
			slot += 1
		}
	}

	return slot
}
