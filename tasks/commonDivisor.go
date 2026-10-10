package tasks

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var cur = head
	for cur.Next != nil {
		next := cur.Next
		// Create a new node
		var newNode = &ListNode{
			Val:  gcd(cur.Val, next.Val),
			Next: next,
		}
		// Reassign head to the newNode
		cur.Next = newNode
		// Move the pointer to the new value
		cur = next
	}

	return head
}
