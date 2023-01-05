package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s.go logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}
	ips := []string{}
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
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			ips = append(ips, ip)
		}
	}
	printSortedUniqIps(ips)
}

func findIPv4(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func printSortedUniqIps(ips []string) {
	sortedAndUniq := sortByValue(uniq(ips))
	for ip, count := range sortedAndUniq {
		fmt.Println(count, ip)
	}
}

func sortByValue[T string | int](m map[T]int) map[T]int {
	keys := make([]T, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	sorted := make(map[T]int)

	for _, key := range keys {
		sorted[key] = m[key]
	}
	return sorted
}

func uniq[T string | int](arr []T) map[T]int {
	uniq := make(map[T]int)

	for _, a := range arr {
		_, exist := uniq[a]
		if exist {
			uniq[a]++
		} else {
			uniq[a] = 1
		}
	}
	return uniq
}
