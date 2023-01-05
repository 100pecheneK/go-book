package main

import (
	"fmt"
	"regexp"
	"time"
)

type Matchers map[*regexp.Regexp]string

func main() {
	logs := []string{"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200] 325504",
		"127.0.0.1 - - [16/Nov/2017:10:16:41 +0200]\"GET /CVEN HTTP/1.1\" 200 12531 \"-\"\"Mozilla/5.0 AppleWebKit/537.36",
		"127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200]\"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
		"[12 Nov 2017:16:27:21 +0300]",
	}

	r := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)
	r2 := regexp.MustCompile(`.*\[(\d\d\s\w+\s\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)

	matchers := make(Matchers)
	matchers[r] = "2/Jan/2006:15:04:05 -0700"
	matchers[r2] = "2 Jan 2006:15:04:05 -0700"

	parse(logs, matchers)
}

func parse(dates []string, matchers Matchers) {
	for _, date := range dates {
		var notAMatch bool
		for reg, layout := range matchers {
			if reg.MatchString(date) {
				notAMatch = false
				match := reg.FindStringSubmatch(date)
				dt, err := time.Parse(layout, match[1])
				if err != nil {
					fmt.Println("Not a valid date time format!")
					break
				}
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
				break
			} else {
				notAMatch = true
			}
		}
		if notAMatch {
			fmt.Println("Not a match")
			notAMatch = false
		}
	}
}
