package lifo

import (
	"fmt"
	"testing"
)

type LifoIntf interface {
	initialize()
	Pop() interface{}
	PrintStack() string
	PrintTop() string
	Append(element interface{})
}

type Lifo struct {
	stack []interface{}
}

func (l *Lifo) initialize() {
	l.stack = make([]interface{}, 0)
}

func (l *Lifo) Pop() interface{} {

	if l.stack == nil {
		l.initialize()
	}

	if len(l.stack) > 0 {
		return l.stack[len(l.stack)-1]
	} else {
		return nil
	}
}

func (l *Lifo) Append(element interface{}) {

	if l.stack == nil {
		l.initialize()
	}
	l.stack = append(l.stack, element)
}

func (l *Lifo) PrintStack() string {

	stack := ""

	if l.stack == nil {
		return ""
	} else {
		for i := len(l.stack) - 1; i >= 0; i-- {
			stack += fmt.Sprintf("%v\n", l.stack[i])
		}
	}
	return stack
}

func (l *Lifo) PrintTop() string {

	if l.stack == nil {
		return ""
	} else {
		return fmt.Sprintf("%v", l.stack[len(l.stack)-1])
	}
}

func TestAppend(t *testing.T) {

	l := Lifo{}
	elements := []interface{}{1, 2, "element", 1.45}
	expectedLength := 0

	for i, val := range elements {

		if expectedLength != len(l.stack) {
			t.Errorf("Expected length %d, got %d", expectedLength, len(l.stack))
		}

		l.Append(val)
		expectedLength++

		if l.stack[i] != val {
			t.Errorf("Expected value %v at position %v, got %v", val, i, l.stack[i])
		}
	}
}

func TestPop(t *testing.T) {

	elements := []interface{}{1, 2, "element", 1.45}
	l := Lifo{}

	// Nothing was added so the first pop should return nil
	popped := l.Pop()
	if popped != nil {
		t.Errorf("Stack was empty, expected nil got %v", popped)
	}

	for _, val := range elements {

		l.stack = append(l.stack, val)
		popped = l.Pop()

		if popped != val {
			t.Errorf("Expected to Pop value %v, got %v", val, popped)
		}
	}
}

func TestPrintTop(t *testing.T) {

	elements := []interface{}{1, 2, "element", 1.45}
	l := Lifo{}

	got := l.PrintTop()
	if got != "" {
		t.Errorf("Expected \"\", got %v", got)
	}

	for _, val := range elements {
		l.stack = append(l.stack, val)

		want := fmt.Sprintf("%v", val)
		got := l.PrintTop()
		if want != got {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}
}

func TestPrintStack(t *testing.T) {

	want := ""
	elements := []interface{}{1, 2, "element", 1.45}
	l := Lifo{}

	got := l.PrintStack()
	if got != want {
		t.Errorf("Expected \"%v\", got %v", want, got)
	}

	for _, val := range elements {
		l.stack = append(l.stack, val)
		want = fmt.Sprintf("%v\n%v", val, want)
		got = l.PrintStack()

		if want != got {
			t.Errorf("Expected %s, got %s", want, got)
		}
	}
}
