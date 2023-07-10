package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"vpn-overrides/overrides"
)

// Напоминание: для соединения нужно создать в папке overrides файл config.json. Подробнее - в readme-файле
func main() {
	sa := syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{192, 168, 10, 2}, //Адрес хоста с TCP-сервером. TODO: попробовать то же самое, но с машиной с белым IP, котоаря не подключена к VPN
	}
	if err := overrides.Connect(128, &sa); err != nil {
		log.Panic(err)
	}

	log.Println("Connected")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter a message:")
		scanner.Scan()
		buff := scanner.Text()
		if _, err := overrides.Write(128, []byte(buff)[:len(buff)]); err != nil {
			log.Panic(err)
		}
		if buff == "/exit" {
			break
		}
		n, err := overrides.Read(128, []byte(buff))
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Server: %s\n", buff[:n])
	}

	log.Println("Exit...")
	if err := overrides.Close(128); err != nil {
		log.Panic(err)
	}
}
