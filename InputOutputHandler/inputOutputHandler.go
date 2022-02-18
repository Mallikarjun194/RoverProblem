package InputOutputHandler

import (
	"awesomeProject1/RoverUtility"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputHandler() {
	fmt.Println("WELCOME TO CLI MODE TO FIND THE ROVER POSITION")
	var MaxX, MaxY int
	fmt.Println("Enter max position seperated by space:")
	_, err := fmt.Scanln(&MaxX, &MaxY)
	if err != nil {
		fmt.Println("Invalid inputs for MaxX and MaxY")
		return
	}
	outputs := make([]string, 0)
	for {
		var location string
		fmt.Println("Enter Rover position seperated by space:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			location = scanner.Text()
		}

		roverValues, err := RoverUtility.RoverValidation(location)
		if err != nil {
			outputs = append(outputs, fmt.Sprintf("Invalid location/co-ordinate %s", roverValues))
			if ExitCheck() {
				continue
			} else {
				break
			}
		} else {
			RoverUtility.X, RoverUtility.Y, RoverUtility.Direction =
				roverValues[0].(int), roverValues[1].(int), roverValues[2].(string)
		}

		var RoverCommand string
		fmt.Println("Command for rover:")
		if scanner.Scan() {
			RoverCommand = scanner.Text()
		}
		RoverCommand = strings.ToUpper(RoverCommand)
		res := RoverUtility.ValidateCommand(RoverCommand)
		if !res {
			outputs = append(outputs, fmt.Sprintf("Invalid command %s", RoverCommand))
			if ExitCheck() {
				continue
			} else {
				break
			}
		} else {
			//fmt.Println(RoverCommand)
			RoverUtility.Move(RoverCommand)
		}

		if RoverUtility.X > MaxX || RoverUtility.Y > MaxY {
			outputs = append(outputs, fmt.Sprintf("Invalid input, Exceeding plateau %d %d", RoverUtility.X,
				RoverUtility.Y))
			if ExitCheck() {
				continue
			} else {
				break
			}
		} else {

			outputs = append(outputs, fmt.Sprintf("%d %d %s", RoverUtility.X, RoverUtility.Y,
				RoverUtility.Direction))
		}
		if ExitCheck() {
			continue
		} else {
			break
		}

	}
	//fmt.Println(X, Y, direction)
	for _, v := range outputs {
		fmt.Println("Rover Current Position:")
		fmt.Println(v)
	}
}

func ExitCheck() bool {
	var exit string
	fmt.Println("To Continue Inputs Y/N")
	fmt.Scanln(&exit)
	if exit == "Y" || exit == "y" {
		return true
	} else {
		return false
	}
}
