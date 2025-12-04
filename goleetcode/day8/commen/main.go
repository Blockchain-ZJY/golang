package commen

type Stack struct {
	items []rune
}

// 入栈
func (s *Stack) Push(ch rune) {
	s.items = append(s.items, ch)
}

// 出栈
func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// 查看栈顶
func (s *Stack) Peek() rune {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

// 判断是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// 栈大小
func (s *Stack) Size() int {
	return len(s.items)
}
