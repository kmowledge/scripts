package main

import (
	"fmt"
	"regexp"
	"time"
)

func formatCheckTime(a *string) {
	for {
		if b, _ := regexp.MatchString(`^(([01]?[0-9]|2[0-3]):[0-5]?[0-9]:[0-5]?[0-9])$`, *a); b {
			break
		}
		fmt.Println("stick to the (24h) time format ---> 00:00:00")
		fmt.Scanln(a) //If you scan to &address of a pointer, you get infinite loop of the print above. Why?
	}
}

// func formatCheckDuration(a *string) {
// 	for {
// 		if b, _ := regexp.MatchString(`[0-9]?[0-9]?[0-9][h][0-5][0-9][m][0-5][0-9][s]`, *a); b {
// 			break
// 		}
// 		fmt.Println("stick to the duration format ---> 00h00m00s")
// 		fmt.Scanln(a) //If you scan to &address of a pointer, you get infinite loop of the print above. Why?
// 	}
// }

// func 30minEarlier(a)

func main() {
	fmt.Println(`	Hello, gimme time in the following format 00:00:00
	To exit the program, answer "enough".`)
	var timeIn string
	fmt.Scanln(&timeIn)
	if timeIn == "enough" {
		return
	}
	formatCheckTime(&timeIn)
	// timeSlice := s.Split(timeIn, ":")
	// timeAmount := fmt.Sprint(timeSlice[0], "h", timeSlice[1], "m", timeSlice[2], "s")
	fmt.Println(`	Now, tell me what duration to substract from this.
	Use format like this: 1h20m39s
	Accepted examples: 60s, 48h1s, 9999m`) //strings. parse +/- adding/substr option
	var durIn string
	fmt.Scanln(&durIn)
	if durIn == "enough" {
		return
	}
	// formatCheckDuration(&durIn)
	timeInType, _ := time.Parse(time.TimeOnly, timeIn) //TimeOnly != Longform(layout)
	durInType, _ := time.ParseDuration(durIn)
	var timeOutType time.Time
	timeOutType = timeInType.Add(-durInType)
	timeOut := timeOutType.Format(time.TimeOnly)
	// timeFormatted := time.parseString(timeIn)
	// timeOut := timeFormatted.Sub(time.Minute * 30)
	fmt.Println("Result: ", timeOut)
}
