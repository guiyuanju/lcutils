package list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(vals []int) *ListNode {
	head := &ListNode{0, nil}
	prev := head
	for _, v := range vals {
		prev.Next = &ListNode{v, nil}
		prev = prev.Next
	}
	return head.Next
}

func (l *ListNode) String() string {
	var res []string
	head := l
	res = append(res, "[")
	for head != nil {
		res = append(res, fmt.Sprintf("%d, ", head.Val))
		head = head.Next
	}
	res = append(res, "]")
	return strings.Join(res, "")
}
