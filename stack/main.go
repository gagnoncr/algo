package main

import (
	"fmt"
)

type Stack struct {
	slice []int
}

func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

func (s *Stack) Pop() int {
	var ret int = s.Peek()
	s.slice = s.slice[0:len(s.slice)-1]
	return ret
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}




func main() {
	var s *Stack = new(Stack)
	s.Push(200)
	s.Push(3)
	s.Push(44)
	s.Push(23)
	s.Push(1)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(99)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}