package main

import (
    "flag"
    "bytes"
    "encoding/binary"
    "os"
    "math"
    "net"
)

type OscResponse struct {
    Voice int32
    Vol float32
    Pitch float32
}

//The UDP listen loop
func oscListen(responseChannel chan OscResponse, sock *net.UDPConn) {
    doLog( "entering OSC packet listen loop" )
    var buf []byte
    for {
        sock.ReadFromUDP(buf)
        doLog("TODO: decode packet")
    }
}

//Write audio to the out channel
func generateNoise(responseChannel chan OscResponse) {
    doLog( "generating audio" )
    var t float64 = 0
    file := os.Stdout
    data := new(bytes.Buffer)
    sample := 0
    for {
        //TODO: check for UDP packets and send data in response
        data.Reset()
        l := (int16) ( 16000 * math.Sin(t) )
        r := (int16) ( 16000 * math.Cos(2*t) )
        binary.Write(data, binary.LittleEndian, l)
        binary.Write(data, binary.LittleEndian, r)
        file.Write(data.Bytes())
        t = t + 0.01
        if sample == 255 {
          sample = -1
          file.Sync()
        }
        sample = sample + 1
    }
}

func udpListen(name string) (sock *net.UDPConn,err error) {
    addr, erru := net.ResolveUDPAddr("udp",name)
    err = erru
    if err == nil {
        sock, err = net.ListenUDP("udp",addr)
        if err == nil {
            return sock,err
        }
    } 
    doLog("could not listen on "+name)
    return nil,err
}

func doLog(msg string) {
    os.Stderr.WriteString(msg+"\n")
    os.Stderr.Sync()
}

func main() {
    flag.Parse()
    args :=  flag.Args()
    if len(args) > 0 {
        addr := args[0]
        sock,err := udpListen(addr)
        if err == nil {
            responseChannel := make(chan OscResponse)
            go oscListen(responseChannel, sock)
            generateNoise(responseChannel)
        } else {
            doLog(err.Error())
        }
    } else {
        doLog("need an arg like 0.0.0.0:9999 for a udp port and bind addr")
    }
}
