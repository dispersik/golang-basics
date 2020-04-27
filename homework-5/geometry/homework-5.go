package geometry;

import (
	"errors"
)

// ErrZeroLen it's an error for zero length side of a Square
var ErrZeroLen = errors.New("Inappropriate value of a")

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (Point, error) {
	if(s.a!=0) {
		return Point{s.start.x + int(s.a), s.start.y - int(s.a)}, nil
	} else {
		return s.start, ErrZeroLen
	}

}

func (s Square) Perimeter() (uint, error) {
	if (s.a!=0) {
		return 4 * s.a, nil
	} else {
		return 0, ErrZeroLen
	}

}

func (s Square) Area() (uint, error) {
	if(s.a!=0) {
		return s.a * s.a, nil
	} else {
		return 0, ErrZeroLen
	}
}