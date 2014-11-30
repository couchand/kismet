package asm

import (
//    "io"
//    "os"
    "fmt"
//    "bufio"
    "regexp"
    "strconv"

    "github.com/couchand/kismet/instruction"
)

type Lexer interface {
    GetToken() int
}

const tokenExpression string = "(!|\\+|-|>R|@|AND|DROP|DUP|OR|OVER|R>|SWAP|XOR|\\[IF]|\\[CALL]|\\[EXIT]|\\[LIT]|DEBUG|[0-9]+)"
var tokenRE *regexp.Regexp = regexp.MustCompile(tokenExpression)

const whitespaceExpression string = "^[ \t\n\r]$"
var whitespaceRE *regexp.Regexp = regexp.MustCompile(whitespaceExpression)

type strLexer struct {
    string
}

func MakeStringLexer(s string) Lexer {
    return &strLexer{s}
}

func (s *strLexer) GetToken() int {
    loc := tokenRE.FindStringIndex(s.string)
    if loc == nil {
        if len(s.string) > 0 && !whitespaceRE.MatchString(s.string) {
            msg := fmt.Sprintf("illegal input found: %v", s.string)
            panic(msg)
        }

        return -1
    }

    if (loc == nil || loc[0] != 0) {
        if !whitespaceRE.MatchString(s.string[:loc[0]]) {
            msg := fmt.Sprintf("illegal input found: %v", s.string[:loc[0]])
            panic(msg)
        }
    }

    t := s.string[loc[0]:loc[1]]
    s.string = s.string[loc[1]:]

    switch t {
    case "!":
        return int(instruction.Store)
    case "+":
        return int(instruction.Add)
    case "-":
        return int(instruction.Sub)
    case ">R":
        return int(instruction.ToR)
    case "@":
        return int(instruction.Fetch)
    case "AND":
        return int(instruction.And)
    case "DROP":
        return int(instruction.Drop)
    case "DUP":
        return int(instruction.Dup)
    case "OR":
        return int(instruction.Or)
    case "OVER":
        return int(instruction.Over)
    case "R>":
        return int(instruction.RFrom)
    case "SWAP":
        return int(instruction.Swap)
    case "XOR":
        return int(instruction.Xor)
    case "[IF]":
        return int(instruction.IfCode)
    case "[CALL]":
        return int(instruction.CallCode)
    case "[EXIT]":
        return int(instruction.Exit)
    case "[LIT]":
        return int(instruction.LitCode)
    case "DEBUG":
        return int(instruction.Debug)
    default:
        val, err := strconv.ParseInt(t, 10, 32)
        if err != nil {
            panic(err)
        }
        return int(val)
    }
}

func ParseLexer(l Lexer) (instructions []instruction.T) {
    instructions = make([]instruction.T, 0)
    for {
        t := l.GetToken()
        switch t {
        case -1:
            return
        case int(instruction.IfCode), int(instruction.CallCode), int(instruction.LitCode):
            v := l.GetToken()
            if v == -1 {
                panic("early end of input, expected param")
            }
            i := instruction.DoubleWide{instruction.Opcode(t), v}
            instructions = append(instructions, i)
        default:
            instructions = append(instructions, instruction.Opcode(t))
        }
    }
}

func ParseString(s string) (instructions []instruction.T) {
    return ParseLexer(MakeStringLexer(s))
}
