package kasm

import (
//    "fmt"

    "github.com/couchand/kismet/instruction"
)

type parser struct {
    lexer Lexer
    labels LabelIndex
}

func MakeParser(l Lexer) parser {
    return parser{l, MakeLabelIndex()}
}

func (p *parser) Parse() (instructions []*InstructionPrototype, err error) {
    var tok Token
    is := make([]*InstructionPrototype, 0)
    for {
        tok, err = p.lexer.GetToken()
        if err != nil {
            return
        }
        switch tok.Type() {
        case EOFToken:
            instructions = is
            return

        default:
            var inst *InstructionPrototype
            inst, err = p.parseInstruction(tok)
            if err != nil {
                return
            }
            is = append(is, inst)
        }
    }
}

func (p *parser) parseInstruction(t Token) (i *InstructionPrototype, err error) {
    var opcode instruction.Opcode
    switch t.Type() {
    case InstructionToken:
        opcode, err = t.Instruction()
        if err != nil {
            return
        }

        i = &InstructionPrototype{Opcode: opcode}

        if !opcode.IsDoubleWide() {
            return
        }

        var lparen, param, rparen Token
        lparen, err = p.lexer.GetToken()
        if err != nil {
            return
        }
        if lparen.Type() != ParenOpenToken {
            err = makeError(lparen, "Expected instruction parameter")
            return
        }
        param, err = p.lexer.GetToken()
        if err != nil {
            return
        }
        if param.Type() != IntegerToken && param.Type() != LabelToken {
            err = makeError(param, "Expected instruction parameter")
            return
        }
        rparen, err = p.lexer.GetToken()
        if err != nil {
            return
        }
        if rparen.Type() != ParenCloseToken {
            err = makeError(rparen, "Expected instruction parameter")
            return
        }

        if param.Type() == IntegerToken {
            var val int
            val, err = param.Integer()
            if err != nil {
                return
            }
            i.Payload = MakeIntegerParameter(val)
        }
        if param.Type() == LabelToken {
            var name string
            name, err = param.Label()
            if err != nil {
                return
            }
            i.Payload = MakeLabelParameter(name)
        }
        return

    case IntegerToken:
        var val int
        val, err = t.Integer()
        if err != nil {
            return
        }
        i = &InstructionPrototype{
            Opcode: instruction.LitCode,
            Payload: MakeIntegerParameter(val),
        }
        return

    case LabelToken:
        var n string
        n, err = t.Label()
        if err != nil {
            return
        }
        var colon, child Token
        colon, err = p.lexer.GetToken()
        if err != nil {
            return
        }
        if colon.Type() != ColonToken {
            err = makeError(colon, "Expected a colon following identifier '%v', saw '%v'", n, colon.Raw())
            return
        }
        child, err = p.lexer.GetToken()
        if err != nil {
            return
        }
        i, err = p.parseInstruction(child)

        _, exists := p.labels[n]
        if exists {
            makeError(t, "Label '%v' already defined", n)
        }

        p.labels[n] = i
        return

    default:
        err = makeError(t, "Unexpected token type '%v'", t.Type())
        return
    }
}
