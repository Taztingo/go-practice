# Merge Sorted Array

## Description
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Constraints:
- nums1.length == m + n
- nums2.length == n
- 0 <= m, n <= 200
- 1 <= m + n <= 200
- -109 <= nums1[i], nums2[j] <= 109

## Summary
This project is a complete implementation of leetcode's merge sorted array problem. The application takes in two arrays in non-descending order and combines them into one.

## URL
https://leetcode.com/problems/merge-sorted-array

## Examples

### Example 1

    Input: 1,2,3,0,0,0 2,5,6
    Output: [1,2,2,3,5,6]
    Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
    The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.


### Example 2:

    Input: 1 ,
    Output: [1]
    Explanation: The arrays we are merging are [1] and [].
    The result of the merge is [1].

### Example 3:

    Input: , 1
    Output: [1]
    Explanation: The arrays we are merging are [] and [1].
    The result of the merge is [1].
    Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.