package render

import (
	"fmt"
	"github.com/ACMEPORTALCOMPANY/html3d/geometry"
	"io/fs"
	"os"
)

func out() error {
	_, err := os.Stat("../out")
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir("../out", fs.ModeDir)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func CSS(className, filename string) error {
	err := out()
	if err != nil {
		return err
	}

	css, err := os.Create(filename + ".css")
	if err != nil {
		return err
	}

	defaultClass := "." + className + " {"
	defaultClass += "\n\tposition: absolute;"
	defaultClass += "\n}\n\n"

	if _, err := css.Write([]byte(defaultClass)); err != nil {
		return err
	}

	style := "#f-%d {"
	style += "\n\ttransform: rotate3D(%.2f, %.2f, %.2f, %.2frad);"
	style += "\n}\n\n"

	if err := css.Close(); err != nil {
		return err
	}

	return nil
}

func HTML(obj *geometry.O3, class, fill, output, stroke string, size int) error {
	err := out()
	if err != nil {
		return err
	}

	html, err := os.Create(output + ".html")
	if err != nil {
		return err
	}

	outerHTMLOpen := fmt.Sprintf("<svg viewBox=\"0 0 %d %d\">", size, size)

	if _, err := html.Write([]byte(outerHTMLOpen)); err != nil {
		return err
	}

	for i, f := range obj.Faces {
		format := "\n\t<polygon class=\"%s\" id=\"f-%d\" points=\"%.2f,%.2f %.2f,%.2f %.2f,%.2f\" fill=\"%s\" stroke=\"%s\" />"
		line := fmt.Sprintf(format, class, i, f.A.X, f.A.Y, f.B.X, f.B.Y, f.C.X, f.C.Y, fill, stroke)
		if _, err := html.Write([]byte(line)); err != nil {
			return err
		}
	}

	outerHTMLClose := "\n</svg>"

	if _, err := html.Write([]byte(outerHTMLClose)); err != nil {
		return err
	}

	if err := html.Close(); err != nil {
		return err
	}

	return nil
}
