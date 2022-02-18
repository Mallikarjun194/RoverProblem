package main

import (
	"awesomeProject1/InputOutputHandler"
	"awesomeProject1/app"
	"fmt"
)

// Position and Location --> (x, y, z), where z = {E, W, N, S}
// Grid Position --> (x=0, y=0, z=N), Max Grid co-ordinates (MAxX=5, MaxY=50)
// L = LeftSpin, R=RightSpin, M=MoveForward
// Input 5 Lines
//   1. Pleatue size: (5,5)
// 	 2. Array of RoverInstruction object, where a. RoverInstruction itself contains 2 lines
// 												1st line --> Rover Position
//												2nd line --> Rover Command

func main() {
	fmt.Println("WELCOME TO ROVER PROBLEM")
	fmt.Println("ENTER YOUR CHOICE")
	fmt.Println("PRESS 1 TO CONTINUE IN CLI MODE")
	fmt.Println("PRESS 2 TO CONTINUE IN RESTFUL MODE")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Please enter a valid mode")
		return
	}
	switch choice {
	case 1:
		InputOutputHandler.InputHandler()
	case 2:
		app.ApiRequests()
	default:
		fmt.Println("Invalid choice")
	}

}
