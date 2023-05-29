package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var nextDirection = getNewDirection()
var moveValues = getMoveValues()

func getNewDirection() map[string]string {
	var nextDirection = make(map[string]string)

	key := fmt.Sprintf("%s-%s", "N", "L")
	nextDirection[key] = "W"

	key = fmt.Sprintf("%s-%s", "W", "L")
	nextDirection[key] = "S"

	key = fmt.Sprintf("%s-%s", "S", "L")
	nextDirection[key] = "E"

	key = fmt.Sprintf("%s-%s", "E", "L")
	nextDirection[key] = "N"

	key = fmt.Sprintf("%s-%s", "N", "R")
	nextDirection[key] = "E"

	key = fmt.Sprintf("%s-%s", "E", "R")
	nextDirection[key] = "S"

	key = fmt.Sprintf("%s-%s", "S", "R")
	nextDirection[key] = "W"

	key = fmt.Sprintf("%s-%s", "W", "R")
	nextDirection[key] = "N"

	return nextDirection
}

func getMoveValues() map[string]int {
	var moveValues = make(map[string]int)

	key := fmt.Sprintf("%s-%s", "N", "F")
	moveValues[key] = 1
	key = fmt.Sprintf("%s-%s", "N", "B")
	moveValues[key] = -1

	key = fmt.Sprintf("%s-%s", "S", "F")
	moveValues[key] = -1
	key = fmt.Sprintf("%s-%s", "S", "B")
	moveValues[key] = 1

	key = fmt.Sprintf("%s-%s", "E", "F")
	moveValues[key] = 1
	key = fmt.Sprintf("%s-%s", "E", "B")
	moveValues[key] = -1

	key = fmt.Sprintf("%s-%s", "W", "F")
	moveValues[key] = -1
	key = fmt.Sprintf("%s-%s", "W", "B")
	moveValues[key] = 1

	return moveValues
}

func getRoomSize() Rect {
	fmt.Println("Enter room size separated by comma:")
	sizeArr := getInputFromCmd()
	if len(sizeArr) != 2 {
		log.Fatal("inital data length is wrong")
	}
	var err error
	var room Rect
	room.Width, err = strconv.Atoi(sizeArr[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	room.Height, err = strconv.Atoi(sizeArr[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	if room.Height <= 0 || room.Width <= 0 {
		log.Fatal("room size is wrong")
	}

	return room
}

func getCarInput(room Rect) CarStruct {
	fmt.Println("Enter car position and direction separated by comma:")
	carArr := getInputFromCmd()
	if len(carArr) != 3 {
		log.Fatal("Initial data length is wrong")
	}
	var err error
	var car CarStruct
	car.X, err = strconv.Atoi(carArr[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	car.Y, err = strconv.Atoi(carArr[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	var direction = Dir(carArr[2])
	if isValidDirection(direction) {
		car.Dir = direction
	} else {
		log.Fatal(errors.New("Wrong direction"))
	}

	if car.X < 0 || car.Y < 0 {
		log.Fatal("Car initial location is wrong")
	}

	if room.Width <= car.X || room.Height <= car.Y {
		log.Fatal("Car is not in the room")
	}

	return car
}

func isValidDirection(direction Dir) bool {
	switch direction {
	case
		"N",
		"W",
		"S",
		"E":
		return true
	}
	return false
}

func getInputFromCmd() []string {
	reader := bufio.NewReader(os.Stdin)
	var err error
	initData, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	initData = strings.TrimSpace(initData)

	return strings.Split(initData, ",")
}

func execAction(room Rect, car CarStruct) {
	fmt.Println("Enter command separated by comma:")
	action := getInputFromCmd()
	carResult, err := runCmd(room, car, action)
	if err != nil {
		log.Fatalln("Fail,", err)
	}
	fmt.Println("Success, final car position and direction: ", carResult)
}

func runCmd(room Rect, car CarStruct, action []string) (CarStruct, error) {
	for _, order := range action {
		switch order {
		case "L", "R":
			key := fmt.Sprintf("%s-%s", car.Dir, order)
			car.Dir = Dir(nextDirection[key])
		case "F", "B":
			if proposedPos, isvalid := makeAMove(room, car, order); isvalid {
				car.X = proposedPos.Easting
				car.Y = proposedPos.Northing
			} else {
				return CarStruct{}, errors.New("Crushed into wall")
			}
		default:
			return CarStruct{}, errors.New("Invalid input command")
		}
	}

	return car, nil
}

func makeAMove(room Rect, car CarStruct, order string) (Pos, bool) {
	var proposedPos Pos
	proposedPos.Easting = car.X
	proposedPos.Northing = car.Y

	aspect := car.Dir
	key := fmt.Sprintf("%s-%s", aspect, order)

	switch aspect {
	case "N", "S":
		proposedPos.Northing = car.Y + moveValues[key]
	case "E", "W":
		proposedPos.Easting = car.X + moveValues[key]
	default:
		panic("invalid aspect")
	}

	return proposedPos, room.contains(proposedPos)
}

func (r Rect) contains(pos Pos) bool {
	return 0 <= pos.Easting && pos.Easting < r.Width && 0 <= pos.Northing && pos.Northing < r.Height
}

func main() {
	room := getRoomSize()
	car := getCarInput(room)
	fmt.Printf("Room size: %d * %d; car's location: (%d, %d), car's direction: %s\n", room.Width, room.Height, car.X, car.Y, car.Dir)

	execAction(room, car)
}
