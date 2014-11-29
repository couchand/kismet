package machine

import (
    "fmt"

    "github.com/couchand/kismet/instruction"
)

type Machine interface {
    Execute()
    Step() bool
}

type mac struct {
    ds Stack
    rs Stack
    mem Memory

    pc int
}

func Make10KMachine(program []instruction.T) Machine {
    ints := make([]int, 0, len(program))
    for _, p := range program {
        for _, w := range p.GetWords() {
            ints = append(ints, w)
        }
    }

    m := Make10KProgrammable()
    m.Program(ints)

    return &mac{
        MakeSliceStack(12),
        MakeSliceStack(4),
        m,
        0,
    }
}

func (m *mac) String() string {
    mem := fmt.Sprintf("%v", m.mem)
    m.mem.SetAddress(m.pc)
    this := m.mem.GetData()
    m.mem.SetAddress(m.pc + 1)
    next := m.mem.GetData()
    return fmt.Sprintf("10K Machine\nProgram Counter: 0x%x Value: 0x%x %x\nData Stack: %v\nReturn Stack: %v\nMemory: %v\n", m.pc, this, next, m.ds, m.rs, mem)
}

func (m *mac) Execute() {
    for {
        done := m.Step()
        if done {
            return
        }
    }
}

func (m *mac) Step() bool {
    m.mem.SetAddress(m.pc)
    inst := m.mem.GetData()
    m.pc = m.pc + 1

    if !instruction.IsOpcode(inst) {
        msg := fmt.Sprintf("illegal instruction %v", inst)
        panic(msg)
    }

    switch op := instruction.Opcode(inst); op {
    case instruction.Halt:
        return true
    case instruction.Store:
        m.mem.SetAddress(m.ds.Pop())
        m.mem.SetData(m.ds.Pop())
    case instruction.Add:
        a, b := m.ds.Pop(), m.ds.Pop()
        m.ds.Push(b + a)
    case instruction.Sub:
        a, b := m.ds.Pop(), m.ds.Pop()
        m.ds.Push(b - a)
    case instruction.ToR:
        m.rs.Push(m.ds.Pop())
    case instruction.Fetch:
        m.mem.SetAddress(m.ds.Pop())
        m.ds.Push(m.mem.GetData())
    case instruction.And:
        a, b := m.ds.Pop(), m.ds.Pop()
        m.ds.Push(a & b)
    case instruction.Drop:
        m.ds.Pop()
    case instruction.Dup:
        d := m.ds.Pop()
        m.ds.Push(d)
        m.ds.Push(d)
    case instruction.Or:
        a, b := m.ds.Pop(), m.ds.Pop()
        m.ds.Push(a | b)
    case instruction.Over:
        top := m.ds.Pop()
        d := m.ds.Pop()
        m.ds.Push(d)
        m.ds.Push(top)
        m.ds.Push(d)
    case instruction.RFrom:
        m.ds.Push(m.rs.Pop())
    case instruction.Swap:
        top := m.ds.Pop()
        d := m.ds.Pop()
        m.ds.Push(top)
        m.ds.Push(d)
    case instruction.Xor:
        a, b := m.ds.Pop(), m.ds.Pop()
        m.ds.Push(a ^ b)
    case instruction.IfCode:
        condition := m.ds.Pop()
        if condition == 0 {
            m.mem.SetAddress(m.pc)
            m.pc = m.mem.GetData()
        } else {
            m.pc = m.pc + 1
        }
    case instruction.CallCode:
        m.mem.SetAddress(m.pc)
        m.pc = m.pc + 1
        m.rs.Push(m.pc)
        m.pc = m.mem.GetData()
    case instruction.Exit:
        m.pc = m.rs.Pop()
    case instruction.LitCode:
        m.mem.SetAddress(m.pc)
        m.pc = m.pc + 1
        m.ds.Push(m.mem.GetData())
    default:
        msg := fmt.Sprintf("unknown opcode: %v", op)
        panic(msg)
    }

    return false
}
