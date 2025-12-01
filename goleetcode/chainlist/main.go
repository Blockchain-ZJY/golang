package main

import (
	"fmt"
)

func main() {
	l := NewListByArryReverse([]int{1, 2, 3, 4, 5, 6, 7, 8})
	ans := removeNthFromEnd(l.Next, 7)
	ans.PrintList()
}

// 19. 删除链表的倒数第 N 个结点

// 找到被删除的前一个节点的
// 如果是头结点直接返回NEXT
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s, f := head, head
	len := 1
	for i := 0; i < n && f.Next != nil; i++ {
		len++
		f = f.Next
	}
	for f.Next != nil {
		len++
		f = f.Next
		s = s.Next
	}
	if len == n {
		//删除头元素
		return head.Next
	}
	if len == n-1 {
		ANS := head.Next.Next
		head.Next = ANS
		return head
	}
	p := s
	p = s.Next.Next
	s.Next = p
	fmt.Println(s.Val, f.Val, len)
	return head
}
