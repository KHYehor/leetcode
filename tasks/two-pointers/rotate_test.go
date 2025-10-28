package two_pointers

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

// Helper function to convert linked list to slice for comparison
func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	result := []int{}
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		k        int
		expected []int
	}{
		{"empty list", []int{}, 5, []int{}},
		{"single element", []int{1}, 10, []int{1}},
		{"no rotation k=0", []int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{"rotate by 1", []int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{"rotate by 2", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"rotate by length", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{"rotate more than length", []int{1, 2, 3, 4, 5}, 7, []int{4, 5, 1, 2, 3}},
		{"two elements rotate 1", []int{1, 2}, 1, []int{2, 1}},
		{"two elements rotate 3", []int{1, 2}, 3, []int{2, 1}},
		{"example from problem", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"another example", []int{0, 1, 2}, 4, []int{2, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.vals)
			result := rotateRight(head, tt.k)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("got %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}
