package tasks

import (
	"reflect"
	"testing"
)

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var out []int
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}

func TestGcd(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{18, 6, 6},
		{6, 10, 2},
		{10, 3, 1},
		{7, 7, 7},
		{1, 1000, 1},
		{1000, 1000, 1000},
		{17, 13, 1},
	}

	for _, tt := range tests {
		if got := gcd(tt.a, tt.b); got != tt.expected {
			t.Errorf("gcd(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
		}
	}
}

func TestInsertGreatestCommonDivisors(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"example 1", []int{18, 6, 10, 3}, []int{18, 6, 6, 2, 10, 1, 3}},
		{"example 2 single node", []int{7}, []int{7}},
		{"two nodes", []int{4, 6}, []int{4, 2, 6}},
		{"equal values", []int{5, 5, 5}, []int{5, 5, 5, 5, 5}},
		{"coprime values", []int{2, 3, 5, 7}, []int{2, 1, 3, 1, 5, 1, 7}},
		{"ones", []int{1, 1}, []int{1, 1, 1}},
		{"max values", []int{1000, 1000}, []int{1000, 1000, 1000}},
		{"divisible chain", []int{12, 24, 48}, []int{12, 12, 24, 24, 48}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			result := insertGreatestCommonDivisors(head)

			// The returned node must be the head of the modified list.
			if result != head {
				t.Errorf("returned node is not the original head (got Val=%d)", result.Val)
			}

			if got := listToSlice(result); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
