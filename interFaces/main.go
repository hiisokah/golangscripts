package main

import (
	"fmt"
	"net"
)

func main() {
	faces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("(+) Error %v", err)
	}

	for _, iface := range faces {

		addr, err := iface.Addrs()
		if err != nil {
			fmt.Println("(+) err", err)
		}

		for _, adr := range addr {
			fmt.Println(iface.Flags)
			fmt.Printf("\n\tInterface: %s | ADDR: %s \n\t | MAC: %s \n\t ", iface.Name, adr, iface.HardwareAddr)

		}
	}
}
