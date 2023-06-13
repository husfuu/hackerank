package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversions(s string) string {
	components := strings.Split(s, ":")
	hour, _ := strconv.Atoi(components[0])
	minute, _ := strconv.Atoi(components[1])
	second, _ := strconv.Atoi(string(components[2][0:2]))
	ampm := string(components[2][2:4])

	if ampm == "AM" && hour == 12 {
		hour = 0
	} else if ampm == "PM" && hour == 12 {
		hour = 12
	} else if ampm == "PM" {
		hour += 12
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func main() {
	var input string = "05:00:00PM"
	res := timeConversions(input)

	fmt.Println(res)
}
