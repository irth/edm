package config_test

import (
	"testing"

	"github.com/irth/edm/config"
	"github.com/stretchr/testify/assert"
)

func TestVariablesGet(t *testing.T) {
	vars := config.Variables{
		"a": map[string]interface{}{
			"b": 3,
			"c": "d",
			"nesting": map[string]interface{}{
				"g": "h",
			},
		},
		"e": "f",
	}

	for _, tt := range []struct {
		Expected interface{}
		Path     []string
	}{
		{vars["a"], []string{"a"}},
		{3, []string{"a", "b"}},
		{"d", []string{"a", "c"}},
		{
			vars["a"].(map[string]interface{})["nesting"],
			[]string{"a", "nesting"},
		},
		{"h", []string{"a", "nesting", "g"}},
		{"f", []string{"e"}},
	} {
		v, err := vars.Get(tt.Path)
		assert.NoError(t, err)
		assert.Equal(t, tt.Expected, v)
	}
}
