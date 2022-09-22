package lifo

import "fmt"

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
