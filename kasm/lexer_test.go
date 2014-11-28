package kasm

import (
    "testing"

    //"fmt"

    "github.com/couchand/kismet/instruction"
)

func TestOutput(t *testing.T) {
    program := "1\n" +
        "1\n" +
        "+\n" +
        "DUP\n" +
        "10\n" +
        "-\n" +
        "IF(14)\n" +
        "1\n" +
        "IF(1)\n"

    l := MakeStringLexer(program)

    //fmt.Printf("Lexer: %v\n", l)

    tok, err := l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
    }
    i, err := tok.Integer()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
    }
    if i != 1 {
        t.Errorf("Expected literal token, got %v", i)
    }
    if tok.LineNumber() != 1 {
        t.Errorf("Expected token to be on line 1")
    }
    if tok.ColumnNumber() != 1 {
        t.Errorf("Expected token to be in column 1")
    }

    //fmt.Printf("Lexer: %v\n", l)

    tok, err = l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
    }
    i, err = tok.Integer()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
    }
    if i != 1 {
        t.Errorf("Expected literal token, got %v", i)
    }
    if tok.LineNumber() != 2 {
        t.Errorf("Expected token to be on line 2")
    }
    if tok.ColumnNumber() != 1 {
        t.Errorf("Expected token to be in column 1")
    }

    //fmt.Printf("Lexer: %v\n", l)

    tok, err = l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
    }
    in, err := tok.Instruction()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
    }
    if in != instruction.Add {
        t.Errorf("Expected add instruction, got %v", i)
    }
    if tok.LineNumber() != 3 {
        t.Errorf("Expected token to be on line 3")
    }
    if tok.ColumnNumber() != 1 {
        t.Errorf("Expected token to be in column 1")
    }

    //fmt.Printf("Lexer: %v\n", l)
}
