package main

import (
    "net"
    "fmt"
)


func startServer() {
    server, err := net.Listen("tcp",":6060")
    if(err == nil) {
        for {
           server.Accept()
        }
    }
}


func main() {
    fmt.Println("Start server:")
    go startServer()
    fmt.Println("After routine")
    fmt.Println("DONE!")
}
