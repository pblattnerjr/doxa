package stack

import (
	"testing"
)

func TestInt(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}
	s.Push(1)
	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}
	if s.Peek().(int) != 1 {
		t.Errorf("Top item on the stack should be 1")
	}
	if s.Pop().(int) != 1 {
		t.Errorf("Top item should have been 1")
	}
	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
	s.Push(1)
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}
	if s.Peek().(int) != 2 {
		t.Errorf("Top of the stack should be 2")
	}
}
func TestString(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}
	s.Push("one")
	if s.Len() != 1 {
		t.Errorf("Length should be 1")
	}
	if s.Peek().(string) != "one" {
		t.Errorf("Top item on the stack should be `one`")
	}
	if s.Pop().(string) != "one" {
		t.Errorf("Top item should have been `one`")
	}
	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
	s.Push("one")
	s.Push("two")
	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}
	if s.Peek().(string) != "two" {
		t.Errorf("Top of the stack should be `two`")
	}
}
func TestStruct(t *testing.T) {
	type Point struct {
		x int
		y int
	}
	s := New()
	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}
	var p1 = Point{0, 1}
	s.Push(p1)
	if s.Len() != 1 {
		t.Errorf("Length should be 1")
	}
	if s.Peek().(Point).x != 0 {
		t.Errorf("Top item on the stack should have x = 0")
	}
	if s.Pop().(Point).y != 1 {
		t.Errorf("Top item should have y = 1")
	}
	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
	s.Push(p1)
	var p2 = Point{1, 2}
	s.Push(p2)
	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}
	if s.Peek().(Point).x != 1 {
		t.Errorf("Top of the stack should have x = 1")
	}
}
