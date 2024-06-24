package parse

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ACMEPORTALCOMPANY/html3d/geometry"
	"os"
	"strconv"
	"strings"
)

var vList []*geometry.V3
var fList []*geometry.F3

func Parse(file *os.File) (*geometry.O3, error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "v":
			v, err := v(line)
			if err != nil {
				return nil, err
			}

			vList = append(vList, v)
		case "f":
			f, err := f(line, vList)
			if err != nil {
				return nil, err
			}

			fList = append(fList, f)
		}
	}

	return &geometry.O3{
		Faces: fList,
	}, nil
}

func v(line []string) (*geometry.V3, error) {
	if len(line) < 4 || len(line) > 5 {
		return nil, errors.New(fmt.Sprintf("invalid geometry.V3: %s", strings.Join(line, " ")))
	}

	x, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		return nil, errors.New("invalid geometry.V3.X: " + line[1])
	}

	y, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		return nil, errors.New("invalid geometry.V3.Y: " + line[2])
	}

	z, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		return nil, errors.New("invalid geometry.V3.Z: " + line[3])
	}

	return &geometry.V3{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func f(line []string, vList []*geometry.V3) (*geometry.F3, error) {
	if len(line) != 4 {
		return nil, errors.New(fmt.Sprintf("invalid geometry.F3: %s", strings.Join(line, " ")))
	}

	aIndex, err := strconv.Atoi(strings.Split(line[1], "/")[0])
	if err != nil {
		return nil, errors.New("invalid geometry.F3.A: " + strings.Split(line[1], "/")[0])
	}

	bIndex, err := strconv.Atoi(strings.Split(line[2], "/")[0])
	if err != nil {
		return nil, errors.New("invalid geometry.F3.B: " + strings.Split(line[2], "/")[0])
	}

	cIndex, err := strconv.Atoi(strings.Split(line[3], "/")[0])
	if err != nil {
		return nil, errors.New("invalid geometry.F3.C: " + strings.Split(line[3], "/")[0])
	}

	return &geometry.F3{
		A: vList[aIndex-1],
		B: vList[bIndex-1],
		C: vList[cIndex-1],
	}, nil
}
