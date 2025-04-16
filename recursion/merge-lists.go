package recursion

import "slices"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	nums1 := getNumArray(list1)
	nums2 := getNumArray(list2)

	merged := append(nums1, nums2...)
	ints := make([]int, len(merged))
	for i, num := range merged {
		ints[i] = *num
	}
	slices.Sort(ints)
}
