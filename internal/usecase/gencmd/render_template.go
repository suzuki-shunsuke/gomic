package gencmd

import (
	"io"
	"strings"
	"text/template"
)

func renderTpl(writer io.Writer, txt string, data interface{}) error {
	t := template.New("dummy")
	tpl, err := t.Parse(strings.Trim(txt, "\n"))
	if err != nil {
		return err
	}
	return tpl.Execute(writer, data)
}
