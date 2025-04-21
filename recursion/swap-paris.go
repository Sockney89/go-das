package recursion

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
