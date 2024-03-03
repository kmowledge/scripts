package main

import (
	"fmt"
	"regexp"
	"time"
)

func formatCheckTime(a *string) {
	for {
		if b, _ := regexp.MatchString(`^(([01]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])$`, *a); b {
			break
		}
		fmt.Println("Error: stick to the given 24h time format ---> 00:00:00")
		fmt.Scanln(a) //If you scan to &address of a pointer, you get infinite loop of the print above. Why.
		if *a == "enough" {
			return
		}
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

func main() {
	fmt.Println("Hello, I can add and substract duration from a point in time, according to your wish.")
	for {
		fmt.Println(`Give me a point in time in the following format 00:00:00.
	To exit the program, answer "enough".`)
		var timeIn string
		fmt.Scanln(&timeIn)
		if timeIn == "enough" {
			return
		}
		formatCheckTime(&timeIn)
		fmt.Println(`	Now, give me an amount of time.
	By default I will add it. To substract add "-" before.
	Use format like this: 1h20m39s
	Accepted examples: 60s, -48h1s, 000h999999m, -2h20m13s`)
		var durIn string
		fmt.Scanln(&durIn)
		if durIn == "enough" {
			return
		}
		// formatCheckDuration(&durIn)
		timeInType, _ := time.Parse(time.TimeOnly, timeIn)
		fmt.Println("Read value: timeInType")
		durInType, _ := time.ParseDuration(durIn)
		var timeOutType time.Time
		timeOutType = timeInType.Add(durInType)
		timeOut := timeOutType.Format(time.TimeOnly)
		fmt.Println("Result: ", timeOut)
	}
}
