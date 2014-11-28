package kasm

import (
    "testing"

    "github.com/couchand/kismet/instruction"
)

func TestOutput(t *testing.T) {

    program := "1\n" +
        "1\n" +
        "+\n" +
        "dup\n" +
        "10\n" +
        "-\n" +
        "jz(14)\n" +
        "1\n" +
        "jz(1)\n"

    l := MakeStringLexer(program)

    //fmt.Printf("Lexer: %v\n", l)

    tok, err := l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    i, err := tok.Integer()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
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
        return
    }
    i, err = tok.Integer()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
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
        return
    }
    in, err := tok.Instruction()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    if in != instruction.Add {
        t.Errorf("Expected add instruction, got %v", in)
    }
    if tok.LineNumber() != 3 {
        t.Errorf("Expected token to be on line 3")
    }
    if tok.ColumnNumber() != 1 {
        t.Errorf("Expected token to be in column 1")
    }

    //fmt.Printf("Lexer: %v\n", l)

    tok, err = l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    in, err = tok.Instruction()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    if in != instruction.Dup {
        t.Errorf("Expected dup instruction, got %v", in)
    }
    if tok.LineNumber() != 4 {
        t.Errorf("Expected token to be on line 4")
    }
    if tok.ColumnNumber() != 1 {
        t.Errorf("Expected token to be in column 1")
    }

    //fmt.Printf("Lexer: %v\n", l)
}

func TestComments(t *testing.T) {
    program := "# this is a program\n0 # nothing\n 1 # something\n"

    l := MakeStringLexer(program)

    tok, err := l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    i, err := tok.Integer()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    if i != 0 {
        t.Errorf("Expected 0, got %v", i)
    }
    if tok.LineNumber() != 2 {
        t.Errorf("Expected token to be on line 2")
    }
    if tok.ColumnNumber() != 1 {
        t.Errorf("Expected token to be in column 1")
    }

    tok, err = l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    i, err = tok.Integer()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    if i != 1 {
        t.Errorf("Expected 1, got %v", i)
    }
    if tok.LineNumber() != 3 {
        t.Errorf("Expected token to be on line 3")
    }
    if tok.ColumnNumber() != 2 {
        t.Errorf("Expected token to be in column 2")
    }

    tok, err = l.GetToken()
    if err != nil {
        t.Errorf("Didn't expect error, got %v", err)
        return
    }
    if tok.Type() != EOFToken {
        t.Errorf("Expected EOF, got %v", tok)
    }
}
