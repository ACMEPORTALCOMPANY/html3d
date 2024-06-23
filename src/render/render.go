package render

import (
	"fmt"
	"io/fs"
	"os"
)

func out() error {
	_, err := os.Stat("out")
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir("out", fs.ModeDir)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func CSS(filename, className string) error {
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
	defaultClass += "\n}"

	if _, err := css.Write([]byte(defaultClass)); err != nil {
		return err
	}

	if err := css.Close(); err != nil {
		return err
	}

	return nil
}

func HTML(filename string, size int) error {
	err := out()
	if err != nil {
		return err
	}

	html, err := os.Create(filename + ".html")
	if err != nil {
		return err
	}

	outerHTMLOpen := fmt.Sprintf("<svg viewBox=\"0 0 %d %d\">", size, size)
	outerHTMLClose := "\n</svg>"

	if _, err := html.Write([]byte(outerHTMLOpen)); err != nil {
		return err
	}

	if _, err := html.Write([]byte(outerHTMLClose)); err != nil {
		return err
	}

	if err := html.Close(); err != nil {
		return err
	}

	return nil
}
