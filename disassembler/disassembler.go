package disassembler

import (
    "fmt"
    "strings"

    "github.com/couchand/kismet/instruction"
)

func instructionToString(t instruction.T) string {
    code := t.GetOpcode()

    if code.IsDoubleWide() {
        if t.GetLength() != 2 {
            panic("expected instruction parameter")
        }
        return fmt.Sprintf("%v\n%v", code.String(), t.GetWords()[1])
    }

    return code.String()
}

func Disassemble(instructions []instruction.T) string {
    lines := make([]string, len(instructions))
    for i, instr := range instructions {
        lines[i] = instructionToString(instr)
    }
    return strings.Join(lines, "\n") + "\n"
}
