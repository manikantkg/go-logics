package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack data structure
type Stack struct {
	elements []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack
// Returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	topIndex := len(s.elements) - 1
	topElement := s.elements[topIndex]
	s.elements = s.elements[:topIndex]
	return topElement, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Top returns the top element of the stack without removing it
// Returns an error if the stack is empty
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Top element:", stack.Top())

	for !stack.IsEmpty() {
		topElement, err := stack.Pop()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println("Popped element:", topElement)
	}

	// Trying to pop from an empty stack
	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
