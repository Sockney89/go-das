package recursion

//Function for swapping adjacent nodes in a linked list
//Link to Leetcode question: https://leetcode.com/problems/swap-nodes-in-pairs
func swapPairs(head *ListNode) *ListNode {

	current := head
	for {
		if current == nil || current.Next == nil {
			break
		}
		current.Val, current.Next.Val = current.Next.Val, current.Val

		current = current.Next.Next
	}

	return head
}
