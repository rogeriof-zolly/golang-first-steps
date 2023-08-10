package main

import "fmt"

func weekDay(weekDayNumber int) string {

	switch weekDayNumber {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid input"
	}
}

func weekDays2(weekDayNumber int) string {

	// In go, it is not necessary to add a break statement after each case
	// Instead, there's a 'fallthrough' statement to go to the next case and run it
	switch {
	case weekDayNumber == 1:
		return "Sunday"
	case weekDayNumber == 2:
		return "Monday"
	case weekDayNumber == 3:
		return "Tuesday"
	case weekDayNumber == 4:
		return "Wednesday"
	case weekDayNumber == 5:
		return "Thursday"
	case weekDayNumber == 6:
		return "Friday"
	case weekDayNumber == 7:
		return "Saturday"
	default:
		return "Invalid input"
	}
}

func main() {

	weekDay1 := weekDay(1)

	fmt.Println(weekDay1)

	weekDay2 := weekDays2(4)

	fmt.Println("--------------------------")
	fmt.Println(weekDay2)
}
