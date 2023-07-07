package examples

import (
	"golang.zx2c4.com/wireguard/overrides"
	"log"
	"syscall"
)

func main() {
	sa := syscall.SockaddrInet4{
		Port: 55555,
		Addr: [4]byte{192, 168, 10, 3},
	}
	if err := overrides.Connect(128, &sa); err != nil {
		log.Panic(err)
	}
	if _, err := overrides.Write(128, []byte("Hello")); err != nil {
		log.Panic(err)
	}
}
