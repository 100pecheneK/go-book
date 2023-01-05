package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

type Connect struct {
	ip      string
	errCode string
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s.go logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}
	ips := []Connect{}
	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error open file %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		r := bufio.NewReader(f)

		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s\n", err)
				break
			}
			ip := findIPv4(line)
			errorCode := findErrorCode(line)
			conn := Connect{
				ip:      ip,
				errCode: errorCode,
			}

			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			ips = append(ips, conn)
		}
	}
	printError404Ips(ips)
}

func findErrorCode(input string) string {
	errors := "(Error 404)|(OK 200)"
	matchMe := regexp.MustCompile(errors)
	return matchMe.FindString(input)
}

func findIPv4(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func printError404Ips(ips []Connect) {
	for _, ip := range ips {
		if ip.errCode == "Error 404" {
			fmt.Printf("Not found in ip: %s\n", ip.ip)
		}
	}
}
