package env

import (
	"os"
	"reflect"
	"testing"
)

func TestParseValue(t *testing.T) {
	os.Setenv("ENV", "develop")
	os.Setenv("PORT", "8080")
	defer os.Setenv("ENV", "")
	defer os.Setenv("PORT", "")

	var s struct {
		ENV  string `env:"ENV"`
		Port string `env:"PORT"`
	}

	expected := struct {
		ENV  string `env:"ENV"`
		Port string `env:"PORT"`
	}{
		ENV:  "develop",
		Port: "8080",
	}

	Parse(&s)

	if !reflect.DeepEqual(expected, s) {
		t.Errorf("%#v is expected but got %#v\n", expected, s)
	}
}

func TestParseValueWithDefault(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Setenv("PORT", "")

	var s = struct {
		ENV  string `env:"ENV"`
		Port string `env:"PORT"`
	}{
		ENV: "local",
	}

	expected := struct {
		ENV  string `env:"ENV"`
		Port string `env:"PORT"`
	}{
		ENV:  "local",
		Port: "8080",
	}

	Parse(&s)

	if !reflect.DeepEqual(expected, s) {
		t.Errorf("%#v is expected but got %#v\n", expected, s)
	}
}
