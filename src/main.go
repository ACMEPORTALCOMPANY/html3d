package main

import (
	"fmt"
	"github.com/ACMEPORTALCOMPANY/html3d/geometry"
	"github.com/ACMEPORTALCOMPANY/html3d/parse"
	"github.com/ACMEPORTALCOMPANY/html3d/render"
	"log"
	"os"
	"strconv"
	"strings"
)

var class = "face"
var fill = "none"
var output = "../out/"
var size = 200
var stroke = "black"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		exitOnError("no args provided")
	}

	path := strings.Split(args[0], "/")
	output += strings.Split(path[len(path)-1], ".")[0]

	if len(args) > 1 {
		flags(args[1:])
	}

	log.Print("file: " + args[0])

	o3 := parseObjFile(args[0]).Normalize(float64(size))
	o2 := o3.Project2D().Normalize(float64(size))

	err := render.HTML(o2, class, fill, output, stroke, size)
	if err != nil {
		exitOnError(err.Error())
	}

	err = render.CSS(class, output)
	if err != nil {
		exitOnError(err.Error())
	}
}

func parseObjFile(path string) *geometry.O3 {
	file, err := os.Open(path)
	if err != nil {
		exitOnError(err.Error())
	} else {
		defer func() {
			if err := file.Close(); err != nil {
				exitOnError(err.Error())
			}
		}()
	}

	obj, err := parse.Parse(file)
	if err != nil {
		exitOnError(err.Error())
	}

	return obj
}

func flags(args []string) {
	if len(args)%2 != 0 {
		exitOnError("flags must take arguments")
	}

	for i := 0; i < len(args); i += 2 {
		switch args[i] {
		case "-class":
			class = args[i+1]
		case "-fill":
			fill = args[i+1]
		case "-output":
			output = "../out/" + args[i+1]
		case "-size":
			s, err := strconv.Atoi(args[i+1])
			if err != nil {
				exitOnError(fmt.Sprintf("unable to parse size %s", args[i+1]))
			} else {
				size = s
			}
		case "-stroke":
			stroke = args[i+1]
		default:
			exitOnError("unknown flag " + args[i])
		}
	}

	log.Print("fill: " + fill)
	log.Print("class: " + class)
	log.Print("output: " + output)
	log.Printf("size: %d", size)
	log.Print("stroke: " + stroke)
}

func exitOnError(msg string) {
	log.Fatalf("ERROR - %s", msg)
}
