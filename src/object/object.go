package object

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type O struct {
	Vs  []V
	VTs []VT
	VNs []VN
	Fs  []F
}

func (o O) String() string {
	str := "{"

	str += "Vs: ["
	for i, v := range o.Vs {
		str += fmt.Sprintf("{%s}", v.String())

		if i != len(o.Vs)-1 {
			str += ", "
		}
	}
	str += "], VTs: ["
	for i, vt := range o.VTs {
		str += fmt.Sprintf("{%s}", vt.String())

		if i != len(o.VTs)-1 {
			str += ", "
		}
	}
	str += "], VNs: ["
	for i, vn := range o.VNs {
		str += fmt.Sprintf("{%s}", vn.String())

		if i != len(o.VNs)-1 {
			str += ", "
		}
	}
	str += "], Fs: ["
	for i, f := range o.Fs {
		str += fmt.Sprintf("{%s}", f.String())

		if i != len(o.Fs)-1 {
			str += ", "
		}
	}

	return str + "]}"
}

type V struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v V) String() string {
	return fmt.Sprintf("X: %.6f, Y: %.6f, Z: %.6f, W: %.6f", v.X, v.Y, v.Z, v.W)
}

type VT struct {
	U float32
	V float32
	W float32
}

func (vt VT) String() string {
	return fmt.Sprintf("U: %.6f, V: %.6f, W: %.6f", vt.U, vt.V, vt.W)
}

type VN struct {
	X float32
	Y float32
	Z float32
}

func (vn VN) String() string {
	return fmt.Sprintf("X: %.6f, Y: %.6f, Z: %.6f", vn.X, vn.Y, vn.Z)
}

type F struct {
	VFs []VF
}

func (f F) String() string {
	str := "VFs: ["

	for i, v := range f.VFs {
		str += fmt.Sprintf("{%s}", v.String())

		if i != len(f.VFs)-1 {
			str += ", "
		}
	}

	return str + "]"
}

type VF struct {
	V  int
	VT int
	VN int
}

func (vf VF) String() string {
	return fmt.Sprintf("V: %d, VT: %d, VN: %d", vf.V, vf.VT, vf.VN)
}

func Obj(file *os.File) O {
	var vs []V
	var vts []VT
	var vns []VN
	var fs []F

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		switch line[0] {
		case "v":
			v := v(line)
			vs = append(vs, v)
		case "vt":
			vt := vt(line)
			vts = append(vts, vt)
		case "vn":
			vn := vn(line)
			vns = append(vns, vn)
		case "f":
			f := f(line)
			fs = append(fs, f)
		}
	}

	return O{
		Vs:  vs,
		VTs: vts,
		VNs: vns,
		Fs:  fs,
	}
}

func v(line []string) V {
	if len(line) < 4 || len(line) > 5 {
		log.Fatalf("ERROR: invalid V: %s.", strings.Join(line, " "))
	}

	x, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid V.X: %s.", line[1])
	}

	y, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid V.Y: %s.", line[2])
	}

	z, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid V.Z: %s.", line[3])
	}

	var w float64
	if len(line) > 4 {
		w, err = strconv.ParseFloat(line[4], 64)
		if err != nil {
			log.Fatalf("ERROR: invalid V.W: %s.", line[4])
		}
	} else {
		w = 1
	}

	return V{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
		W: float32(w),
	}
}

func vt(line []string) VT {
	if len(line) < 2 || len(line) > 4 {
		log.Fatalf("ERROR: invalid VT: %s.", strings.Join(line, " "))
	}

	u, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid VT.U: %s.", line[1])
	}

	var v float64
	if len(line) > 2 {
		v, err = strconv.ParseFloat(line[2], 64)
		if err != nil {
			log.Fatalf("ERROR: invalid VT.V: %s.", line[2])
		}
	} else {
		v = 0
	}

	var w float64
	if len(line) > 3 {
		w, err = strconv.ParseFloat(line[3], 64)
		if err != nil {
			log.Fatalf("ERROR: invalid VT.W: %s.", line[3])
		}
	} else {
		w = 0
	}

	return VT{
		U: float32(u),
		V: float32(v),
		W: float32(w),
	}
}

func vn(line []string) VN {
	if len(line) != 4 {
		log.Fatalf("ERROR: invalid VN: %s.", strings.Join(line, " "))
	}

	x, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid VN.X: %s.", line[1])
	}

	y, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid VN.Y: %s.", line[2])
	}

	z, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		log.Fatalf("ERROR: invalid VN.Z: %s.", line[3])
	}

	return VN{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}
}

func f(line []string) F {
	if len(line) < 4 {
		log.Fatalf("ERROR: invalid F: %s.", strings.Join(line, " "))
	}

	var vfs []VF
	for _, vStr := range line[1:] {
		vParts := strings.Split(vStr, "/")

		v, err := strconv.Atoi(vParts[0])
		if err != nil {
			log.Fatalf("ERROR: invalid VF.V: %s.", vParts[0])
		}
		v -= 1

		var vt int
		if len(vParts) > 1 {
			if strings.Trim(vParts[1], " ") == "" {
				vt = -1
			} else {
				vt, err = strconv.Atoi(vParts[1])
				if err != nil {
					log.Fatalf("ERROR: invalid VF.VT: %s.", vParts[1])
				}
				vt -= 1
			}
		} else {
			vt = -1
		}

		var vn int
		if len(vParts) > 2 {
			vn, err = strconv.Atoi(vParts[2])
			if err != nil {
				log.Fatalf("ERROR: invalid VF.VN: %s.", vParts[2])
			}
			vn -= 1
		} else {
			vn = -1
		}

		vf := VF{
			V:  v,
			VT: vt,
			VN: vn,
		}

		vfs = append(vfs, vf)
	}

	return F{
		VFs: vfs,
	}
}
