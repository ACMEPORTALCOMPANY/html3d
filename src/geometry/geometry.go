package geometry

import (
	"fmt"
	"math"
)

type V3 struct {
	X, Y, Z float64
}

func (v *V3) String() string {
	return fmt.Sprintf("V3 - X: %.4f, Y: %.4f, Z: %.4f", v.X, v.Y, v.Z)
}

func (v *V3) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func (v *V3) Normalize() *V3 {
	mag := v.Magnitude()

	v.X = v.X / mag
	v.Y = v.Y / mag
	v.Z = v.Z / mag

	return v
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

type F3 struct {
	A, B, C *V3
}

func (f *F3) String() string {
	return fmt.Sprintf("F3 - A: [ %s ],  B: [ %s ], C: [ %s ]", f.A.String(), f.B.String(), f.C.String())
}

func (f *F3) Project2D() *F3 {
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

	return &F3{
		A: &V3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		B: &V3{
			X: ab.Dot(u),
			Y: ab.Dot(v),
			Z: 0,
		},
		C: &V3{
			X: ac.Dot(u),
			Y: ac.Dot(v),
			Z: 0,
		},
	}
}

type O3 struct {
	Faces []*F3
}

func (o *O3) Normalize(scale float64) *O3 {
	var xMin, xMax, yMin, yMax, zMin, zMax float64

	var vs []*V3
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

		if vs[i].Z > zMax {
			zMax = vs[i].Z
		}

		if vs[i].Z < zMin {
			zMin = vs[i].Z
		}
	}

	diagonal := math.Sqrt(math.Pow(xMax-xMin, 2) + math.Pow(yMax-yMin, 2) + math.Pow(zMax-zMin, 2))

	normalize := func(x, xMin, xMargin float64) float64 {
		normX := (x - xMin) / diagonal

		return (normX * scale) + xMargin
	}

	normalizeY := func(y, yMin, yMargin float64) float64 {
		normY := (y - yMin) / diagonal

		return scale - (normY * scale) - yMargin
	}

	for _, f := range o.Faces {
		fXMin := math.Min(math.Min(f.A.X, f.B.X), f.C.X)
		fXMax := math.Max(math.Max(f.A.X, f.B.X), f.C.X)
		fYMin := math.Min(math.Min(f.A.Y, f.B.Y), f.C.Y)
		fYMax := math.Max(math.Max(f.A.Y, f.B.Y), f.C.Y)
		fZMin := math.Min(math.Min(f.A.Z, f.B.Z), f.C.Z)
		fZMax := math.Max(math.Max(f.A.Z, f.B.Z), f.C.Z)

		normXRange := (fXMax - fXMin) / diagonal
		normYRange := (fYMax - fYMin) / diagonal
		normZRange := (fZMax - fZMin) / diagonal

		xMargin := (scale - scale*normXRange) / 2
		yMargin := (scale - scale*normYRange) / 2
		zMargin := (scale - scale*normZRange) / 2

		f.A.X, f.A.Y, f.A.Z = normalize(f.A.X, fXMin, xMargin), normalizeY(f.A.Y, fYMin, yMargin), normalize(f.A.Z, fZMin, zMargin)
		f.B.X, f.B.Y, f.B.Z = normalize(f.B.X, fXMin, xMargin), normalizeY(f.B.Y, fYMin, yMargin), normalize(f.B.Z, fZMin, zMargin)
		f.C.X, f.C.Y, f.C.Z = normalize(f.C.X, fXMin, xMargin), normalizeY(f.C.Y, fYMin, yMargin), normalize(f.C.Z, fZMin, zMargin)
	}

	return o
}

func (o *O3) Project2D() *O3 {

	var faces []*F3
	for _, f := range o.Faces {
		faces = append(faces, f.Project2D())
	}

	return &O3{
		Faces: faces,
	}
}
