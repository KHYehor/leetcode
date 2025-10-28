package two_pointers

// https://leetcode.com/problems/rotate-list/?envType=problem-list-v2&envId=two-pointers

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// Empty list or List with one element
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Copy values to the array
	arr := make([]int, 0)
	cur := head
	for {
		arr = append(arr, cur.Val)
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}

	rot := k % len(arr)
	// No rotation needed
	if rot == 0 {
		return head
	}

	// Iterate over the list
	cur = head
	for i := 0; i < len(arr); i++ {
		// Apply elements from the end
		if i < rot {
			id := len(arr) - rot + i
			cur.Val = arr[id]
		} else {
			id := i - rot
			cur.Val = arr[id]
		}
		cur = cur.Next
	}

	return head
}
