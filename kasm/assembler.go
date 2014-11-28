package kasm

import (
    "fmt"

    "github.com/couchand/kismet/instruction"
)

type assembler struct {
    instructions []*InstructionPrototype
    labels LabelIndex
}

func MakeAssembler(is []*InstructionPrototype, ls LabelIndex) assembler {
    return assembler{is, ls}
}

func (a assembler) Layout() {
    offset := 0
    for _, instr := range a.instructions {
        //fmt.Printf("Laying out %v at %v\n", instr, offset)

        instr.Offset = offset
        if instr.Opcode.IsDoubleWide() {
            offset += 2
        } else {
            offset += 1
        }
    }
}

func (a assembler) Assemble() (instructions []instruction.T) {
    instructions = make([]instruction.T, len(a.instructions))

    for i, prototype := range a.instructions {
        instructions[i] = a.assembleInstruction(prototype)
    }

    return
}

func (a assembler) assembleInstruction(p *InstructionPrototype) instruction.T {
    if p.Opcode.IsDoubleWide() {
        return instruction.DoubleWide{p.Opcode, a.resolvePayload(p.Payload)}
    } else {
        return p.Opcode
    }
}

func (a assembler) resolvePayload(p Parameter) int {
    if p.IsInt() {
        return p.IntVal()
    }
    if p.IsVar() {
        n := p.VarName()
        pointer, exists := a.labels[n]

        if !exists {
            msg := fmt.Sprintf("Label '%v' doesn't exist.", n)
            panic(msg)
        }

        return pointer.Offset
    }
    msg := fmt.Sprintf("Unknown parameter type: %v", p)
    panic(msg)
}

func ParseString(s string) (instructions []instruction.T) {
    return parseLexer(MakeStringLexer(s))
}

func ParseFile(s, f string) (instructions []instruction.T) {
    return parseLexer(MakeFileLexer(s, f))
}

func parseLexer(l Lexer) (instructions []instruction.T) {
    p := MakeParser(l)

    parsed, err := p.Parse()
    if err != nil {
        panic(err)
    }

    a := MakeAssembler(parsed, p.labels)
    a.Layout()
    return a.Assemble()
}
