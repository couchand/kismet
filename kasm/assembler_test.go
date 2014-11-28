package kasm

import (
    "testing"

    //"fmt"

    "github.com/couchand/kismet/instruction"
)

func TestAssembler(t *testing.T) {
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
    }

    a := MakeAssembler(is, p.labels)
    a.Layout()

    assembled := a.Assemble()

    if assembled[0].GetOpcode() != instruction.LitCode {
        t.Errorf("Expected literal instruction")
    }
    if assembled[1].GetOpcode() != instruction.LitCode {
        t.Errorf("Expected literal instruction")
    }
    if assembled[2].GetOpcode() != instruction.Add {
        t.Errorf("Expected add instruction")
    }

    words := assembled[6].GetWords()
    if words[1] != 15 {
        t.Errorf("Expected address resolved to 14, got %v", words[1])
    }
}
