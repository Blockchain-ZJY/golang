package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {

			return
		}
		inorder(root.Left) // 左
		ans = append(ans, root.Val)
		inorder(root.Right) // 右
	}
	inorder(root)
	return ans
}

func NewTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右子节点
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func main() {

}

// 101. 对称二叉树(这个不行❗ 中序遍历加翻转无法确认一个二叉树)
func isSymmetric(root *TreeNode) bool {
	l := []int{}
	r := []int{}
	var inorder func(root *TreeNode, ans *[]int)
	inorder = func(root *TreeNode, ans *[]int) {
		if root == nil {
			return
		}
		inorder(root.Left, ans) // 左
		*ans = append(*ans, root.Val)
		inorder(root.Right, ans) // 右
	}
	inorder(root.Left, &l)
	inorder(root.Right, &r)
	fmt.Println(l, r)
	slices.Reverse(l)
	return slices.Compare(l, r) == 0
}

// 需要用到后续遍历,左右中 只有收集到树的情况才能在中的时候判断结果

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}

// 543. 二叉树的直径
// 后序遍历找到当前节点左右边深度最大值
func diameterOfBinaryTree(root *TreeNode) int {
	maxv := 0
	getmax(root, &maxv)
	return maxv
}

func getmax(root *TreeNode, maxv *int) {
	if root == nil {
		return
	}
	*maxv = max(*maxv, treedeepth(root.Left)+treedeepth(root.Right))
	getmax(root.Left, maxv)
	getmax(root.Right, maxv)
	return
}

// 意义, 以当前节点为根节点的最大高度
func treedeepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := treedeepth(root.Left)
	r := treedeepth(root.Right)
	return max(l, r) + 1
}

func diameterOfBinaryTreeBetter(root *TreeNode) int {
	maxv := 0
	var depth func(*TreeNode) int

	// 意义, 以当前节点为节点的最长高度
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		maxv = max(maxv, left+right)
		return max(left, right) + 1
	}

	depth(root)
	return maxv
}

// 102. 二叉树的层序遍历
// 队列
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{}
	q = append(q, root)

	for len(q) != 0 {
		layer := []int{}
		len := len(q)
		for i := 0; i < len; i++ {
			head := q[0]
			q = q[1:]
			layer = append(layer, head.Val)
			if head.Left != nil {
				q = append(q, head.Left)
			}
			if head.Right != nil {
				q = append(q, head.Right)
			}
		}
		ans = append(ans, layer)
	}
	return
}

// 98. 验证二叉搜索树
// 这个不行,单独看该节点左右满足大小关系不够,不完全
// 用BST中序遍历是严格有序的特点,需要记录前置节点
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	pre = nil
	var helper func(root *TreeNode) bool
	//中序遍历
	helper = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !helper(root.Left) {
			return false
		}
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		return helper(root.Right)
	}
	return helper(root)
}
