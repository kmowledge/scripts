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

/*
>>> Bug report: Strange program behaviours in response to atypical user inputs.

1. First input (point in time):
1.1. General observations

1.1.1. If user adds a word before writing a point in time in correct format, it disturbs the "hour" figure and resets it to 0.
1.1.2. If the word is added after the correct time input, it skips the second input and prints the result - same as if it was sent blank.


1.2. Examples of bug behaviour

1.2.1. This input triggers error message 2 times. Why.
kurczak 1:30:00
Error: stick to the given 24h time format ---> 00:00:00
Error: stick to the given 24h time format ---> 00:00:00

1.2.2. a) This input breaks the loop, allowing to move to the 2nd input. Why.
	 b) But reads only 1 char from hour.
kurczak 11:30:00
Error: stick to the given 24h time format ---> 00:00:00
Read value: 0000-01-01 01:30:00 +0000 UTC

*/
