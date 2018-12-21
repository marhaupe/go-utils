package pkg

// Stack is a Stack implementation
type Stack struct {
	top int
	el  []interface{}
}

// New initializes a Stack ready to be used
func New() *Stack {
	return &Stack{
		top: -1,
		el:  make([]interface{}, 0, 5),
	}
}

// Pop removes and returns the topmost element and true or returns nil and false
func (s *Stack) Pop() (interface{}, bool) {
	if s.top == -1 {
		return nil, false
	}
	popped := s.el[s.top]
	s.el = s.el[:s.top]
	s.top--
	return popped, true
}

// Push puts i on top of the stack
func (s *Stack) Push(i interface{}) {
	s.el = append(s.el, i)
	s.top++
}

// Peek returns the topmost element and true or returns nil and false, but does not remove
// anything from the stack
func (s *Stack) Peek() (interface{}, bool) {
	if s.top == -1 {
		return nil, false
	}
	return s.el[s.top], true
}
