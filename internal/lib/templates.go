package lib

import (
	"html/template"
	"strings"
)

func FlashTemplates(filenames ...string) (*template.Template, error) {
	return template.ParseFiles(append(filenames,
		"pkg/templates/templates/flash.html",
		"pkg/templates/templates/nav.html")...)
}

func TrimProtocol(url string) string {
	if strings.HasPrefix(url, "http://") {
		return strings.TrimPrefix(url, "http://")
	} else if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}
	return url
}

func TemplatesFlashFuncMap(val string) (*template.Template, error) {
	TemplateFuncs := template.FuncMap{
		"trimProtocol": TrimProtocol,
	}

	// First, create a new template.
	tmpl := template.New("").Funcs(TemplateFuncs)

	// Then, parse the file.
	tmpl, err := tmpl.ParseFiles(val)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
