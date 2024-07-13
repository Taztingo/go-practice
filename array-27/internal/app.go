package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	if len(os.Args) != 3 {
		fmt.Printf("Unexpected number of args. Expected %d, but received %d\n", 3, len(os.Args))
		fmt.Println("./app 1,2,3 2")
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

	val, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	length := RemoveElement(nums, val)

	fmt.Printf("Output: length=%d nums=%v\n", length, nums)
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
