package instruction

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
)

func IsOpcode(d int) bool {
    return d >= int(Halt) && d <= int(LitCode)
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
