package main

// func reverseList(head *ListNode) *ListNode {
// 	pre := head
// 	p := head.
// 	for p.Next != nil {
// 		fmt.Println(p.Val)
// 		p = p.Next
// 	}
// 	fmt.pl
// 	return p
// }

func main() {
	l := NewListByArry([]int{1, 2, 3, 4, 5})
	re := reverseList(l)
	re.PrintListWithHead()
}
