package RoverUtility

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var X int
var Y int
var Direction string

func Rover(location string) ([]interface{}, error) {
	roverValues := make([]interface{}, 0)

	arr := strings.Split(location, " ")
	X, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("X must be digit", err)
		return roverValues, err
	}
	Y, err1 := strconv.Atoi(arr[1])
	if err1 != nil {
		fmt.Println("Y must be digit", err1)
		return roverValues, err1
	}
	Direction = arr[2]
	Direction = strings.ToUpper(Direction)
	//expected_direction_values := ["N", "S", "E", "W"]
	expectedDirectionValues := []string{"N", "E", "W", "S"}
	res := contains(expectedDirectionValues, Direction)
	if !res {
		fmt.Println(fmt.Sprintf("direction  %s is invalid", Direction))
		return roverValues, errors.New("")
	}
	roverValues = append(roverValues, X)
	roverValues = append(roverValues, Y)
	roverValues = append(roverValues, Direction)
	return roverValues, nil

}

func SpinLeft() {

	switch Direction {
	case "N":
		Direction = "W"
		break
	case "W":
		Direction = "S"
		break
	case "S":
		Direction = "E"
		break
	case "E":
		Direction = "N"
		break
	default:
		fmt.Println("Invalid direction")
		break
	}
}

func SpinRight() {

	switch Direction {
	case "N":
		Direction = "E"
		break
	case "W":
		Direction = "N"
		break
	case "S":
		Direction = "W"
		break
	case "E":
		Direction = "S"
		break
	default:
		fmt.Println("Invalid direction")
		break
	}
}

func MoveForward() {
	switch Direction {
	case "N":
		Y += 1
		break
	case "W":
		X -= 1
		break
	case "S":
		Y -= 1
		break
	case "E":
		X += 1
		break
	default:
		fmt.Println("Invalid direction")
		break
	}
}

func Move(RoverCommand string) {
	for i := 0; i < len(RoverCommand); i++ {
		s := string(RoverCommand[i])
		switch s {
		case "L":
			SpinLeft()
			break
		case "R":
			SpinRight()
			break
		case "M":
			MoveForward()
			break
		default:
			fmt.Println("Invalid direction")
			break

		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ValidateCommand(RoverCommand string) bool {
	byteArray := []byte(RoverCommand)
	flag := true
	for _, v := range byteArray {
		if v == 'L' || v == 'R' || v == 'M' {
			flag = true
		} else {
			flag = false
			break
		}
	}
	return flag
}
