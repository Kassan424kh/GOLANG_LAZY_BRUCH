package LazyPoint

import (
	"math"

	P "github.com/Kassan424kh/golang-lazy-brush/V1/point"
)

type LazyPoint struct {
	X, Y float64
}

/**
 * Update the x and y values
 *
 * @param {Point} point
 */
func (l *LazyPoint) Update(p P.Point) {
	l.X = p.X
	l.Y = p.Y
}

/**
 * Move the point to another position using an angle and distance
 *
 * @param {number} angle The angle in radians
 * @param {number} distance How much the point should be moved
 */
func (l *LazyPoint) MoveByAngle(angle, distance float64) {
	// Rotate the angle based on the browser coordinate system ([0,0] in the top left)
	angleRotated := angle + (math.Pi / 2)

	l.X = l.X + (math.Sin(angleRotated) * distance)
	l.Y = l.Y - (math.Cos(angleRotated) * distance)
}

/**
 * Check if this point is the same as another point
 *
 * @param {Point} point
 * @returns {boolean}
 */
func (l *LazyPoint) EqualsTo(p P.Point) bool {
	return l.X == p.X && l.Y == p.Y
}

/**
 * Get the difference for x and y axis to another point
 *
 * @param {Point} point
 * @returns {Point}
 */
func (l *LazyPoint) GetDifferenceTo(p P.Point) P.Point {
	return P.Point{
		X: l.X - p.X,
		Y: l.Y - p.Y,
	}
}

/**
 * Calculate distance to another point
 *
 * @param {Point} point
 * @returns {Point}
 */
func (l *LazyPoint) GetDistanceTo(p P.Point) float64 {
	diff := l.GetDifferenceTo(p)

	return math.Sqrt(math.Pow(diff.X, 2) + math.Pow(diff.Y, 2))
}

/**
 * Calculate the angle to another point
 *
 * @param {Point} point
 * @returns {Point}
 */
func (l *LazyPoint) GetAngleTo(p P.Point) float64 {
	diff := l.GetDifferenceTo(p)

	return math.Atan2(diff.Y, diff.X)
}

/**
 * Return a simple object with x and y properties
 *
 * @returns {object}
 */
func (l *LazyPoint) ToObject() map[string]int {
	return map[string]int{
		"X": int(math.Round(l.X)),
		"Y": int(math.Round(l.Y)),
	}
}
