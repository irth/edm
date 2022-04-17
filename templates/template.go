package templates

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Template struct {
	*template.Template
	Variables [][]string
}

func LoadTemplate(path string) (*Template, error) {
	t := Template{}
	t.Template = template.New("main")

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	codeBytes, err := ioutil.ReadAll(f)
	code := string(codeBytes)

	vars, err := GetTemplateVariables(code)
	if err != nil {
		return nil, err
	}
	t.Variables = vars

	t.Parse(code)

	return &t, nil
}
