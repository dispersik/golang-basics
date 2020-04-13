package main

import (
	"fmt"
	"sort"
)

func median(i []int) float32 {
	sort.Ints(i)
	l := len(i)
	if l%2 == 0 {
		return (float32(i[int(l/2)-1]) + float32(i[int(l/2)])) / 2
	} else {
		return float32(i[int(l/2)])
	}
}

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y - int(s.a)}
}

func (s Square) Perimeter() uint {
	return 4 * s.a
}

func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	fmt.Println(median([]int{2, 1, 3, 4, 7}))
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
