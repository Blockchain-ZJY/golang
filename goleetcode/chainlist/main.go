package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

func main() {

	// l1 := NewListByArryReverse([]int{1, 3, 4})
	// ans := mergeTwoLists(l.Next, l1.Next)
	// ans.Next.PrintList()

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for headA != nil && headB != nil {
		a = a.Next
		b = b.Next
	}
	count := 0
	if a == nil {
		for b != nil {
			b = b.Next
			count++
		}
		for i := 0; i < count; i++ {
			headB = headB.Next
		}
		for headB != headA {
			headA = headA.Next
			headB = headB.Next
		}
		return headB
	} else {
		for a != nil {
			a = a.Next
			count++
		}
		for i := 0; i < count; i++ {
			headA = headA.Next
		}
		for headB != headA {
			headA = headA.Next
			headB = headB.Next
		}
		return headB
	}

}

func isPalindrome(head *ListNode) bool {
	sk := arraystack.New()
	len := 0 //5/2= 2
	// 4/2 2
	a := head

	for a != nil {
		len++
		a = a.Next
	}
	if len%2 == 1 {
		count := 0
		for head != nil {
			if count < len/2 {
				sk.Push(head.Val)
				head = head.Next
			} else if count == len/2 {
				count++
				head = head.Next
				continue
			} else {
				if v, ok := sk.Peek(); ok && v != head.Val {
					return false
				} else {
					sk.Pop()
					head = head.Next
				}
			}
			count++
		}
	}
	if len%2 == 0 {
		count := 0
		for head != nil {
			if count < len/2 {
				sk.Push(head.Val)
				head = head.Next

			} else {
				if v, ok := sk.Peek(); ok && v != head.Val {
					return false
				} else {
					sk.Pop()
					head = head.Next
				}
			}
			count++
		}

	}
	if sk.Empty() {
		return true
	}
	return false
}

// 32. 最长有效括号
func longestValidParentheses(s string) int {
	stack := arraystack.New()
	stack.Push(-1)
	ans := 0
	for i := range s {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if !stack.Empty() {
				top, _ := stack.Peek()
				ans = max(ans, i-top.(int))
			} else {
				stack.Push(i)
			}

		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseList1(head *ListNode) *ListNode {
	var root *ListNode
	for head != nil {
		head.Next, root, head = root, head, head.Next
	}
	return root
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}
	return dummy.Next
}

// 删除链表的倒数第 N 个结点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	len := 0
	s := head
	for s != nil {
		len++
		s = s.Next
	}
	if n == len {
		return head.Next
	}

	fast, slow := head, head
	for i := 0; i < n && fast.Next != nil; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if n == 1 {
		slow.Next = nil
		return head
	}
	slow.Next = slow.Next.Next
	return head
}

// 归并排序(链表版本)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	r := getmid(head)
	head = sortList(head)
	r = sortList(r)
	return merge(head, r)
}

func getmid(head *ListNode) *ListNode {
	f, s := head.Next, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	t := s.Next
	s.Next = nil
	return t
}

func merge(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}
	if l == nil {
		cur.Next = r
	} else {
		cur.Next = l
	}
	return dummy.Next
}
