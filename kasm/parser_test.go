package kasm

import (
    "testing"

    //"fmt"

    "github.com/couchand/kismet/instruction"
)

func TestParser(t *testing.T) {
    program := "1\n" +
        "loop:1\n" +
        "+\n" +
        "DUP\n" +
        "10\n" +
        "-\n" +
        "JZ(end)\n" +
        "1\n" +
        "JZ(loop)\n" +
        "end:0\n"

    l := MakeStringLexer(program)
    p := MakeParser(l)

    is, err := p.Parse()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }

    if is[0].Opcode != instruction.LitCode {
        t.Errorf("Expected literal instruction")
    }
    if is[0].Payload.IntVal() != 1 {
        t.Errorf("Expected literal value 1")
    }
    if is[1].Opcode != instruction.LitCode {
        t.Errorf("Expected literal instruction")
    }
    if is[1].Payload.IntVal() != 1 {
        t.Errorf("Expected literal value 1")
    }
    if is[2].Opcode != instruction.Add {
        t.Errorf("Expected add instruction")
    }
}

func TestLabelLiteral(t *testing.T) {
    program := "0\n 1\n two:+\n [two]\n"

    l := MakeStringLexer(program)
    p := MakeParser(l)

    is, err := p.Parse()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }

    if is[0].Opcode != instruction.LitCode {
        t.Errorf("Expected literal instruction")
    }
    if is[0].Payload.IntVal() != 0 {
        t.Errorf("Expected literal value 0")
    }
    if is[3].Opcode != instruction.LitCode {
        t.Errorf("Expected literal instruction")
    }
    if is[3].Payload.VarName() != "two" {
        t.Errorf("Expected label name two")
    }
}
