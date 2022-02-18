package app

import (
	"awesomeProject1/RoverUtility"
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var InputFile = "RoverPosition.json"

func writeFile(b []byte) bool {
	err3 := ioutil.WriteFile(InputFile, b, 0644)
	if err3 != nil {
		//fmt.Println("Error occur when writing to a file: %s", err3)
		return false
	}
	return true
}

var Error models.ErrorHandler
var RoverProb []models.RoverProblem

func RoverPosition(w http.ResponseWriter, r *http.Request) {
	//
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	var RoverPro models.RoverProblem
	error := json.Unmarshal(reqBody, &RoverPro)
	if error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		Error.ErrorMsg = "Error occur when marshalling request data"
		json.NewEncoder(w).Encode(&Error)
	} else {
		rand.Seed(time.Now().UnixNano())

		//RoverPro.Id, _ = h.Encode([]int{int(now.Unix())})
		RoverPro.Id = strconv.Itoa(1000 + rand.Intn(10000-1000))
		MaxX := RoverPro.MaxX
		MaxY := RoverPro.MaxY
		location := RoverPro.Location
		command := RoverPro.Command
		print(MaxX, "\n", MaxY, "\n", location, "\n", command, "\n")
		roverValues, err := RoverUtility.Rover(location)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			Error.ErrorMsg = "Invalid json payload error, Invalid Location " + location
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Error)
			return
		}
		RoverUtility.X, RoverUtility.Y, RoverUtility.Direction = roverValues[0].(int), roverValues[1].(int), roverValues[2].(string)
		command = strings.ToUpper(command)
		res := RoverUtility.ValidateCommand(command)
		if !res {
			fmt.Println("Invalid command", command)
			w.Header().Add("Content-Type", "application/json")
			Error.ErrorMsg = "Invalid json payload error, Invalid command " + command
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Error)

			return
		}
		//fmt.Println(RoverCommand)
		RoverUtility.Move(command)
		if RoverUtility.X > MaxX || RoverUtility.Y > MaxY {
			fmt.Println("Positions got exceeded ", RoverUtility.X, RoverUtility.Y)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			Error.ErrorMsg = fmt.Sprintf("Positions got exceeded %d %d", RoverUtility.X, RoverUtility.Y)
			json.NewEncoder(w).Encode(&Error)
			return

		}
		FinalX := strconv.Itoa(RoverUtility.X)
		FinalY := strconv.Itoa(RoverUtility.Y)
		fmt.Printf("%T %T %T", FinalX, FinalY, RoverUtility.Direction)

		//fmt.Sprintf("%s %s %s", FinalX, FinalY, direction)
		result := fmt.Sprintf("%s %s %s", FinalX, FinalY, RoverUtility.Direction)
		//res = X+" "+Y+" "+direction
		RoverPro.RoverPosition = result
		RoverProb = append(RoverProb, RoverPro)
		b, err := json.MarshalIndent(RoverProb, "", "\t")
		if err != nil {
			fmt.Printf("Error occur when marshalling request data: %v", err)
		}
		if writeFile(b) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Printf("Postion of Rover:%v\n", RoverPro.RoverPosition)
			json.NewEncoder(w).Encode(RoverPro)
		}

	}
}

func getAllRoverPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(RoverProb)
}

func getRoverPositionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, rover := range RoverProb {
		if rover.Id == key {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(rover)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	Error.ErrorMsg = "Requested ID not found"
	json.NewEncoder(w).Encode(&Error)

}

func ReadFile() {
	if _, err := os.Stat(InputFile); err == nil {
		content, err1 := ioutil.ReadFile(InputFile)
		if err1 != nil {
			fmt.Println("Error while opening file", err1)
		}
		if len(content) > 0 {
			err2 := json.Unmarshal(content, &RoverProb)
			if err2 != nil {
				fmt.Println("Error while unmarshalling the content", err2)
			}
		}
	}
}

func ApiRequests() {
	//
	fmt.Println("WELCOME TO API MODE TO FIND THE ROVER POSITION")
	ReadFile()
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", RoverPosition).Methods("POST")
	route.HandleFunc("/", getAllRoverPosition).Methods("GET")
	route.HandleFunc("/{id}", getRoverPositionById).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", route))
}
