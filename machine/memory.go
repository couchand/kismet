package machine

import (
    "fmt"
)

type Memory interface {
    SetAddress(int)
    GetData() int
    SetData(int)
}

type Programmable interface {
    Memory
    Program([]int)
}

type tenK struct {
    address int
    data [1024]int
}

func (m *tenK) String() string {
    return fmt.Sprintf("Addr: 0x%x Data: 0x%x", m.address, m.data[m.address])
}

func (m *tenK) SetAddress(a int) {
    if a >= 1024 {
        msg := fmt.Sprintf("address out of bounds: %v", a)
        panic(msg)
    }

    m.address = a
}

func (m *tenK) GetData() int {
    return m.data[m.address]
}

func (m *tenK) SetData(v int) {
    m.data[m.address] = v
}

func (m *tenK) Program(ps []int) {
    if len(ps) >= 1024 {
        msg := fmt.Sprintf("program too large: %v", len(ps))
        panic(msg)
    }

    for i, d := range ps {
        m.data[i] = d
    }
}

func Make10KProgrammable() Programmable {
    return new(tenK)
}
