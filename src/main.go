package main

import (
	"log"
	"os"
	"strings"

	"github.com/ACMEPORTALCOMPANY/html3d/faces"
	"github.com/ACMEPORTALCOMPANY/html3d/object"
	"github.com/ungerik/go3d/vec2"
	"github.com/ungerik/go3d/vec3"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		path := args[0]

		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("ERROR: unable to open [ %s ]: %s", path, err.Error())
		}
		defer file.Close()
		o := object.Obj(file)

		var faces2D [][]vec2.T
		for _, f := range o.Fs {
			vA := o.Vs[f.VFs[0].V]
			a := &vec3.T{vA.X, vA.Y, vA.Z}

			vB := o.Vs[f.VFs[1].V]
			b := &vec3.T{vB.X, vB.Y, vB.Z}

			vC := o.Vs[f.VFs[2].V]
			c := &vec3.T{vC.X, vC.Y, vC.Z}

			face2D := faces.From3D(a, b, c)
			faces2D = append(faces2D, face2D)
		}

		faces.NormalizeFaces(faces2D)

		filePath := strings.Split(file.Name(), "/")
		filename := strings.Split(filePath[len(filePath)-1], ".")[0]

		className := "face"

		if len(args) > 1 {
			flags := args[1:]

			for i := 0; i < len(flags); i += 2 {
				if len(flags) < i+2 {
					log.Fatalf("ERROR: flag must be provided arg")
				}

				switch flags[i] {
				case "-o":
					filename = flags[i+1]
				case "-c":
					className = flags[i+1]
				default:
					log.Fatalf("ERROR: illegal arg [ %s ]", flags[i])
				}
			}
		}

		faces.Render(filename, className, faces2D)
	} else {
		log.Fatal("ERROR: no args provided")
	}
}
