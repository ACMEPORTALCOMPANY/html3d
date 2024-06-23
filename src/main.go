package main

import (
	"fmt"
	"github.com/ACMEPORTALCOMPANY/html3d/render"
	"log"
	"os"
	"strconv"
	"strings"
)

var s = 200
var o = "out/"
var c = "face"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		exitOnError("no args provided")
	}

	parts := strings.Split(args[0], "/")
	o += strings.Split(parts[len(parts)-1], ".")[0]

	if len(args) > 1 {
		flags(args[1:])
	}

	log.Printf("running on %s w/ c = %s, o = %s, s = %d", args[0], c, o, s)

	file, err := os.Open(args[0])
	if err != nil {
		exitOnError(err.Error())
	} else {
		defer file.Close()
	}

	err = render.HTML(o, s)
	if err != nil {
		exitOnError(err.Error())
	}

	err = render.CSS(o, c)
	if err != nil {
		exitOnError(err.Error())
	}
}

func flags(args []string) {
	if len(args)%2 != 0 {
		exitOnError("flags must take arguments")
	}

	for i := 0; i < len(args); i += 2 {
		switch args[i] {
		case "-c":
			c = args[i+1]
		case "-o":
			o = "out/" + args[i+1]
		case "-s":
			size, err := strconv.Atoi(args[i+1])
			if err != nil {
				exitOnError(fmt.Sprintf("unable to parse size %s", args[i+1]))
			} else {
				s = size
			}
		default:
			exitOnError("unknown flag " + args[i])
		}
	}
}

func exitOnError(msg string) {
	log.Fatalf("ERROR - %s", msg)
}
