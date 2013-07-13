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

func oscListen(responseChannel chan<-OscResponse, sock *net.UDPConn) {
    //var udpBuf [1024]byte
    //TODO: do something with incoming UDP packets from controller
}

func generateNoise(responseChannel chan<-OscResponse) {
    var t float64 = 0
    file := os.Stdout
    data := new(bytes.Buffer)
    for {
        //TODO: check for UDP packets and send data in response
        data.Reset()
        l := (int16) ( 16000 * math.Sin(t) )
        r := (int16) ( 16000 * math.Cos(2*t) )
        binary.Write(data, binary.LittleEndian, l)
        binary.Write(data, binary.LittleEndian, r)
        file.Write(data.Bytes())
        file.Sync()
        t = t + 0.01
    }
}

func main() {
    flag.Parse()
    args :=  flag.Args()

    //Start listening on our UDP port and make noise in response
    addr, err := net.ResolveUDPAddr("udp",args[0])
    if err == nil {
        sock, erru := net.ListenUDP("udp",addr)
        if erru == nil {
            responseChannel := make(chan OscResponse)
            go oscListen(responseChannel,sock)
            generateNoise(responseChannel)
        } else {
            os.Stderr.WriteString("could not listen")
        }
    } else {
      os.Stderr.WriteString("could not open address")
    }
}
