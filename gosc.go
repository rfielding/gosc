package main

import (
    "bytes"
    "encoding/binary"
//    "fmt"
//    "bufio"
    "os"
    "math"
)

func main() {
    var t float64 = 0
    file := os.Stdout
    data := new(bytes.Buffer)
    for {
        data.Reset()
        l := (int16) ( 30000 * math.Sin(t) )
        r := (int16) ( 30000 * math.Cos(2*t) )
        binary.Write(data, binary.LittleEndian, l)
        binary.Write(data, binary.LittleEndian, r)
        file.Write(data.Bytes())
        file.Sync()
        t = t + 0.01
    }
}
