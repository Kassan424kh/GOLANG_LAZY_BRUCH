package LazyBrush

import (
	LP "github.com/Kassan424kh/golang-lazy-brush/v1/lazy-point"
	P "github.com/Kassan424kh/golang-lazy-brush/v1/point"
)

/**
 * constructor
 *
 * @param {object} settings
 * @param {number} settings.radius The radius for the lazy area
 * @param {boolean} settings.enabled
 */
type LazyBrush struct {
	Radius    float64
	isEnabled bool

	Pointer LP.LazyPoint
	Brush   LP.LazyPoint

	Angle    float64
	Distance float64
	hasMoved bool
}

func Init() LazyBrush {
	lb := LazyBrush{
		Radius:    30,
		Pointer:   LP.LazyPoint{X: 0, Y: 0},
		Brush:     LP.LazyPoint{X: 0, Y: 0},
		Angle:     0,
		Distance:  0,
		isEnabled: true,
		hasMoved:  false,
	}

	return lb
}

/**
 * Enable lazy brush calculations.
 *
 */
func (l *LazyBrush) Enable() {
	l.isEnabled = true
}

/**
 * Disable lazy brush calculations.
 *
 */
func (l *LazyBrush) Disable() {
	l.isEnabled = false
}

/**
 * @returns {boolean}
 */
func (l *LazyBrush) IsEnabled() bool {
	return l.isEnabled
}

/**
 * Update the radius
 *
 * @param {number} radius
 */
func (l *LazyBrush) SetRadius(r float64) {
	l.Radius = r
}

/**
 * Return the current radius
 *
 * @returns {number}
 */
func (l *LazyBrush) GetRadius() float64 {
	return l.Radius
}

/**
 * Return the brush coordinates as a simple object
 *
 * @returns {object}
 */
func (l *LazyBrush) GetBrushCoordinates() map[string]int {
	return l.Brush.ToObject()
}

/**
 * Return the pointer coordinates as a simple object
 *
 * @returns {object}
 */
func (l *LazyBrush) GetPointerCoordinates() map[string]int {
	return l.Pointer.ToObject()
}

/**
 * Return the brush as a LazyPoint
 *
 * @returns {LazyPoint}
 */
func (l *LazyBrush) GetBrush() LP.LazyPoint {
	return l.Brush
}

/**
 * Return the pointer as a LazyPoint
 *
 * @returns {LazyPoint}
 */
func (l *LazyBrush) GetPointer() LP.LazyPoint {
	return l.Pointer
}

/**
 * Return the angle between pointer and brush
 *
 * @returns {number} Angle in radians
 */
func (l *LazyBrush) GetAngle() float64 {
	return l.Angle
}

/**
 * Return the distance between pointer and brush
 *
 * @returns {number} Distance in pixels
 */
func (l *LazyBrush) GetDistance() float64 {
	return l.Distance
}

/**
 * Return if the previous update has moved the brush.
 *
 * @returns {boolean} Whether the brush moved previously.
 */
func (l *LazyBrush) BrushHasMoved() bool {
	return l.hasMoved
}

/**
 * Updates the pointer point and calculates the new brush point.
 *
 * @param {Point} newPointerPoint
 * @param {Object} options
 * @param {Boolean} options.both Force update pointer and brush
 * @returns {Boolean} Whether any of the two points changed
 */
func (l *LazyBrush) Update(newPointerPoint P.Point, both bool) bool {
	l.hasMoved = false
	if l.Pointer.EqualsTo(newPointerPoint) && !both {
		return false
	}

	l.Pointer.Update(newPointerPoint)

	if both {
		l.hasMoved = true
		l.Brush.Update(newPointerPoint)
		return true
	}

	if l.isEnabled {
		brushPointer := P.Point{X: l.Brush.X, Y: l.Brush.Y}
		l.Distance = l.Pointer.GetDistanceTo(brushPointer)
		l.Angle = l.Pointer.GetAngleTo(brushPointer)

		if l.Distance > l.Radius {
			l.Brush.MoveByAngle(l.Angle, l.Distance-l.Radius)
			l.hasMoved = true
		}
	} else {
		l.Distance = 0
		l.Angle = 0
		l.Brush.Update(newPointerPoint)
		l.hasMoved = true
	}

	return true
}
