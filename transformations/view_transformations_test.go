package transformations

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: The transformation matrix for the default orientation
// Given from ← point(0, 0, 0)
// And to ← point(0, 0, -1)
// And up ← vector(0, 1, 0)
// When t ← view_transform(from, to, up)
// Then t = identity_matrix
func Test_the_Transformation_Matrix_for_the_Default_Orientation(t *testing.T) {
	// Given from ← point(0, 0, 0)
	fromPoint := tuples.Point(0, 0, 0)
	// And to ← point(0, 0, -1)
	toPoint := tuples.Point(0, 0, -1)
	// And up ← vector(0, 1, 0)
	upVector := tuples.Vector(0, 1, 0)
	// When t ← view_transform(from, to, up)
	trans := NewViewTransform(fromPoint, toPoint, upVector)
	// Then t = identity_matrix
	if !matrix.Identity(4).Equals(trans) {
		t.Errorf("NewViewTransform(%v, %v, %v) = %v, expected %v", fromPoint, toPoint, upVector, trans, matrix.Identity(4))
	}
}

// Scenario: A view transformation matrix looking in positive z direction
// Given from ← point(0, 0, 0)
// And to ← point(0, 0, 1)
// And up ← vector(0, 1, 0)
// When t ← view_transform(from, to, up)
// Then t = scaling(-1, 1, -1)
func Test_a_View_Transformation_Matrix_Looking_in_positive_Z_Direction(t *testing.T) {
	// Given from ← point(0, 0, 0)
	fromPoint := tuples.Point(0, 0, 0)
	// And to ← point(0, 0, 1)
	toPoint := tuples.Point(0, 0, 1)
	// And up ← vector(0, 1, 0)
	upVector := tuples.Vector(0, 1, 0)
	// When t ← view_transform(from, to, up)
	trans := NewViewTransform(fromPoint, toPoint, upVector)
	// Expected
	wanted := Scaling(-1, 1, -1)
	// Then t = identity_matrix
	if !wanted.Equals(trans) {
		t.Errorf("NewViewTransform(%v, %v, %v) = %v, expected %v", fromPoint, toPoint, upVector, trans, wanted)
	}
}

// Scenario: The view transformation moves the world
// Given from ← point(0, 0, 8)
// And to ← point(0, 0, 0)
// And up ← vector(0, 1, 0)
// When t ← view_transform(from, to, up)
// Then t = translation(0, 0, -8)
func Test_The_View_Transformation_Moves_the_World(t *testing.T) {
	// Given from ← point(0, 0, 8)
	// And to ← point(0, 0, 0)
	// And up ← vector(0, 1, 0)
	// When trans ← view_transform(from, to, up)
	// Then trans = translation(0, 0, -8)
}

// Scenario: An arbitrary view transformation
// Given from ← point(1, 3, 2)
// And to ← point(4, -2, 8)
// And up ← vector(1, 1, 0)
// When t ← view_transform(from, to, up)
// Then t is the following 4x4 matrix:
//       | -0.50709 | 0.50709 |  0.67612 | -2.36643 |
//       |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
//       | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
//       |  0.00000 | 0.00000 |  0.00000 |  1.00000 |
func Test_an_Arbitrary_View_Transformation(t *testing.T) {
	// Given from ← point(1, 3, 2)
	// And to ← point(4, -2, 8)
	// And up ← vector(1, 1, 0)
	// When t ← view_transform(from, to, up)
	// Then t is the following 4x4 matrix:
	//       | -0.50709 | 0.50709 |  0.67612 | -2.36643 |
	//       |  0.76772 | 0.60609 |  0.12122 | -2.82843 |
	//       | -0.35857 | 0.59761 | -0.71714 |  0.00000 |
	//       |  0.00000 | 0.00000 |  0.00000 |  1.00000 |
}
