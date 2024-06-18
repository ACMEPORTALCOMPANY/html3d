package faces

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
)

func From3D(a, b, c *vec3.T) []vec2.T {
	ab := b.Sub(a)
	ac := c.Sub(a)

	n := vec3.Cross(ab, ac)

	u := ab.Normalized()
	v := vec3.Cross(&n, &u)
	v = *v.Normalize()

	aPrime := vec2.T{float32(0), float32(0)}
	bPrime := vec2.T{vec3.Dot(ab, &u), vec3.Dot(ab, &v)}
	cPrime := vec2.T{vec3.Dot(ac, &u), vec3.Dot(ac, &v)}

	return []vec2.T{aPrime, bPrime, cPrime}
}

func NormalizeFaces(faces2D [][]vec2.T) {
	var xMin, xMax, yMin, yMax float32

	var points2D []vec2.T
	for i := range faces2D {
		points2D = append(points2D, faces2D[i]...)
	}

	for i := range points2D {
		if points2D[i][0] < xMin {
			xMin = points2D[i][0]
		}

		if points2D[i][0] > xMax {
			xMax = points2D[i][0]
		}

		if points2D[i][1] < yMin {
			yMin = points2D[i][1]
		}

		if points2D[i][1] > yMax {
			yMax = points2D[i][1]
		}
	}

	xRange := xMax - xMin
	yRange := yMax - yMin
	diag := math.Sqrt(math.Pow(float64(xRange), 2) + math.Pow(float64(yRange), 2))

	for i := range faces2D {
		for j := range faces2D[i] {
			faces2D[i][j][0] = ((faces2D[i][j][0] - xMin) / float32(diag)) * 100
			faces2D[i][j][1] = ((faces2D[i][j][1] - yMin) / float32(diag)) * 100
		}
	}
}

func Render(filename, class string, faces2D [][]vec2.T) {
	css, err := os.Create(filename + ".css")
	if err != nil {
		log.Fatalf("ERROR: unable to create [ %s.css ]: %s", filename, err.Error())
	}

	html, err := os.Create(filename + ".html")
	if err != nil {
		log.Fatalf("ERROR: unable to create [ %s.html ]: %s", filename, err.Error())
	}

	defer func() {
		if err := css.Close(); err != nil {
			log.Fatalf("ERROR: unable to close [ %s.css ]: %s", filename, err.Error())
		}

		if err := html.Close(); err != nil {
			log.Fatalf("ERROR: unable to close [ %s.html ]: %s", filename, err.Error())
		}
	}()

	for i := range faces2D {
		styleFormat := "#f-%d{"
		styleFormat += "\n\t-webkit-clip-path: polygon(%.2f%% %.2f%%,%.2f%% %.2f%%,%.2f%% %.2f%%);"
		styleFormat += "\n\tclip-path: polygon(%.2f%% %.2f%%,%.2f%% %.2f%%,%.2f%% %.2f%%);"
		styleFormat += "\n}"

		xMin := faces2D[i][0][0]
		xMax := faces2D[i][0][0]
		yMin := faces2D[i][0][1]
		yMax := faces2D[i][0][1]

		for j := range faces2D[i] {
			if faces2D[i][j][0] < xMin {
				xMin = faces2D[i][j][0]
			}

			if faces2D[i][j][0] > xMax {
				xMax = faces2D[i][j][0]
			}

			if faces2D[i][j][1] < yMin {
				yMin = faces2D[i][j][1]
			}

			if faces2D[i][j][1] > yMax {
				yMax = faces2D[i][j][1]
			}
		}

		xRange := xMax - xMin
		yRange := yMax - yMin

		xMargin := (100 - (xRange)) / 2
		yMargin := (100 - (yRange)) / 2

		x1 := faces2D[i][0][0] + xMargin - xMin
		y1 := 100 - faces2D[i][0][1] - yMargin
		x2 := faces2D[i][1][0] + xMargin - xMin
		y2 := 100 - faces2D[i][1][1] - yMargin
		x3 := faces2D[i][2][0] + xMargin - xMin
		y3 := 100 - faces2D[i][2][1] - yMargin

		style := fmt.Sprintf(styleFormat, i, x1, y1, x2, y2, x3, y3, x1, y1, x2, y2, x3, y3)

		markdown := fmt.Sprintf("<svg class=\"%s\" id=\"f-%d\"></svg>", class, i)
		if i != len(faces2D)-1 {
			style += "\n"
			markdown += "\n"
		}

		if _, err := css.Write([]byte(style)); err != nil {
			panic(err)
		}

		if _, err := html.Write([]byte(markdown)); err != nil {
			panic(err)
		}
	}
}
