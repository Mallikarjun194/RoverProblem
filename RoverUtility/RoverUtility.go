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

func RoverValidation(location string) ([]interface{}, error) {
	roverValues := make([]interface{}, 0)

	arr := strings.Split(location, " ")
	X, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("X must be digit", err)
		return roverValues, err
	}
	Y, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("Y must be digit", err)
		return roverValues, err
	}
	Direction = arr[2]
	Direction = strings.ToUpper(Direction)
	//expected_direction_values := ["N", "S", "E", "W"]
	expectedDirectionValues := []string{"N", "E", "W", "S"}
	res := contains(expectedDirectionValues, Direction)
	if !res {
		//fmt.Println(fmt.Sprintf("direction  %s is invalid", Direction))
		return roverValues, errors.New("direction is invalid")
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
	case "W":
		Direction = "S"
	case "S":
		Direction = "E"
	case "E":
		Direction = "N"
	}
}

func SpinRight() {

	switch Direction {
	case "N":
		Direction = "E"
	case "W":
		Direction = "N"
	case "S":
		Direction = "W"
	case "E":
		Direction = "S"
	}
}

func MoveForward() {
	switch Direction {
	case "N":
		Y += 1
	case "W":
		X -= 1
	case "S":
		Y -= 1
	case "E":
		X += 1
	}
}

func Move(RoverCommand string) {
	for i := 0; i < len(RoverCommand); i++ {
		s := string(RoverCommand[i])
		switch s {
		case "L":
			SpinLeft()
		case "R":
			SpinRight()
		case "M":
			MoveForward()

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
