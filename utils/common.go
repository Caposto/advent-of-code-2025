package utils

type Stack[T any] []T

func (s *Stack[T]) Push(item  T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T // Return zero value for the type if empty
        return zero, false
    }
    index := len(*s) - 1
    item := (*s)[index]
    *s = (*s)[:index] // Slice off the top element
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }
    index := len(*s) - 1
    return (*s)[index], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(*s) == 0
}

func Check(e error) {
    if e != nil {
        panic(e)
    }
}