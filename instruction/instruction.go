package instruction

import "fmt"

type Opcode int

const (
    Halt Opcode = iota
    Store
    Add
    Sub
    ToR
    Fetch
    And
    Drop
    Dup
    Or
    Over
    RFrom
    Swap
    Xor
    IfCode
    CallCode
    Exit
    LitCode
    Debug
)

func IsOpcode(d int) bool {
    return d >= int(Halt) && d <= int(Debug)
}

func IsDoubleWideOpcode(d int) bool {
    return d == int(IfCode) || d == int(CallCode) || d == int(LitCode)
}

type T interface {
    GetOpcode() Opcode
    GetLength() int
    GetWords() []int
}

func (o Opcode) GetOpcode() Opcode {
    return o
}

func (o Opcode) GetLength() int {
    return 1
}

func (o Opcode) GetWords() []int {
    return []int{ int(o) }
}

func (o Opcode) IsDoubleWide() bool {
    return o == IfCode || o == CallCode || o == LitCode
}

func (o Opcode) String() string {
    switch o {
    case Halt:
        return "HALT"
    case Store:
        return "!"
    case Add:
        return "+"
    case Sub:
        return "-"
    case ToR:
        return ">R"
    case Fetch:
        return "@"
    case And:
        return "AND"
    case Drop:
        return "DROP"
    case Dup:
        return "DUP"
    case Or:
        return "OR"
    case Over:
        return "OVER"
    case RFrom:
        return "R>"
    case Swap:
        return "SWAP"
    case Xor:
        return "XOR"
    case IfCode:
        return "[IF]"
    case CallCode:
        return "[CALL]"
    case Exit:
        return "[EXIT]"
    case LitCode:
        return "[LIT]"
    case Debug:
        return "DEBUG"
    default:
        msg := fmt.Sprintf("unknown opcode %v", int(o))
        panic(msg)
    }
}

type DoubleWide struct {
    Opcode
    Payload int
}

func (d DoubleWide) GetOpcode() Opcode {
    return d.Opcode
}

func (d DoubleWide) GetLength() int {
    return 2
}

func (d DoubleWide) GetWords() []int {
    return []int{ int(d.Opcode), d.Payload }
}

func If(p int) T {
    return DoubleWide{IfCode, p}
}

func Call(p int) T {
    return DoubleWide{CallCode, p}
}

func Lit(p int) T {
    return DoubleWide{LitCode, p}
}
