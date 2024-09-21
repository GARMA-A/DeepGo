package main

type stackImpl struct{
	stack []int
}

func (s *stackImpl) push(val int){
	s.stack = append(s.stack, val)
}

func (s *stackImpl) pop() int{
	if s.isEmpty(){
		return -1
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}

func (s *stackImpl) peek() int{
	if s.isEmpty(){
		return -1
	}
	return s.stack[len(s.stack)-1]
}

func (s *stackImpl) isEmpty() bool{
	return len(s.stack) == 0
}

