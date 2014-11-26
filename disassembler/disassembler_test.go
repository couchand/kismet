package disassembler

import (
    "testing"

    "github.com/couchand/kismet/instruction"
)

func TestOutput(t *testing.T) {
    program := []instruction.T{
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

    expected := "[LIT] 1\n" +
        "[LIT] 1\n" +
        "+\n" +
        "DUP\n" +
        "[LIT] 10\n" +
        "-\n" +
        "[IF] 14\n" +
        "[LIT] 1\n" +
        "[IF] 1\n"

    actual := Disassemble(program)

    if actual != expected {
        t.Errorf("Expected: %v\nSaw: %v\n", expected, actual)
        return
    }
}
