package geometry;

import (
	"testing"
)

func TestEnd(t *testing.T) {
	testCases := []struct {
		in Square
		out Point
		err error
	}{
		{Square{Point{0,0},0},Point{0,0},ErrZeroLen},
		{Square{Point{0,1},1},Point{1,0},nil},
		{Square{Point{-1,1},1},Point{0,0},nil},
	}

	for _, c := range testCases {
		res, err := c.in.End()
		if res != c.out || err != c.err {
			t.Errorf("End(%v) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}
}

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		in Square
		out uint
		err error
	}{
		{Square{Point{0,0},0},0,ErrZeroLen},
		{Square{Point{0,1},1},4,nil},
		{Square{Point{-1,1},1},4,nil},
	}

	for _, c := range testCases {
		res, err := c.in.Perimeter()
		if res != c.out || err != c.err {
			t.Errorf("Perimeter(%v) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		in Square
		out uint
		err error
	}{
		{Square{Point{0,0},0},0,ErrZeroLen},
		{Square{Point{0,1},1},1,nil},
		{Square{Point{-1,1},1},1,nil},
	}

	for _, c := range testCases {
		res, err := c.in.Area()
		if res != c.out || err != c.err {
			t.Errorf("Area(%v) returns %v, %v; expected %v, %v", c.in, res, err, c.out, c.err)
		}
	}
}