package models

// Struct for rover i/o

type RoverProblem struct {
	Id            string `json:"id" validate:"required"`
	MaxX          int    `json:"max_x" validate:"required"`
	MaxY          int    `json:"max_y" validate:"required"`
	Location      string `json:"location" validate:"required"`
	Command       string `json:"command" validate:"required"`
	RoverPosition string `json:"roverPosition"`
}

// Struct to handle errors

type ErrorHandler struct {
	ErrorMsg string `json:"error_msg"`
}
