package lib

import "text/template"

func TemplatesWithFlash(filenames ...string) (*template.Template, error) {
	return template.ParseFiles(append(filenames, "pkg/templates/flash.html")...)
}
