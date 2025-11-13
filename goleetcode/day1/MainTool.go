package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrintList() {
	readnode := new(ListNode)
	readnode = l
	for readnode != nil {
		fmt.Print(readnode.Val, " ")
		readnode = readnode.Next
	}
	fmt.Println()
}

func (l *ListNode) Getlen() int {
	conter := l
	lens := 0
	for conter != nil {
		lens++
		conter = conter.Next
	}
	return lens
}

func (l *ListNode) CreateList(nums []int) *ListNode {
	head := new(ListNode)
	l1 := head
	for i := 0; i < len(nums); i++ {
		l1.Val = nums[i]
		if i == len(nums)-1 {
			l1.Next = nil
		} else {
			l1.Next = new(ListNode)
		}
		l1 = l1.Next
	}
	return head
}
