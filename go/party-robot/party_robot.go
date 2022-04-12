package partyrobot

import (
	"fmt"
	"math"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return ("Happy birthday " + name + "! You are now " + fmt.Sprintf("%d", age) + " years old!")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	firstLine := Welcome(name)
	secondLine := ("You have been assigned to table " + fmt.Sprintf("%03d", table) + ". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", (math.Round(distance*10)/10)) + " meters from here.")
	thirdLine := "You will be sitting next to " + neighbor + "."
	return firstLine + "\n" + secondLine + "\n" + thirdLine
}
