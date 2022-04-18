package config

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Variables map[string]interface{}

func (v Variables) GetFor(service string, container string, path []string) (interface{}, error) {
	paths := [][]string{
		{service, container},
		{service},
		{"global"},
	}
	for _, p := range paths {
		p = append(p, path...)
		v, err := v.Get(p)
		if err != nil {
			var ok bool
			if err, ok = err.(LookupError); ok {
				continue
			}
			return nil, errors.Wrapf(err, "error looking up %s", p)
		}
		return v, nil
	}
	return v, LookupError{fmt.Sprintf("couldn't find %s in %s.%s.*, %s.*, and global.*", strings.Join(path, "."), service, container, service)}
}

type LookupError struct{ e string }

func (v LookupError) Error() string { return v.e }

func (v Variables) Get(path []string) (interface{}, error) {
	if len(path) == 0 {
		return map[string]interface{}(v), nil
	}

	var el map[string]interface{} = v
	var ok bool
	i := 0
	for i < len(path)-1 {
		elUntyped, ok := el[path[i]]
		if !ok {
			return nil, LookupError{fmt.Sprintf("unknown property: %s", strings.Join(path[:i], "."))}
		}
		el, ok = elUntyped.(map[string]interface{})
		if !ok {
			return nil, LookupError{fmt.Sprintf("dictionary expected: %s", strings.Join(path[:i], "."))}
		}
		i += 1
	}

	val, ok := el[path[i]]
	if !ok {
		return nil, LookupError{fmt.Sprintf("unknown property: %s", strings.Join(path[:i], "."))}
	}

	return val, nil
}
