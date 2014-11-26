package loader

import (
    "io"
    "os"
//    "fmt"
//    "bytes"
    "bufio"
    "encoding/binary"

    "github.com/couchand/kismet/instruction"
)

func Load(filename string) (instructions []instruction.T, err error) {
    file, err := os.Open(filename)
    if err != nil {
        return
    }

    defer func() {
        e := file.Close()
        if err != nil && e != nil {
            panic(e)
        } else if e != nil {
            err = e
        }
    }()

    r := bufio.NewReader(file)

    return LoadReader(r)
}

func convert(bs []byte) int {
    return int(binary.LittleEndian.Uint32(bs))
}

func bytesToInstructions(bs []byte) (instructions []instruction.T) {
    if len(bs) < 4 {
        if len(bs) != 0 {
            //fmt.Printf("remaining: 0x%x\n", bs)
            panic("frame shift error!")
        }

        return []instruction.T{}
    }
    instructions = make([]instruction.T, 0, len(bs))

    for len(bs) >= 4 {
        code := instruction.Opcode(convert(bs[0:4]))
        var instr instruction.T
        switch code {
        case instruction.IfCode, instruction.CallCode, instruction.LitCode:
            instr = instruction.DoubleWide{ code, convert(bs[4:8]) }
            bs = bs[8:]
        default:
            instr = code
            bs = bs[4:]
        }
        instructions = append(instructions, instr)
    }

    if len(bs) != 0 {
        //fmt.Printf("remaining: 0x%x\n", bs)
        panic("frame shift error!")
    }
    return
}

func LoadReader(reader io.Reader) (instructions []instruction.T, err error) {
    bytes := make([]byte, 1024)
    n, err := reader.Read(bytes)
    if err != nil {
        return
    }

    //fmt.Printf("Read %v bytes.\n", n)

    instructions = bytesToInstructions(bytes[:n])
    return
}
