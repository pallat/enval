package env

func Example_parse() {
	var env struct {
		PORT string `env:"PORT"`
		ENV  string `env:"ENV"`
	}

	enval.Parse(&env)
}
