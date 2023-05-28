package main

type CarStruct struct {
	X, Y int
	Dir  Dir
}

type Dir string
type Pos struct{ Easting, Northing int }
type Rect struct{ Width, Height int }
