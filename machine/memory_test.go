package machine

import (
    "testing"
)

func TestStore(t *testing.T) {
    m := Make10KProgrammable()

    m.SetAddress(42)
    m.SetData(42)

    m.SetAddress(666)
    m.SetData(65535)

    m.SetAddress(42)

    if d := m.GetData(); d != 42 {
        t.Errorf("Expected stored data, got %v", d)
    }

    m.SetAddress(666)

    if d := m.GetData(); d != 65535 {
        t.Errorf("Expected stored data, got %v", d)
    }
}

func TestProgram(t *testing.T) {
    m := Make10KProgrammable()

    m.Program([]int{42, 666, 1337, 65535})

    m.SetAddress(0)
    if d := m.GetData(); d != 42 {
        t.Errorf("Expected stored data, got %v", d)
    }

    m.SetAddress(1)
    if d := m.GetData(); d != 666 {
        t.Errorf("Expected stored data, got %v", d)
    }

    m.SetAddress(2)
    if d := m.GetData(); d != 1337 {
        t.Errorf("Expected stored data, got %v", d)
    }

    m.SetAddress(3)
    if d := m.GetData(); d != 65535 {
        t.Errorf("Expected stored data, got %v", d)
    }
}
