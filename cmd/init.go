package cmd

func StartApp() {
	counter := NewCounter()
	buildDirector := NewCounterBuilder(counter)

	buildDirector.BuildCounter()
}
