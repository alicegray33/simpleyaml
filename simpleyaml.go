// a Go package to interact with arbitrary YAML.
//
// Example:
//
//	var data = []byte(`
//	name: smallfish
//	age: 99
//	bool: true
//	bb:
//	    cc:
//	        dd:
//	            - 111
//	            - 222
//	            - 333
//	`
//
//	y, err := simpleyaml.NewYaml(data)
//	if err != nil {
//		// ERROR
//	}
//
//	if v, err := y.Get("name").String(); err == nil {
//		fmt.Println("value:", v)
//	}
//
//	// y.Get("age").Int()
//	// y.Get("bool").Bool()
//	// y.Get("bb").Get("cc").Get("dd").Array()
//	// y.Get("bb").Get("cc").Get("dd").GetIndex(1).Int()
//	// y.GetPath("bb", "cc", "ee").String()
package simpleyaml

import (
	"errors"

	"github.com/geofffranks/yaml"
)

type Yaml struct {
	data interface{}
}

// NewYaml returns a pointer to a new `Yaml` object after unmarshaling `body` bytes
func NewYaml(body []byte) (*Yaml, error) {
	val := yaml.MapSlice{}
	err := yaml.Unmarshal(body, &val)
	if err != nil {
		return nil, errors.New("unmarshal []byte to yaml failed: " + err.Error())
	}
	return &Yaml{val}, nil
}

// Get returns a pointer to a new `Yaml` object for `key` in its `map` representation
//
// Example:
//
//	y.Get("xx").Get("yy").Int()
func (y *Yaml) Get(key string) *Yaml {
	m, err := y.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Yaml{val}
		}
	}
	return &Yaml{nil}
}

// GetPath searches for the item as specified by the branch
//
// Example:
//
//	y.GetPath("bb", "cc").Int()
func (y *Yaml) GetPath(branch ...string) *Yaml {
	yin := y
	for _, p := range branch {
		yin = yin.Get(p)
	}
	return yin
}

// Array type asserts to an `array`
func (y *Yaml) Array() ([]interface{}, error) {
	if a, ok := (y.data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

// GetIndex returns a pointer to a new `Yaml` object.
// for `index` in its `array` representation
//
// Example:
//
//	y.Get("xx").GetIndex(1).String()
func (y *Yaml) GetIndex(index int) *Yaml {
	a, err := y.Array()
	if err == nil {
		if len(a) > index {
			return &Yaml{a[index]}
		}
	}
	return &Yaml{nil}
}

// Int type asserts to `int`
func (y *Yaml) Int() (int, error) {
	if v, ok := (y.data).(int); ok {
		return v, nil
	}
	return 0, errors.New("type assertion to int failed")
}

// Bool type asserts to `bool`
func (y *Yaml) Bool() (bool, error) {
	if v, ok := (y.data).(bool); ok {
		return v, nil
	}
	return false, errors.New("type assertion to bool failed")
}

// String type asserts to `string`
func (y *Yaml) String() (string, error) {
	if v, ok := (y.data).(string); ok {
		return v, nil
	}
	return "", errors.New("type assertion to string failed")
}

// Map type asserts to `map`
func (y *Yaml) Map() (map[interface{}]interface{}, error) {
	if m, ok := (y.data).(map[interface{}]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[interface]interface{} failed")
}
