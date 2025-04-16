package recursion

func getNumArray(ln *ListNode) []*int {

	var numArray []*int

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
