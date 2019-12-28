package main

import (
	"fmt"
)

type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

func (r Rect) String() string {
	return fmt.Sprintf("Rect ==> width=%v, height=%v", r.Width, r.Height)
}

func (r Rect) DoubleSize() {
	r.Width *= 2
	r.Height *= 2
	fmt.Println("DoubleSize()", r)
}

func main() {
	r := Rect{2, 4}
	fmt.Printf("Rect area=%v\n", r.Area())
	fmt.Println(r)

	r.DoubleSize()
	fmt.Println("main()", r)
}
