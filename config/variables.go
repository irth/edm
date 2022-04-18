package config

import (
	"fmt"
	"strings"
)

type Variables map[string]interface{}

func (v Variables) Get(path []string) (interface{}, error) {
	var el map[string]interface{} = v
	var ok bool
	i := 0
	for i < len(path)-1 {
		elUntyped, ok := el[path[i]]
		if !ok {
			return nil, fmt.Errorf("unknown property: %s", strings.Join(path[:i], "."))
		}
		el, ok = elUntyped.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("dictionary expected: %s", strings.Join(path[:i], "."))
		}
		i += 1
	}

	val, ok := el[path[i]]
	if !ok {
		return nil, fmt.Errorf("unknown property: %s", strings.Join(path[:i], "."))
	}

	return val, nil
}
