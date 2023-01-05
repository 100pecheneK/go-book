# Part 1

## 1

Напишите программу которая вычесляет сумму всех аргументов командной строки которые являются действительными числами

## 2

Напишите программу вычисляющую среднее значение всех чисел с плавающей запятой переданных программе в качестве аргументов командной строки

## 3

Напишите программу которая считыает целые числа до тех пор пока не встретит во входных данных слово END

## 4

Изменить программу, чтобы данные журнала записывались одновременно в два файла журнала

```go
package main
import  (
	"fmt"
	"log"
	"os"
)

var LOGFILE = './tmp/go1.log'

func main(){
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
```
