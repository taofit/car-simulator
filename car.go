package main

type CarStruct struct {
	X, Y int
	Dir  Dir
}

type Dir string
type Pos struct{ Easting, Northing string }
type Rect struct{ Width, Height int }
