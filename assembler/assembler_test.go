package assembler

import (
    "testing"

    "github.com/couchand/kismet/instruction"
)

func TestOutput(t *testing.T) {
    program := []instruction.T{
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

    Assemble(program, "testing.ko")
}
