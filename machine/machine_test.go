package machine

import (
    "testing"

    "github.com/couchand/kismet/asm"
)

func TestMachine(t *testing.T) {
    m := Make10KMachine(asm.ParseString("[LIT] 1 [LIT] 2")).(*mac)

    m.Step()

    if m.ds.Len() != 1 {
        t.Errorf("Expected stack to have one element")
        return
    }

    v := m.ds.Pop()
    if v != 1 {
        t.Errorf("Expected stack value to be 1, got %v", v)
        return
    }

    m.Step()

    if m.ds.Len() != 1 {
        t.Errorf("Expected stack to have one element")
        return
    }

    v = m.ds.Pop()
    if v != 2 {
        t.Errorf("Expected stack value to be 2, got %v", v)
        return
    }
}
