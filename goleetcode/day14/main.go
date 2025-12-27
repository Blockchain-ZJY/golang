package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

func main() {
	// fmt.Println(multiply("2", "3"))
	// fmt.Println(addStrings("0", "0"))
	// fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 44))
	fmt.Println(minimumBoxes([]int{5, 5, 5}, []int{2, 4, 2, 7}))
}

//977.有序数组的平方

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	index := len(nums) - 1
	for l <= r {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res[index] = nums[l] * nums[l]
			l++
		} else {
			res[index] = nums[r] * nums[r]
			r--
		}
		index--
	}
	return res
}

// 541. 反转字符串 II
func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		if i+k <= len(b) {
			reverse(b[i : i+k])
		} else {
			reverse(b[i:])
		}
	}
	return string(b)
}

func reverse(b []byte) {
	l, r := 0, len(b)
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

// 移除字符串里面的空格(每个单词之间有且只有一个空格)
func removeExtraSpaces(s []byte) ([]byte, int) {
	fast, slow := 0, 0
	for ; fast < len(s); fast++ {
		if s[fast] != ' ' {
			if slow != 0 {
				s[slow] = ' '
				slow++
			}
			for s[fast] != ' ' && fast < len(s) {
				s[slow] = s[fast]
				slow++
				fast++
			}
		}
	}
	return s[:slow], slow
}

// 151. 反转字符串中的单词
func reverseWords(s string) string {
	b := []byte(s)
	b, lenb := removeExtraSpaces(b)
	slices.Reverse(b)
	pre := 0
	for i := 0; i < lenb; i++ {
		for i < lenb && b[i] != ' ' {
			i++
		}
		slices.Reverse(b[pre:i])
		if i+1 < lenb {
			pre = i + 1
		}
	}
	return string(b)
}

func revs(b []byte, l, r int) {
	r = r - 1
	for l < r {
		b[l], b[r] = b[r], b[l]
		r--
		l++
	}
}

//58. 最后一个单词的长度

func lengthOfLastWord(s string) int {
	l := len(s)
	end := l - 1
	for s[end] == ' ' {
		end--
	}
	st := end
	for s[st] != ' ' {
		st--
	}
	return end - st
}

// 43. 字符串相乘--会超时
func multiply(num1 string, num2 string) string {

	var getnum func(a string) int
	getnum = func(a string) int {
		ans := 0
		index := 0
		for i := len(a) - 1; i >= 0; i-- {
			tp := int(a[i] - '0')
			ans += tp * int(math.Pow(float64(10), float64(index)))
			index++
		}
		return ans
	}
	ans := getnum(num1) * getnum(num2)
	res := []byte{}
	if ans == 0 {
		return "0"
	}
	for ans != 0 {
		res = append(res, byte(ans%10)+'0')
		ans /= 10
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

// 模拟乘法 123* 45
func multiply1(num1 string, num2 string) string {
	//记录所有获取的string
	ans := []string{}
	for i := len(num1) - 1; i >= 0; i-- {
		add := 0
		pow := 0
		var temp strings.Builder
		for j := len(num2) - 1; j >= 0; j-- {
			sum := (int(num1[i]-'0')*int(num2[j]-'0') + add)
			pow++
			temp.WriteByte(byte(sum%10) + '0')
			add = sum / 10
		}
		if add != 0 {
			temp.WriteByte(byte(add) + '0')
		}
		res := []byte(temp.String())
		slices.Reverse(res)
		fmt.Println(string(res))
		ans = append(ans, string(res))
	}
	for i := 0; i < len(ans); i++ {
		for j := 0; j < i; j++ {
			ans[i] = ans[i] + "0"
		}
	}
	tepm := "0"
	for i := 0; i < len(ans); i++ {
		tepm = addStrings(tepm, ans[i])
	}
	if tepm[0] == '0' {
		return "0"
	}
	return tepm
}

// 415. 字符串相加
func addStrings(num1 string, num2 string) string {
	n, m := []byte(num1), []byte(num2)
	i, j := len(num1)-1, len(num2)-1
	ans := []byte{}
	sum := 0
	addone := 0
	for i >= 0 || j >= 0 || addone != 0 {
		for i >= 0 && j >= 0 {
			sum = int(n[i]-'0'+m[j]-'0') + addone
			ans = append(ans, byte(sum%10+'0'))
			addone = sum / 10
			i--
			j--
		}
		if i < 0 && j >= 0 {
			for j >= 0 {
				sum = int(m[j]-'0') + addone
				ans = append(ans, byte(sum%10+'0'))
				addone = sum / 10
				j--
			}
		}
		if i >= 0 && j < 0 {
			for i >= 0 {
				sum = int(n[i]-'0') + addone
				ans = append(ans, byte(sum%10+'0'))
				addone = sum / 10
				i--
			}
		}
		if addone != 0 {
			ans = append(ans, '1')
			break
		}
	}
	revs(ans, 0, len(ans))
	return string(ans)
}

func addStrings1(num1, num2 string) string {
	var (
		i, j   = len(num1) - 1, len(num2) - 1 // 指向 num1 和 num2 的指针
		carry  int                            // 进位
		result strings.Builder                // golang 最高效的拼接字符串
	)
	for i >= 0 || j >= 0 || carry > 0 {
		ch1, ch2 := 0, 0
		if i >= 0 {
			ch1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			ch2 = int(num2[j] - '0')
			j--
		}
		sum := ch1 + ch2 + carry
		result.WriteByte(byte(sum%10) + '0')
		carry = sum / 10
	}

	// 结果是逆序的，需要反转
	res := []byte(result.String())
	slices.Reverse(res)
	return string(res)
}

//989. 数组形式的整数加法

func addToArrayForm(num []int, k int) []int {
	l := len(num) - 1
	ans := []int{}
	addone := 0
	for l >= 0 || addone != 0 || k != 0 {
		n1, n2 := 0, 0
		if l >= 0 {
			n1 = int(num[l])
			l--
		}
		if k != 0 {
			n2 = k % 10
			k /= 10
		}
		sum := n1 + n2 + addone
		addone = sum / 10
		ans = append(ans, sum%10)
	}
	slices.Reverse(ans)
	return ans
}

/*
2054. 两个最好的不重叠活动
这个算法的核心是：

	1.按结束时间排序，保证查找时不会漏掉可能的组合。
	2.用栈维护「结束时间递增、价值递增」的活动集合。
	3.每次通过二分查找快速找到能和当前活动不重叠的最佳活动。
*/
func maxTwoEvents(events [][]int) int {
	ans := 0
	slices.SortFunc(events, func(a, b []int) int { // 如果是负数代表 a<b
		return a[1] - b[1] // 小的在前就是升序
	})
	type pair struct {
		endTime, value int
	}
	st := []pair{{}} //哨兵
	for _, e := range events {
		starttime, value := e[0], e[2]
		//要的是最后一个满足条件的,因为越往后它的价值越高
		// sort.Search 返回的是“第一个 true 的位置”，不是“最后一个 true 的位置”。
		i := sort.Search(len(st), func(i int) bool { return st[i].endTime >= starttime }) - 1
		ans = max(ans, st[i].value+value)
		if value > st[len(st)-1].value {
			st = append(st, pair{e[1], value})
		}
	}
	return ans
}

type Node struct {
	Val    int
	Random *Node
	Next   *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		t := &Node{
			Val: cur.Val,
		}
		m[cur] = t
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}

// 3074. 重新分装苹果
func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for i := range apple {
		sum += apple[i]
	}
	sort.Ints(capacity)
	l := len(capacity) - 1
	for sum > 0 {
		sum -= capacity[l]
		l--
	}
	return len(capacity) - l - 1
}
