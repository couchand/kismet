package assembler

import (
    "io"
    "os"
    "fmt"
    "bytes"
    "bufio"
    "encoding/binary"

    "github.com/couchand/kismet/instruction"
)

func Assemble(instructions []instruction.T, filename string) (err error) {
    file, err := os.Create(filename)
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

    w := bufio.NewWriter(file)

    err = AssembleWriter(instructions, w)
    if err != nil {
        return
    }

    err = w.Flush()
    return
}

func intsToBytes(ints []int) ([]byte) {
    fmt.Println("byting", ints)
    convert := func(v int) (bs []byte) {
        bs = make([]byte, 4)
        binary.LittleEndian.PutUint32(bs, uint32(v))
        return
    }
    all := make([][]byte, len(ints))
    for idx, i := range ints {
        all[idx] = convert(i)
    }
    fmt.Println("got", all)
    return bytes.Join(all, []byte{})
}

func AssembleWriter(instructions []instruction.T, writer io.Writer) (err error) {
    for _, instr := range instructions {
        _, err = writer.Write(intsToBytes(instr.GetWords()))
        if err != nil {
            return
        }
    }
    return
}
