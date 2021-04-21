package main

import "fmt"

type point struct {
	x, y float32
}

type circle struct {
	centre point
	radius float32
}

func main() {
	c := circle{
		centre: point{1, 1},
		radius: 1,
	}
	fmt.Println(getRadius(c))
	updateRadius(&c, 1.1)
	fmt.Println(getRadius(c))
}

func getRadius(c circle) float32 {
	return c.radius
}

func updateRadius(c *circle, add float32) {
	c.radius += add
}

func valueDoesNotEscape() {
	c := circle{centre: point{1, 1}, radius: 1}
	getRadius(c)
}

func pointerDoesNotEscape() {
	c := circle{centre: point{1, 1}, radius: 1}
	updateRadius(&c, 1.1)
}

func escapes() *circle {
	c := circle{centre: point{1, 1}, radius: 1}
	return &c
}
