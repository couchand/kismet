package kasm

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"

    "github.com/couchand/kismet/instruction"
)

type Parameter struct {
    raw string
    isInt bool
    intVal int
    isVar bool
    name string
}

func (p Parameter) IsInt() bool {
    return p.isInt
}

func (p Parameter) IsVar() bool {
    return p.isVar
}

func (p Parameter) IntVal() int {
    if !p.IsInt() {
        msg := fmt.Sprintf("value is not an int: %v", p)
        panic(msg)
    }
    return p.intVal
}

func (p Parameter) VarName() string {
    if !p.IsVar() {
        msg := fmt.Sprintf("value is not a var: %v", p)
        panic(msg)
    }
    return p.name
}

var numberRE = regexp.MustCompile("^[0-9]+$")

func MakeParameter(p string) Parameter {
    if numberRE.MatchString(p) {
        val, err := strconv.ParseInt(p, 10, 32)
        if err != nil {
            panic(err)
        }
        return MakeIntegerParameter(int(val))
    }
    return MakeLabelParameter(strings.Trim(p, " "))
}

func MakeIntegerParameter(i int) Parameter {
    return Parameter{ raw: fmt.Sprintf("%v", i), isInt: true, intVal: i }
}

func MakeLabelParameter(n string) Parameter {
    return Parameter{ raw: n, isVar: true, name: n }
}

type InstructionPrototype struct {
    Opcode instruction.Opcode
    Payload Parameter
    Offset int
}

type LabelIndex map[string]*InstructionPrototype

func MakeLabelIndex() map[string]*InstructionPrototype {
    return make(map[string]*InstructionPrototype)
}
