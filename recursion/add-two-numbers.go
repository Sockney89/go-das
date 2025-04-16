package recursion

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InvokeAddTwoNumbers() {

	testNode := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	testNode1 := &ListNode{Val: 5, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 7}}}}

	res := addTwoNumbers(testNode, testNode1)

	fmt.Println("Result:", res)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := getNumArray(l1)
	num2 := getNumArray(l2)
	sum := getSumArray(num1, num2)
	res := convSumArrayToListNode(&sum)

	return res
}

func getSumArray(num1 []*int, num2 []*int) []int {

	var longer []*int
	var shorter []*int

	if len(num1) > len(num2) {
		longer = num1
		shorter = num2
	} else {
		longer = num2
		shorter = num1
	}

	var sumArray []int

	var carry int = 0

	for i := range len(longer) {
		if i >= len(shorter) {
			sum := *longer[i] + carry
			digit := sum % 10
			sumArray = append(sumArray, digit)
			carry = sum / 10

		} else {
			sum := (*longer[i] + *shorter[i]) + carry
			digit := sum % 10
			sumArray = append(sumArray, digit)
			carry = sum / 10
		}
		if carry > 0 && i == len(longer)-1 {
			sumArray = append(sumArray, carry)
		}
	}
	return sumArray
}

func convSumArrayToListNode(numArray *[]int) *ListNode {
	var head *ListNode
	var tail *ListNode

	tail = &ListNode{Val: (*numArray)[0]}

	if len(*numArray) == 1 {
		return tail
	}

	remainder := (*numArray)[1:]
	head = convSumArrayToListNode(&remainder)
	return &ListNode{Val: tail.Val, Next: head}
}
