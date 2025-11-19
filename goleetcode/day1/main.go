package main

// 两数之和
func twoSum(nums []int, target int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
				return ans
			}
		}
	}
	return ans
}

// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var longerone *ListNode
	var shorterone *ListNode
	var ans *ListNode
	if l1.Getlen() > l2.Getlen() {
		longerone = l1
		ans = l1
		shorterone = l2
	} else {
		longerone = l2
		shorterone = l1
		ans = l2
	}
	nextAddition := 0 //进位初始为0
	for longerone != nil {
		temp := (longerone.Val + shorterone.Val + nextAddition) % 10
		nextAddition = (longerone.Val + shorterone.Val + nextAddition) / 10
		longerone.Val = temp

		if nextAddition != 0 && longerone.Next == nil {
			longerone.Next = new(ListNode)
			longerone.Next.Val = nextAddition
			break
		}

		if shorterone.Next == nil && longerone.Next != nil {
			for longerone.Next != nil {
				temp = (longerone.Next.Val + nextAddition) % 10
				nextAddition = (longerone.Next.Val + nextAddition) / 10
				longerone.Next.Val = temp
				longerone = longerone.Next
				if nextAddition != 0 && longerone.Next == nil {
					longerone.Next = new(ListNode)
					longerone.Next.Val = nextAddition
					break
				}
			}
			break
		}
		longerone = longerone.Next
		shorterone = shorterone.Next
	}
	// 长的移动
	return ans
}

func main() {
	list1 := []int{5}
	list2 := []int{5}
	l1 := new(ListNode)
	l2 := new(ListNode)
	head1 := l1.CreateList(list1)
	head2 := l2.CreateList(list2)

	ans := addTwoNumbers(head1, head2)
	ans.PrintList()
}
