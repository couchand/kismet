package asm

import (
    "testing"

    "github.com/couchand/kismet/instruction"
)

func TestOutput(t *testing.T) {
    expected := []instruction.T{
        instruction.Lit(1),
        instruction.Lit(1),
        instruction.Add,
        instruction.Dup,
        instruction.Lit(10),
        instruction.Sub,
        instruction.If(14),
        instruction.Lit(1),
        instruction.If(1),
    }

    program := "[LIT] 1\n" +
        "[LIT] 1\n" +
        "+\n" +
        "DUP\n" +
        "[LIT] 10\n" +
        "-\n" +
        "[IF] 14\n" +
        "[LIT] 1\n" +
        "[IF] 1\n"

    actual := ParseString(program)

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
