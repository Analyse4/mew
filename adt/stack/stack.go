package stack

type Stack struct {
	buffer []interface{}
}

func New() *Stack {
	return &Stack{make([]interface{}, 0)}
}

func (s *Stack) Push(v interface{}) {
	s.buffer = append(s.buffer, v)
}

func (s *Stack) Pop() interface{} {
	v := s.buffer[len(s.buffer)-1]
	s.buffer = s.buffer[:len(s.buffer)-1]
	return v
}

func (s *Stack) IsEmpty() bool {
	return len(s.buffer) == 0
}

func (s *Stack) Size() int {
	return len(s.buffer)
}
