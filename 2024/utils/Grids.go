package utils

import "image"

func AllDirections() []image.Point {
	return []image.Point{
		image.Point{0, -1},
		image.Point{1, 0},
		image.Point{0, 1},
		image.Point{-1, 0},
	}
}
