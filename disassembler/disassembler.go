package disassembler

import (
    "fmt"
    "strings"

    "github.com/couchand/kismet/instruction"
)

func instructionToString(t instruction.T) string {
    switch code := t.GetOpcode(); code {
    case instruction.Store:
        return "!"
    case instruction.Add:
        return "+"
    case instruction.Sub:
        return "-"
    case instruction.ToR:
        return ">R"
    case instruction.Fetch:
        return "@"
    case instruction.And:
        return "AND"
    case instruction.Drop:
        return "DROP"
    case instruction.Dup:
        return "DUP"
    case instruction.Or:
        return "OR"
    case instruction.Over:
        return "OVER"
    case instruction.RFrom:
        return "R>"
    case instruction.Swap:
        return "SWAP"
    case instruction.Xor:
        return "XOR"
    case instruction.IfCode:
        if t.GetLength() != 2 {
            panic("expected if jump address")
        }
        return fmt.Sprintf("[IF] %v", t.GetWords()[1])
    case instruction.CallCode:
        if t.GetLength() != 2 {
            panic("expected call jump address")
        }
        return fmt.Sprintf("[CALL] %v", t.GetWords()[1])
    case instruction.Exit:
        return "[EXIT]"
    case instruction.LitCode:
        if t.GetLength() != 2 {
            panic("expected literal")
        }
        return fmt.Sprintf("[LIT] %v", t.GetWords()[1])
    default:
        msg := fmt.Sprintf("unknown opcode %v", code)
        panic(msg)
    }
}

func Disassemble(instructions []instruction.T) string {
    lines := make([]string, len(instructions))
    for i, instr := range instructions {
        lines[i] = instructionToString(instr)
    }
    return strings.Join(lines, "\n") + "\n"
}
