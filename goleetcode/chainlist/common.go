package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListByArry(arr []int) *ListNode {
	head := &ListNode{
		Next: nil,
	}

	for _, v := range arr {
		head.Next = &ListNode{
			Val:  v,
			Next: head.Next,
		}
	}
	return head
}

//尾插法 和数组顺序一致
func NewListByArryReverse(arr []int) *ListNode {
	head := &ListNode{
		Next: nil,
		Val:  0,
	}
	p := head
	for _, v := range arr {
		p.Next = &ListNode{
			Val:  v,
			Next: p.Next,
		}
		p = p.Next
	}
	return head
}

func (l *ListNode) PrintList() {
	l = l.Next
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}

func (l *ListNode) PrintListWithHead() {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}

// 链表翻转
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	p := head.Next
	if p.Next == nil {
		head.Next = nil
		p.Next = head
		return p
	}
	head.Next = nil
	q := p.Next
	p.Next = head
	var temp *ListNode
	for q != nil {
		temp = q
		q = q.Next
		temp.Next = p
		p = temp
	}
	return p
}

// 递归的方式实现
// 返回以该节点开始的翻转节点  1, 2, 3, 4, 5, nil
func getreverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// head = 4, head.Next = 5
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
