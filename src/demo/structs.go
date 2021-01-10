package main

import "net"

type Server struct {
    IPv4,Port string
}

func main() {

    google := Server {
        IPv4: "8.8.8.8",
        Port: "8080",
    }

    net.Dial("tcp",google.IPv4)

}

