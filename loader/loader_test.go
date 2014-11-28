package loader

import (
    "testing"

//    "fmt"

    "github.com/couchand/kismet/instruction"
)

func TestOutput(t *testing.T) {
    expected := []instruction.T{
        instruction.Lit(1),
        instruction.Dup,
        instruction.Lit(1),
        instruction.Add,
        instruction.Dup,
        instruction.Lit(10),
        instruction.Sub,
        instruction.If(14),
        instruction.Lit(0),
        instruction.If(2),
    }

    actual, err := Load("testing.ko")

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    //fmt.Printf("Loaded: %v\n", actual)

    if len(actual) != len(expected) {
        t.Errorf("Expected same length of program")
        return
    }

    for i := range actual {
        e := expected[i].GetOpcode()
        a := actual[i].GetOpcode()
        if a != e {
            t.Errorf("Expected opcode %v, got %v", e, a)
            return
        }

        if e.GetLength() == 2 {
            if a.GetLength() != 2 {
                t.Errorf("Expected double wide")
                return
            }

            if e.GetWords()[1] != a.GetWords()[1] {
                t.Errorf("Expected literal value")
                return
            }
        }
    }
}
