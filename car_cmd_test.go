package main

import (
	"errors"
	"testing"
)

type expectedStruct struct {
	CarStruct
	error
}

type cmdTest struct {
	room     Rect
	car      CarStruct
	action   []string
	expected expectedStruct
}

var cmdTests = []cmdTest{
	{Rect{5, 5}, CarStruct{2, 3, "N"}, []string{"F", "R", "F", "F"}, expectedStruct{CarStruct{4, 4, "E"}, nil}},
	{Rect{5, 5}, CarStruct{2, 3, "N"}, []string{"R", "F", "F", "L", "F"}, expectedStruct{CarStruct{4, 4, "N"}, nil}},
	{Rect{6, 5}, CarStruct{3, 2, "E"}, []string{"F", "L", "F", "L", "F"}, expectedStruct{CarStruct{3, 3, "W"}, nil}},
	{Rect{15, 5}, CarStruct{3, 2, "E"}, []string{"F", "L", "F", "R", "F", "F", "F", "F"}, expectedStruct{CarStruct{8, 3, "E"}, nil}},
	{Rect{5, 5}, CarStruct{3, 2, "E"}, []string{"F", "L", "F", "R", "F", "F", "F", "F"}, expectedStruct{CarStruct{}, errors.New("Crushed into wall")}},
	{Rect{5, 5}, CarStruct{3, 2, "E"}, []string{"F", "L", "F", "R", "f", "F", "F", "F"}, expectedStruct{CarStruct{}, errors.New("Invalid input command")}},
}

func TestCmd(t *testing.T) {
	for _, test := range cmdTests {
		car, err := runCmd(test.room, test.car, test.action)
		if car.X != test.expected.X || car.Y != test.expected.Y {
			t.Errorf("Output car's position(%d, %d) not equal to expected (%d, %d)", car.X, car.Y, test.expected.X, test.expected.Y)
		}

		if car.Dir != test.expected.Dir {
			t.Errorf("Output car's direction: %s not equal to expected direction: %s", car.Dir, test.expected.Dir)
		}

		if err != nil && test.expected.error != nil && err.Error() != test.expected.Error() {
			t.Errorf("Output car's error %s not equal to expected error %s", err, test.expected.error)
		}
	}
}
