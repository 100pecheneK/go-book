package main

import (
	"fmt"
	"net"
	"os"
)

func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need one arg")
		return
	}
	input := args[1]
	IPAdress := net.ParseIP(input)
	if IPAdress == nil {
		IPs, err := lookHostname(input)
		if err == nil {
			for _, IP := range IPs {
				fmt.Println(IP)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}

	}
}
