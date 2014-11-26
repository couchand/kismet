package machine

import (
    "fmt"
    "strings"
)

type Stack interface {
    Push(int)
    Pop() int
    Len() int
}

type sliceStack struct {
    data []int
}

func (s *sliceStack) String() string {
    ds := make([]string, len(s.data))
    for i, d := range s.data {
        ds[i] = fmt.Sprintf("%x", d)
    }
    return fmt.Sprintf("(%v)", strings.Join(ds, " "))
}

func (s *sliceStack) Push(v int) {
    s.data = append(s.data, v)
}

func (s *sliceStack) Pop() int {
    if len(s.data) == 0 {
        panic("stack underflow")
    }

    last := len(s.data) - 1
    d := s.data[last]
    s.data = s.data[:last]
    return d
}

func (s *sliceStack) Len() int {
    return len(s.data)
}

func MakeSliceStack(initialCapacity int) Stack {
    return &sliceStack{make([]int, 0, initialCapacity)}
}
