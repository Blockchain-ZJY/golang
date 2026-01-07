package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
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

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) (ans *TreeNode) {
	var hasnode func(node, p, q *TreeNode) int
	// 从该节点往下是否包含p,q节点,找到最接近的
	hasnode = func(node, p, q *TreeNode) (res int) {
		if node == nil {
			return 0
		}
		if node == q || node == p {
			res++
		}
		res += hasnode(node.Left, p, q) + hasnode(node.Right, p, q)
		if res == 2 && ans == nil {
			ans = node
		}
		return
	}
	hasnode(root, p, q)
	return ans
}

// 257. 二叉树的所有路径
func binaryTreePaths(root *TreeNode) (ans []string) {
	var dfs func(path string, node *TreeNode)
	dfs = func(path string, node *TreeNode) {
		if node == nil {
			return
		}
		path = path + strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			ans = append(ans, path)
			return
		}
		path += "->"
		dfs(path, root.Left)
		dfs(path, root.Right)
	}
	dfs("", root)
	return
}

// 1161. 最大层内元素和
func maxLevelSum(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	bigerone := -math.MaxInt
	q := []*TreeNode{}
	q = append(q, root)
	index := 0
	for len(q) != 0 {
		layer := []int{}
		len := len(q)
		index++
		sum := 0
		for i := 0; i < len; i++ {
			head := q[0]
			sum += head.Val
			q = q[1:]
			layer = append(layer, head.Val)
			if head.Left != nil {
				q = append(q, head.Left)
			}
			if head.Right != nil {
				q = append(q, head.Right)
			}
		}
		if bigerone < sum {
			ans = index
			bigerone = sum
		}
		fmt.Println(bigerone)
	}
	return
}

// 199. 二叉树的右视图
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	//层次遍历收集最后一个结果
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			t := q[0]
			if i == l-1 {
				ans = append(ans, t.Val)
			}
			q = q[1:] // 出队
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
	}
	return
}

// 114. 二叉树展开为链表 先序遍历相同,先序遍历记录pre,每次
func flatten(root *TreeNode) {
	pre := &TreeNode{}
	var helpfunc func(root *TreeNode)
	helpfunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		if pre != nil {
			pre.Right = root
			pre.Left = nil
		}
		right := root.Right
		pre = root
		helpfunc(root.Left)
		helpfunc(right)
	}
	helpfunc(root)
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	if len(inorder) == 1 {
		return root
	}
	// 找到中序左边的数组
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	inorderleft := inorder[:index]
	inorderright := inorder[index+1:] // ⭐ 修复这里
	preorderleft := preorder[1 : 1+index]
	preorderright := preorder[1+index:]
	root.Left = buildTree(preorderleft, inorderleft)
	root.Right = buildTree(preorderright, inorderright)
	return root
}

// 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return (targetSum-root.Val == 0)
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 路径总和
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var path []int

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)

		path = path[:len(path)-1] // 回溯
	}

	dfs(root, 0)
	return ans
}

// 437. 路径总和 III
func pathSum1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	//函数定义 由node节点开头,后续链路和为targetSum的数量
	var dfs func(node *TreeNode, sum int) int
	dfs = func(node *TreeNode, sum int) (ct int) {
		if node == nil {
			return 0
		}
		if sum-node.Val == 0 {
			ct++
		}
		ct += dfs(node.Left, sum-node.Val) + dfs(node.Right, sum-node.Val)
		return
	}
	return dfs(root, targetSum) + pathSum1(root.Left, targetSum) + pathSum1(root.Right, targetSum)
}
