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

	fmt.Printf("Output: length=%d nums=%v\n", unique, nums)
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := 1
	pointer := 1
	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		current := nums[i]

		// We are at a new number
		if current != prev {
			counter = 0
		}

		// We can safely copy it over
		if counter < 2 {
			nums[pointer] = current
			pointer += 1
		}
		counter += 1
	}

	return pointer
}
