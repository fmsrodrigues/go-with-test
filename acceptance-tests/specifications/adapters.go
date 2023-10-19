package specifications

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

type CurseAdapter func(name string) string

func (c CurseAdapter) Curse(name string) (string, error) {
	return c(name), nil
}
