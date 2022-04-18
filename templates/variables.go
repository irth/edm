package templates

import (
	"text/template/parse"
)

func GetTemplateVariables(tmpl string) ([][]string, error) {
	vars := [][]string{}

	template, err := ParseTemplateAST(tmpl)
	if err != nil {
		return nil, err
	}

	visitors := Visitors{
		FieldNode: func(n *parse.FieldNode) (bool, error) {
			vars = append(vars, n.Ident)
			return true, nil
		},
	}

	err = Walk(template.Root, visitors)
	if err != nil {
		return nil, err

	}

	return vars, nil
}
