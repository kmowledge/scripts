package main

import (
	"fmt"
	"regexp"
	"time"
)

func formatCheckTime(a *string) {
	for {
		if b, _ := regexp.MatchString(`[0-23]\:[0-5][0-9]\:[0-5][0-9]`, *a); b {
			break
		}
		fmt.Println("stick to the (24h) time format ---> 00:00:00")
		fmt.Scanln(a) //If you scan to &address of a pointer, you get infinite loop of the print above. Why?
	}
}

func formatCheckDuration(a *string) {
	for {
		if b, _ := regexp.MatchString(`[0-23][h][0-5][0-9][m][0-5][0-9][s]`, *a); b {
			break
		}
		fmt.Println("stick to the duration format ---> 00h00m00s")
		fmt.Scanln(a) //If you scan to &address of a pointer, you get infinite loop of the print above. Why?
	}
}

// func 30minEarlier(a)

func main() {
	fmt.Println(`	Hello, gimme time in the following format 00:00:00
	and I will tell you what time was 30 minutes ago.\n
	To exit the program, answer "enough".`)
	// for/if+return jeśli enough to pomiń wszystko, zakończ program
	var timeIn string
	fmt.Scanln(&timeIn)
	formatCheckTime(&timeIn)
	// timeSlice := s.Split(timeIn, ":")
	// timeAmount := fmt.Sprint(timeSlice[0], "h", timeSlice[1], "m", timeSlice[2], "s")
	timeOut, _ := time.Parse("TimeOnly", timeIn) //TimeOnly != Longform(layout)
	fmt.Println(`	Now, tell me what duration to substract from this.
	Use format like this: 1h20m39s`) //strings. parse +/- adding/substr option
	var durIn string
	fmt.Scanln(&durIn)
	formatCheckDuration(&durIn)

	// timeFormatted := time.parseString(timeIn)
	// timeOut := timeFormatted.Sub(time.Minute * 30)
	fmt.Println("You said: ", timeOut)
}
