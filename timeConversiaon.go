
package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	// Write your code here
	var time string

	if string(s[8]) == "A" {
		if s[:2] == "12" {
			time = "00" + s[2:8]
		} else {
			time = s[:8]
		}
	} else {
		if s[:2] == "12" {
			time = s[:8]
		} else {
			converted_int, _ := strconv.Atoi(s[:2])
			converted_str := strconv.Itoa(converted_int + 12)
			time = converted_str + s[2:8]
		}
	}

	return time
}

func main() {
	var s string

	fmt.Scanf("%s\n", &s)
	fmt.Println(timeConversion(s))
}