package main

// loads a.sto and executes

import (
    "fmt"
    "flag"
    "io/ioutil"
    "strconv"
    "regexp"

    "github.com/couchand/kismet/asm"
    "github.com/couchand/kismet/kasm"
    "github.com/couchand/kismet/assembler"
    "github.com/couchand/kismet/disassembler"
    "github.com/couchand/kismet/instruction"
    "github.com/couchand/kismet/loader"
    "github.com/couchand/kismet/machine"
)

var rangeRE *regexp.Regexp = regexp.MustCompile("^([0-9a-f]+)-([0-9a-f]+)$")

func runFile(f string, debug, quiet bool) {
    p, err := loader.Load(f)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    m := machine.Make10KMachine(p)

    watches := make([]int, 0)

    for {
        if !quiet {
            fmt.Printf("%v\n", m)
            for _, addr := range watches {
                val := m.MemoryAt(addr)
                fmt.Printf("Watch [0x%.8x]: 0x%.8x\n", addr, val)
            }
        }

        if debug {
            var cmd string
            var param string
            fmt.Scanln(&cmd, &param)

            if len(cmd) > 0 {
                switch cmd[0] {
                case 'R', 'r':
                    debug = false
                case 'W', 'w':
                    var addrstr string
                    if len(param) == 0 {
                        if len(cmd[1:]) == 0 {
                            break
                        }

                        addrstr = cmd[1:]
                    } else {
                        addrstr = param
                    }
                    if rangeRE.MatchString(addrstr) {
                        groups := rangeRE.FindStringSubmatch(addrstr)
                        start, err := strconv.ParseInt(groups[1], 16, 32)
                        if err != nil {
                            fmt.Println("Error:", err)
                        } else {
                            end, err := strconv.ParseInt(groups[2], 16, 32)
                            if err != nil {
                                fmt.Println("Error:", err)
                            } else {
                                for i := int(start); i <= int(end); i += 1 {
                                    watches = append(watches, i)
                                }
                            }
                        }
                    } else {
                        addr, err := strconv.ParseInt(addrstr, 16, 32)
                        if err != nil {
                            fmt.Println("Error:", err)
                        } else {
                            watches = append(watches, int(addr))
                        }
                    }
                }
            }
        }

        done := m.Step()
        if done {
            if m.State() == machine.HaltState {
                break
            }
            debug = true
            quiet = false
        }
    }

    if quiet {
        fmt.Printf("%v\n", m)
    }
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
