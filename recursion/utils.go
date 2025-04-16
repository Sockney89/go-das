package recursion

func getNumArray(ln *ListNode) []*int {

	var numArray []*int

	if ln == nil {
		return []*int{}
	}
	for {
		numArray = append(numArray, &ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}

	//slices.Reverse(numArray)

	return numArray

}

func convNumArrayToListNode(numArray *[]int) *ListNode {
	var head *ListNode
	var tail *ListNode

	if numArray == nil || len(*numArray) == 0 {
		return nil
	}
	tail = &ListNode{Val: (*numArray)[0]}

	if len(*numArray) == 1 {
		return tail
	}

	remainder := (*numArray)[1:]
	head = convNumArrayToListNode(&remainder)
	return &ListNode{Val: tail.Val, Next: head}
}
