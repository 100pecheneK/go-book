# part 3

## Напишите генератор константы iota для дней недели

## Напишите программу Go которая бы преобразовывала заданный массив в хеш таблицу

## Написать генератор константы iota для степеней числа четыре

## Написать собственную версию parseData.go

```go
func main(){
	var myDate string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}
	myDate = os.Args[1]
	d, err := time.Parse("02 January 2006", myDate)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}
}
```

## Написать собственную версию parseTime.go

```go
func main(){
	var myTime string
	if len(os.Args) != 2 {
		fmt.Println("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}
	myTime = os.Args[1]
	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}
}
```

## Написать версию timeData.go которая бы обрабатывала два формата даты и времени

```go
func main(){
	logs := []string{"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
                     "127.0.0.1 - - [16/Nov/2017:10:16:41 +0200]
										\"GET /CVEN HTTP/1.1\" 200 12531 \"-\"
										\"Mozilla/5.0 AppleWebKit/537.36",
										"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200]
										\"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
										"[12/Nov/2017:16:27:21 +0300]", "[12/Nov/2017:20:88:21 +0200]",
										"[12/Nov/2017:20:21 +0200]",
										}
	for _, logEntry := range logs {
		r := regexp.MustCompile(`.*[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*\].*`)
		if r.MatchString(logEntry){
			match := r.FindStringSubmatch(logEntry)
			dt, err := time.Parse("2/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			} else {
				fmt.Println("Not a valid date time format!")
			}
		} else {
			fmt.Println("Not a match!")
		}
	}
```
