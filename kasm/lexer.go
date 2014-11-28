package kasm

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"

    "github.com/couchand/kismet/instruction"
)

type TokenType int

const (
    EOFToken TokenType = iota
    InstructionToken
    IntegerToken
    LabelToken
    ParenOpenToken
    ParenCloseToken
    ColonToken
)

type Token interface {
    Type() TokenType

    Raw() string
    Instruction() (instruction.Opcode, error)
    Integer() (int, error)
    Label() (string, error)

    FileName() string
    LineNumber() int
    ColumnNumber() int
}

type Lexer interface {
    GetToken() (Token, error)
}

func makeMatcher(expr string) (*regexp.Regexp) {
    return regexp.MustCompile("^" + expr + "$")
}

const colonExpression string = ":"
var colonRE *regexp.Regexp = makeMatcher(colonExpression)

const parenOpenExpression string = "\\("
var parenOpenRE *regexp.Regexp = makeMatcher(parenOpenExpression)

const parenCloseExpression string = "\\)"
var parenCloseRE *regexp.Regexp = makeMatcher(parenCloseExpression)

const integerExpression string = "[0-9]+"
var integerRE *regexp.Regexp = makeMatcher(integerExpression)

const labelExpression string = "[a-zA-Z_][a-zA-Z_0-9]*"
var labelRE *regexp.Regexp = makeMatcher(labelExpression)

const instructionExpression string = "(!|\\+|-|>R|@|&|DROP|DUP|\\||OVER|R>|SWAP|XOR|JZ|CALL|RETURN)"
var instructionRE *regexp.Regexp = makeMatcher(instructionExpression)

const whitespaceExpression string = "[ \t\n\r]"
var whitespaceRE *regexp.Regexp = makeMatcher(whitespaceExpression)

const newlineExpression string = "(\n\r|\r\n|\n|\r)"
var newlineRE *regexp.Regexp = makeMatcher(whitespaceExpression)

var opts = []string{
    colonExpression,
    parenOpenExpression,
    parenCloseExpression,
    integerExpression,
    instructionExpression,
    labelExpression,
}
var options = strings.Join(opts, "|")
var nextTokenRE = regexp.MustCompile("(" + options + ")")

type rawToken struct {
    raw string
    typ TokenType
    file string
    line int
    column int
}

func (r rawToken) Type() TokenType {
    return r.typ
}

func (r rawToken) Raw() string {
    return r.raw
}

func (r rawToken) FileName() string {
    return r.file
}

func (r rawToken) LineNumber() int {
    return r.line
}

func (r rawToken) ColumnNumber() int {
    return r.column
}

type lexerError string

func (l lexerError) Error() string {
    return string(l)
}

func makeError(token Token, format string, params... interface{}) error {
    ps := make([]interface{}, len(params) + 3)
    ps[len(params) + 0] = token.FileName()
    ps[len(params) + 1] = token.LineNumber()
    ps[len(params) + 2] = token.ColumnNumber()
    msg := fmt.Sprintf(format + " (%v, line %v, column %v)", ps...)
    return lexerError(msg)
}

func (r rawToken) Instruction() (opcode instruction.Opcode, err error) {
    if r.Type() != InstructionToken {
        err = makeError(r, "Token '%v' is not an instruction", r.raw)
        return
    }
    switch r.raw {
    case "!":
        return instruction.Store, nil
    case "+":
        return instruction.Add, nil
    case "-":
        return instruction.Sub, nil
    case ">R":
        return instruction.ToR, nil
    case "@":
        return instruction.Fetch, nil
    case "&":
        return instruction.And, nil
    case "DROP":
        return instruction.Drop, nil
    case "DUP":
        return instruction.Dup, nil
    case "|":
        return instruction.Or, nil
    case "OVER":
        return instruction.Over, nil
    case "R>":
        return instruction.RFrom, nil
    case "SWAP":
        return instruction.Swap, nil
    case "XOR":
        return instruction.Xor, nil
    case "JZ":
        return instruction.IfCode, nil
    case "CALL":
        return instruction.CallCode, nil
    case "RETURN":
        return instruction.Exit, nil
    default:
        // application error
        panic(makeError(r, "Unknown instruction '%v' found", r.raw))
    }
}

func (r rawToken) Integer() (val int, err error) {
    if r.Type() != IntegerToken {
        err = makeError(r, "Token '%v' is not an integer", r.raw)
        return
    }
    parsed, err := strconv.ParseInt(r.raw, 10, 32)
    if err != nil {
        return
    }
    val = int(parsed)
    return
}

func (r rawToken) Label() (name string, err error) {
    if r.Type() != LabelToken {
        err = makeError(r, "Token '%v' is not a label", r.raw)
        return
    }
    name = r.raw
    return
}

func MakeToken(raw string, file string, line int, column int) Token {
    tok := func(t TokenType) Token {
        return rawToken{raw, t, file, line, column}
    }
    switch {
    case len(raw) == 0:
        return tok(EOFToken)
    case colonRE.MatchString(raw):
        return tok(ColonToken)
    case parenOpenRE.MatchString(raw):
        return tok(ParenOpenToken)
    case parenCloseRE.MatchString(raw):
        return tok(ParenCloseToken)
    case integerRE.MatchString(raw):
        return tok(IntegerToken)
    case instructionRE.MatchString(raw):
        return tok(InstructionToken)
    case labelRE.MatchString(raw):
        return tok(LabelToken)
    default:
        panic("unrecognized token")
    }
}

func makeErrorLexer(lexer *strLexer, format string, params... interface{}) error {
    ps := make([]interface{}, len(params) + 3)
    ps[len(params) + 0] = lexer.file
    ps[len(params) + 1] = lexer.line
    ps[len(params) + 2] = lexer.column
    msg := fmt.Sprintf(format + " (%v, line %v, column %v)", ps...)
    return lexerError(msg)
}

type strLexer struct {
    string

    file string
    line int
    column int
}

func MakeStringLexer(s string) Lexer {
    return &strLexer{s, "<raw string>", 1, 1}
}

func (s *strLexer) SetFileName(f string) {
    s.file = f
}

func addPos(s *strLexer, w string) {
    newlines := newlineRE.FindAllStringIndex(w, 99)
    if len(newlines) > 0 {
        last := newlines[len(newlines) - 1]
        s.line += len(newlines)
        s.column = len(w) - last[1] + 1
    } else {
        s.column += len(w)
    }
}

func (s *strLexer) GetToken() (tok Token, err error) {
    loc := nextTokenRE.FindStringIndex(s.string)
    if loc == nil {
        if len(s.string) > 0 && !whitespaceRE.MatchString(s.string) {
            err = makeErrorLexer(s, "illegal input found: %v", s.string)
            return
        }

        if len(s.string) > 0 {
            addPos(s, s.string)
        }

        tok = MakeToken("", s.file, s.line, s.column)
        return
    }

    if loc[0] != 0 {
        prefix := s.string[:loc[0]]
        if !whitespaceRE.MatchString(prefix) {
            err = makeErrorLexer(s, "illegal input found: %v", s.string[:loc[0]])
            return
        }
        addPos(s, prefix)
    }

    t := s.string[loc[0]:loc[1]]

    tok = MakeToken(t, s.file, s.line, s.column)

    s.string = s.string[loc[1]:]
    s.column += len(t)

    return
}