package spheres

import (
	"math"
	"testing"

	"github.com/bas-velthuizen/go-raytracer/materials"

	"github.com/bas-velthuizen/go-raytracer/transformations"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: A sphere's default transformation
// Given s ← sphere()
// Then s.Transform = identity_matrix
func Test_Sphere_has_Default_Transformation(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// Expected
	wanted := matrix.Identity(4)
	// Then
	if !wanted.Equals(s.Transform) {
		t.Errorf("Transform of %v = %v, expected %v", s, s.Transform, wanted)
	}
}

// Scenario: Changing a sphere's transformation
// Given s ← sphere()
// And t ← translation(2, 3, 4)
// When set_transform(s, t)
// Then s.Transform = t
func Test_Changing_a_Sphere_s_Transformation(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	trans := transformations.Translation(2, 3, 4)
	// When
	s.SetTransform(trans)
	// Then
	if !trans.Equals(s.Transform) {
		t.Errorf("Transform of %v = %v, expected %v", s, s.Transform, trans)
	}
}

// Scenario: The normal on a sphere at a point on the x axis
// Given s ← sphere()
// When n ← normal_at(s, point(1, 0, 0))
// Then n = vector(1, 0, 0)
func Test_Normal_of_a_Sphere_at_a_Point_on_the_X_Axis(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	point := tuples.Point(1, 0, 0)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(1, 0, 0)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal on a sphere at a point on the y axis
// Given s ← sphere()
// When n ← normal_at(s, point(0, 1, 0))
// Then n = vector(0, 1, 0)
func Test_Normal_of_a_Sphere_at_a_Point_on_the_Y_Axis(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	point := tuples.Point(0, 1, 0)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 1, 0)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal on a sphere at a point on the z axis
// Given s ← sphere()
// When n ← normal_at(s, point(0, 0, 1))
// Then n = vector(0, 0, 1)
func Test_Normal_of_a_Sphere_at_a_Point_on_the_Z_Axis(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	point := tuples.Point(0, 0, 1)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 0, 1)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal on a sphere at a non-axial point
// Given s ← sphere()
// When n ← normal_at(s, point(√3/3, √3/3, √3/3))
// Then n = vector(√3/3, √3/3, √3/3)
func Test_Normal_of_a_Sphere_at_a_NonAxial_Point(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	val := math.Sqrt(3) / 3.0
	point := tuples.Point(val, val, val)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(val, val, val)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: The normal is a normalized vector
// Given s ← sphere()
// When n ← normal_at(s, point(√3/3, √3/3, √3/3))
// Then n = normalize(n)
func Test_Normal_is_a_Normalized_Vector(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	val := math.Sqrt(3) / 3.0
	point := tuples.Point(val, val, val)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := n.Normalize()
	// And
	if !n.IsVector() {
		t.Errorf("normal_at( %v, %v ) is Point , wanted Vector", s, point)
	}
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v (normalized) , wanted %v", s, point, n, wanted)
	}
}

// Scenario: Computing the normal on a translated sphere
// Given s ← sphere()
// And set_transform(s, translation(0, 1, 0))
// When n ← normal_at(s, point(0, 1.70711, -0.70711))
// Then n = vector(0, 0.70711, -0.70711)
func Test_Computing_the_Normal_on_a_Translated_Sphere(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	s.SetTransform(transformations.Translation(0, 1, 0))
	// And
	point := tuples.Point(0, 1.70711, -0.70711)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 0.70711, -0.70711)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: Computing the normal on a scaled sphere
// Given s ← sphere()
// And set_transform(s, scaling(1, 0.5, 1))
// When n ← normal_at(s, point(0, √2/2, -√2/2))
// Then n = vector(0, 0.97014, -0.24254)
func Test_Computing_the_Normal_on_a_scaled_Sphere(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	s.SetTransform(transformations.Scaling(1, 0.5, 1))
	// And
	point := tuples.Point(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
	// When
	n := s.NormalAt(point)
	// Expected
	wanted := tuples.Vector(0, 0.97014, -0.24254)
	// Then
	if !wanted.Equals(*n) {
		t.Errorf("normal_at( %v, %v ) = %v , wanted %v", s, point, n, wanted)
	}
}

// Scenario: A sphere has a default material
// Given s ← sphere()
// When m ← s.material
// Then m = material()
func Test_a_Sphere_has_a_Default_Material(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// When
	wantedMaterial := materials.DefaultMaterial()
	if !wantedMaterial.Equals(s.Material) {
		t.Errorf("%v has material %v, expected %v", s, s.Material, wantedMaterial)
	}
}

// Scenario: A sphere may be assigned a material
// Given s ← sphere()
// And m ← material()
// And m.ambient ← 1
// When s.material ← m
// Then s.material = m
func Test_a_Sphere_May_Be_Assigned_a_Material(t *testing.T) {
	// Given
	s := NewUnitSphere()
	// And
	m := materials.DefaultMaterial()
	// And
	m.Ambient = 1.0
	// When
	s.Material = m
	// Then
	if !m.Equals(s.Material) {
		t.Errorf("%v has material %v, expected %v", s, s.Material, m)
	}
}
