package env

import (
	"os"
	"reflect"
)

const (
	tagPrefix = "env"
)

// Parse parse environment variable to any struct with tag name "env"
func Parse(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	num := t.NumField()
	for i := 0; i < num; i++ {
		c := t.Field(i).Tag
		envKey := c.Get(tagPrefix)
		val := os.Getenv(envKey)
		if val == "" {
			continue
		}

		field := reflect.New(reflect.TypeOf(val))
		field.Elem().Set(reflect.ValueOf(val))

		v.FieldByName(t.Field(i).Name).Set(field.Elem())
	}
}
