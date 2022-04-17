package templates

import (
	"fmt"
	"text/template"
	"text/template/parse"
)

func ParseTemplateAST(tmpl string) (*parse.Tree, error) {
	funcNames := []string{
		"and", "call", "html", "index", "slice", "js", "len", "not", "or",
		"print", "printf", "println", "urlquery", "eq", "ge", "gt", "le",
		"lt", "ne",
	}
	funcs := template.FuncMap{}
	for _, f := range funcNames {
		funcs[f] = func() {}
	}
	parsed, err := parse.Parse("template", string(tmpl), "{{", "}}", funcs)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse template: %w", err)
	}

	return parsed["template"], nil
}
