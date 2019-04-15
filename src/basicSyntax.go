package main

import "fmt"

type point struct {
	x, y int
}

func (p point) getX() (px int) {
	px = p.x
	return px
}

func (p *point) getY() int {
	return p.y
}

// func setPointXY(p point, x, y int) {
// 	p.x = x
// 	p.y = y
// }
func setPointXY(p *point, x, y int) {
	p.x = x
	p.y = y
}
func baseAndPointer() {
	point1 := point{1, 2}
	println(point1.x, point1.y)

	println(point1.getX())
	println(point1.getY())

	pointptr := &point1
	println(pointptr.getX())
	println(pointptr.getY())

	setPointXY(&point1, 100, 200)
	println(pointptr.getX())
	println(pointptr.getY())
	println(point1.getX())
	println(point1.getY())

	setPointXY(pointptr, 1000, 2000)
	println(pointptr.getX())
	println(pointptr.getY())
	println(point1.getX())
	println(point1.getY())

	fmt.Println(point1)

}
