package main

import (
	"fmt"
)

func main() {
	l := NewListByArryReverse([]int{1, 2, 4})
	l1 := NewListByArryReverse([]int{1, 3, 4})
	ans := mergeTwoLists(l.Next, l1.Next)
	ans.Next.PrintList()
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

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	p := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			p.Next = list2
			return head
		}
		if list2 == nil {
			p.Next = list1
			return head
		}
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	return head.Next
}
