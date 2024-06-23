package parse

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type V struct {
	X, Y, Z float32
}

func (v *V) String() string {
	return fmt.Sprintf("V - X: %.4f, Y: %.4f, Z: %.4f", v.X, v.Y, v.Z)
}

type F struct {
	A, B, C *V
}

func (f *F) String() string {
	return fmt.Sprintf("F - A: [ %s ],  B: [ %s ], C: [ %s ]", f.A.String(), f.B.String(), f.C.String())
}

type Object struct {
	Faces []*F
}

var vList []*V
var fList []*F

func Parse(file *os.File) (*Object, error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "v":
			v, err := v(line)
			if err != nil {
				return nil, err
			}

			log.Print(v.String())

			vList = append(vList, v)
		case "f":
			f, err := f(line, vList)
			if err != nil {
				return nil, err
			}

			log.Print(f.String())

			fList = append(fList, f)
		}
	}

	return &Object{
		Faces: fList,
	}, nil
}

func v(line []string) (*V, error) {
	if len(line) < 4 || len(line) > 5 {
		return nil, errors.New(fmt.Sprintf("invalid V: %s", strings.Join(line, " ")))
	}

	x, err := strconv.ParseFloat(line[1], 32)
	if err != nil {
		return nil, errors.New("invalid V.X: " + line[1])
	}

	y, err := strconv.ParseFloat(line[2], 32)
	if err != nil {
		return nil, errors.New("invalid V.Y: " + line[2])
	}

	z, err := strconv.ParseFloat(line[3], 32)
	if err != nil {
		return nil, errors.New("invalid V.Z: " + line[3])
	}

	return &V{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}, nil
}

func f(line []string, vList []*V) (*F, error) {
	if len(line) != 4 {
		return nil, errors.New(fmt.Sprintf("invalid F: %s", strings.Join(line, " ")))
	}

	aIndex, err := strconv.Atoi(strings.Split(line[1], "/")[0])
	if err != nil {
		return nil, errors.New("invalid F.A: " + strings.Split(line[1], "/")[0])
	}

	bIndex, err := strconv.Atoi(strings.Split(line[2], "/")[0])
	if err != nil {
		return nil, errors.New("invalid F.B: " + strings.Split(line[2], "/")[0])
	}

	cIndex, err := strconv.Atoi(strings.Split(line[3], "/")[0])
	if err != nil {
		return nil, errors.New("invalid F.C: " + strings.Split(line[3], "/")[0])
	}

	return &F{
		A: vList[aIndex-1],
		B: vList[bIndex-1],
		C: vList[cIndex-1],
	}, nil
}
