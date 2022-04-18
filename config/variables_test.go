package config_test

import (
	"testing"

	"github.com/irth/edm/config"
	"github.com/stretchr/testify/assert"
)

func TestVariablesGetFor(t *testing.T) {
	{
		vars := config.Variables{
			"global": map[string]interface{}{
				"a": 1,
			},
		}

		v, err := vars.GetFor("service1", "container1", []string{"a"})
		assert.NoError(t, err)
		assert.Equal(t, 1, v)

		vars["service1"] = map[string]interface{}{
			"a": 2,
		}
		v, err = vars.GetFor("service1", "container1", []string{"a"})
		assert.NoError(t, err)
		assert.Equal(t, 2, v)

		vars["service1"] = map[string]interface{}{
			"a": 2,
			"container1": map[string]interface{}{
				"a": 3,
			},
		}
		v, err = vars.GetFor("service1", "container1", []string{"a"})
		assert.NoError(t, err)
		assert.Equal(t, 3, v)

		_, err = vars.GetFor("service1", "container1", []string{"wrong", "path"})
		assert.IsType(t, config.LookupError{}, err)
	}
}

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

	_, err := vars.Get([]string{"wrong", "path"})
	assert.IsType(t, config.LookupError{}, err)

	self, err := vars.Get([]string{})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}(vars), self)
}
