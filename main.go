package main

// loads a.sto and executes

import (
    "fmt"
    "flag"
    "io/ioutil"

    "github.com/couchand/kismet/asm"
    "github.com/couchand/kismet/kasm"
    "github.com/couchand/kismet/assembler"
    "github.com/couchand/kismet/disassembler"
    "github.com/couchand/kismet/instruction"
    "github.com/couchand/kismet/loader"
    "github.com/couchand/kismet/machine"
)

func runFile(f string, debug, quiet bool) {
    p, err := loader.Load(f)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    m := machine.Make10KMachine(p)

    for {
        if !quiet {
            fmt.Printf("%v\n", m)
        }

        done := m.Step()
        if done {
            break
        }

        if debug {
            var cmd string
            fmt.Scanln(&cmd)
        }
    }

    fmt.Printf("%v\n", m)
    fmt.Printf("Done.\n\n")
}

func assembleFile(f, input, output string, useAsm bool) {
    var instructions []instruction.T
    if useAsm {
        instructions = asm.ParseString(f)
    } else {
        instructions = kasm.ParseFile(f, input)
    }
    assembler.Assemble(instructions, output)
}

func disassembleFile(f, output string) {
    p, err := loader.Load(f)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    str := disassembler.Disassemble(p)
    err = ioutil.WriteFile(output, []byte(str), 0644)
    if err != nil {
        fmt.Println( "Error:", err)
        return
    }
}

func main() {
    run, assemble, disassemble, output := "", "", "", ""
    useAsm, debug, quiet := false, false, false

    flag.StringVar(&run, "run", "", "run the object file")
    flag.StringVar(&assemble, "assemble", "", "assemble the input file")
    flag.StringVar(&disassemble, "disassemble", "", "disassemble the input file")
    flag.StringVar(&output, "output", "a.out", "output file name")
    flag.BoolVar(&useAsm, "asm", false, "use basic asm assembler rather than kasm")
    flag.BoolVar(&debug, "debug", false, "debug object file")
    flag.BoolVar(&quiet, "quiet", false, "quiet output")

    flag.Parse()

    if len(run) != 0 && len(assemble) != 0 && len(disassemble) != 0 {
        fmt.Println("Please provide only one command.")
        return
    }

    if len(run) != 0 {
        runFile(run, debug, quiet)
        return
    }
    if len(assemble) != 0 {
        program, err := ioutil.ReadFile(assemble)
        if err != nil {
            panic(err)
        }

        assembleFile(string(program), assemble, output, useAsm)
        return
    }
    if len(disassemble) != 0 {
        disassembleFile(disassemble, output)
        return
    }
    fmt.Println("Must provide a command.")
}
