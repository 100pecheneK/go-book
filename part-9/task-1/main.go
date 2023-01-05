// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var FILES = [3]string{"text_1.txt", "text_2.txt", "text_3.txt"}
// var str = "misha"

// func main() {
// 	files := make(chan *os.File)
// 	counts := make(chan int)

// 	for _, filepath := range FILES {
// 		go add(filepath, files)
// 		go read(files, counts)
// 	}

// 	calclateCounts(counts)
// }

// func countInFile(str string, file *os.File) (sum int) {
// 	sum = 0
// 	s := bufio.NewScanner(file)
// 	s.Split(bufio.ScanWords)
// 	for s.Scan() {
// 		if s.Text() == str {
// 			sum++
// 		}
// 	}
// 	return sum
// }
// func add(filepath string, f chan *os.File) {
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("add:", filepath)
// 	f <- file
// }
// func read(f chan *os.File, in chan<- int) {
// 	file := <-f
// 	defer file.Close()
// 	count := countInFile(str, file)
// 	fmt.Printf("Count in %s: %d\n", file.Name(), count)
// 	in <- count
// }

// func calclateCounts(in chan int) {
// 	// v, ok := <-in

// 	// fmt.Println(v, ok)
// 	// v, ok = <-in
// 	// fmt.Println(v, ok)
// 	// v, ok = <-in
// 	// fmt.Println(v, ok)

// 	sum := 0

// 	for x2 := range in {
// 		sum = sum + x2
// 	}
// 	fmt.Printf("Count in files %d\n", sum)
// }
