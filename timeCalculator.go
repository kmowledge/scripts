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

func main() {
	fmt.Println("\nHello, I can add and substract duration from a point in time, according to your wish.\n")
	for ; ; fmt.Println("\n\n\t\t* * * Once again? * * *\n") {
		time.Sleep(2 * time.Second)
		fmt.Println(`Give me a point in time in the following format 00:00:00, for example 13:22:59.
To exit the program, answer "enough".
	
	`)
		var timeIn string
		fmt.Scanln(&timeIn)
		if timeIn == "enough" {
			return
		}
		formatCheckTime(&timeIn)
		time.Sleep(1 * time.Second)
		fmt.Println(`
Good. Now, give me an amount of time.`)
		time.Sleep(500 * time.Millisecond)
		fmt.Println(`Use format like this: 1h20m39s`)
		time.Sleep(500 * time.Millisecond)
		fmt.Println(`By default I will add it. To substract add "-" before.`)
		time.Sleep(500 * time.Millisecond)
		fmt.Println(`Examples: 60s, -48h1s, 000h999999m, -2h20m13s

`)
		var durIn string
		fmt.Scanln(&durIn)
		if durIn == "enough" {
			return
		}
		// formatCheckDuration(&durIn)
		timeInType, _ := time.Parse(time.TimeOnly, timeIn)
		// fmt.Println("Read value: timeInType") --- for testing
		durInType, _ := time.ParseDuration(durIn)
		var timeOutType time.Time
		timeOutType = timeInType.Add(durInType)
		timeOut := timeOutType.Format(time.TimeOnly)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("\n%v", timeIn)
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" + ")
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("(%v)", durIn)
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" = ")
		time.Sleep(1 * time.Second)
		fmt.Printf("%v\n", timeOut)
		time.Sleep(2 * time.Second)
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
