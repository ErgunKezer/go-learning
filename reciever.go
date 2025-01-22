package main

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) setRadius(newRadius float64) {
	c.Radius = newRadius
}

func (c Circle) getCircle() Circle {
	return c
}
