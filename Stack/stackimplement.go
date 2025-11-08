package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) { //insert value
	s.items = append(s.items, i)
	fmt.Println("pushed", i)

}
func (s *Stack) Pop() int { //delete value
	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		return -1
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	fmt.Println("poped", last)
	return last
}
func (s *Stack) Peek() int { //show top element without removing
	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]

}

func main() {
	var s Stack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("stack now", s.items)
	fmt.Println("top element", s.Peek())

	s.Pop()
	fmt.Println("stack after one pop", s.items)
	s.Pop()
	fmt.Println("top element", s.Peek())
	s.Pop()

	s.Pop()

}
