package assembler

import (
    "io"
    "os"
//    "fmt"
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

    w := bufio.NewWriter(file)

    defer func() {
        e := w.Flush()
        if err != nil && e != nil {
            panic(e)
        } else if e != nil {
            err = e
        }
        e = file.Close()
        if err != nil && e != nil {
            panic(e)
        } else if e != nil {
            err = e
        }
    }()

    err = AssembleWriter(instructions, w)
    if err != nil {
        return
    }

    err = w.Flush()
    return
}

func intsToBytes(ints []int) ([]byte) {
    //fmt.Println("byting", ints)
    convert := func(v int) (bs []byte) {
        bs = make([]byte, 4)
        binary.LittleEndian.PutUint32(bs, uint32(v))
        return
    }
    all := make([][]byte, len(ints))
    for idx, i := range ints {
        all[idx] = convert(i)
    }
    //fmt.Println("got", all)
    return bytes.Join(all, []byte{})
}

func AssembleWriter(instructions []instruction.T, writer io.Writer) (err error) {
    //fmt.Printf("Assembling %v instructions.\n", len(instructions))
    wordCount, byteCount := 0, 0
    for _, instr := range instructions {
        words := instr.GetWords()
        bs := intsToBytes(words)
        _, err = writer.Write(bs)
        if err != nil {
            return
        }
        wordCount += len(words)
        byteCount += len(bs)
    }
    //fmt.Printf("Assembled %v words and %v bytes.\n", wordCount, byteCount)
    return
}
