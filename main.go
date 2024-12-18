package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func main() {
	var name string
	var userInput int
	fmt.Println("Welcome to the Master Mind Game:")
	fmt.Println("Enter the Name")
	fmt.Scan(&name)
	fmt.Println("Start guessing")
	// do the logic fo putting it into the data base

	numberMaps := GenerateRandomNumber()
	var count = 0
	totalTime := time.Now()
	for true {
		startTime := time.Now()
		var left = 0
		var right = 0
		fmt.Scan(&userInput)
		clearLine()

		userInput1 := userInput
		position4 := userInput % 10
		userInput = userInput / 10
		position3 := userInput % 10
		userInput = userInput / 10
		position2 := userInput % 10
		userInput = userInput / 10
		position1 := userInput

		if position1 == position2 || position2 == position3 || position3 == position4 || position1 == position4 || position2 == position4 || position1 == position3 {
			fmt.Println("The digts are repeated. Please try again.")
			continue
		}
		value4, exist4 := numberMaps[position4]
		if exist4 {

			if value4 == 4 {
				right++
			} else {
				left++
			}
		}

		value3, exist3 := numberMaps[position3]
		if exist3 {

			if value3 == 3 {
				right++
			} else {
				left++
			}
		}

		value2, exist2 := numberMaps[position2]
		if exist2 {

			if value2 == 2 {
				right++
			} else {
				left++
			}
		}

		value1, exist1 := numberMaps[position1]
		if exist1 {

			if value1 == 1 {
				right++
			} else {
				left++
			}
		}
		if right == 4 {
			fmt.Print(userInput1)

			fmt.Print(color.BlueString(fmt.Sprintf("%v", left)))
			fmt.Print(":")
			fmt.Print(color.GreenString(fmt.Sprintf("%v", right)))
			totalDuration := time.Since(totalTime)
			fmt.Printf("\t Total Duration: %v\n", totalDuration)

			break
		} else {
			count++
			fmt.Print(userInput1)
			fmt.Print("\t ")
			fmt.Print(color.BlueString(fmt.Sprintf("%v", left)))
			fmt.Print(":")
			fmt.Print(color.GreenString(fmt.Sprintf("%v", right)))
			duration := time.Since(startTime)
			fmt.Printf("\t #%d chances taken |  Duration: %v\n", count, duration)
			left = 0
			right = 0
			fmt.Println("Take another guess")
		}
	}

}

func GenerateRandomNumber() map[int]int {

	numberMaps := make(map[int]int)
	max := 10
	min := 0
	position := 1

	for len(numberMaps) < 4 {

		randomNumber := rand.Intn(max-min) + min

		if position == 1 && randomNumber == 0 {
			continue

		}
		if _, exists := numberMaps[randomNumber]; exists {
			continue
		} else {
			numberMaps[randomNumber] = position
			position++

		}

	}
	return numberMaps
}

func clearLine() {
	fmt.Print("\033[K")
}
