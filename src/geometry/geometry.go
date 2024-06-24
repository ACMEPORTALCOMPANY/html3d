package geometry

import (
	"fmt"
	"math"
)

type V2 struct {
	X, Y float64
}

func (v *V2) String() string {
	return fmt.Sprintf("V2 - X: %.4f, Y: %.4f", v.X, v.Y)
}

func (v *V2) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

type V3 struct {
	X, Y, Z float64
}

func (v *V3) String() string {
	return fmt.Sprintf("V3 - X: %.4f, Y: %.4f, Z: %.4f", v.X, v.Y, v.Z)
}

func (v *V3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v *V3) Normalize() {
	mag := v.Magnitude()

	v.X = v.X / mag
	v.Y = v.Y / mag
	v.Z = v.Z / mag
}

func (v *V3) Cross(b *V3) *V3 {
	return &V3{
		X: v.Y*b.Z - v.Z*b.Y,
		Y: v.Z*b.X - v.X*b.Z,
		Z: v.X*b.Y - v.Y*b.X,
	}
}

func (v *V3) Dot(b *V3) float64 {
	return (v.X * b.X) + (v.Y * b.Y) + (v.Z * b.Z)
}

type F2 struct {
	A, B, C *V2
}

func (f *F2) String() string {
	return fmt.Sprintf("F2 - A: [ %s ], B: [ %s ], C: [ %s ]", f.A.String(), f.B.String(), f.C.String())
}

type F3 struct {
	A, B, C *V3
}

func (f *F3) String() string {
	return fmt.Sprintf("F3 - A: [ %s ],  B: [ %s ], C: [ %s ]", f.A.String(), f.B.String(), f.C.String())
}

func (f *F3) Project2D() *F2 {
	ab := &V3{
		X: f.B.X - f.A.X,
		Y: f.B.Y - f.A.Y,
		Z: f.B.Z - f.A.Z,
	}

	ac := &V3{
		X: f.C.X - f.A.X,
		Y: f.C.Y - f.A.Y,
		Z: f.C.Z - f.A.Z,
	}

	n := ab.Cross(ac)

	u := &V3{
		X: ab.X,
		Y: ab.Y,
		Z: ab.Z,
	}

	u.Normalize()
	v := n.Cross(u)
	v.Normalize()

	return &F2{
		A: &V2{
			X: 0,
			Y: 0,
		},
		B: &V2{
			X: ab.Dot(u),
			Y: ab.Dot(v),
		},
		C: &V2{
			X: ac.Dot(u),
			Y: ac.Dot(v),
		},
	}
}

type O2 struct {
	Faces []*F2
}

func (o *O2) Normalize(scale float64) {
	var xMin, xMax, yMin, yMax float64

	var vs []*V2
	for _, f := range o.Faces {
		vs = append(vs, f.A, f.B, f.C)
	}

	for i := range vs {
		if vs[i].X < xMin {
			xMin = vs[i].X
		}

		if vs[i].X > xMax {
			xMax = vs[i].X
		}

		if vs[i].Y < yMin {
			yMin = vs[i].Y
		}

		if vs[i].Y > yMax {
			yMax = vs[i].Y
		}
	}

	diagonal := math.Sqrt(math.Pow(xMax-xMin, 2) + math.Pow(yMax-yMin, 2))

	normalizeX := func(x float64) float64 {
		return ((x - xMin) / diagonal) * scale
	}

	normalizeY := func(y float64) float64 {
		return ((y - yMin) / diagonal) * scale
	}

	for _, f := range o.Faces {
		f.A.X, f.A.Y = normalizeX(f.A.X), normalizeY(f.A.Y)
		f.B.X, f.B.Y = normalizeX(f.B.X), normalizeY(f.B.Y)
		f.C.X, f.C.Y = normalizeX(f.C.X), normalizeY(f.C.Y)
	}
}

type O3 struct {
	Faces []*F3
}

func (o *O3) Project2D() *O2 {

	var faces []*F2
	for _, f := range o.Faces {
		faces = append(faces, f.Project2D())
	}

	return &O2{
		Faces: faces,
	}
}
