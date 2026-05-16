package main

func main() {
	var modules = []Module{
		NewModule("bank"),
		NewModule("mine"),
	}

	InitializeModules(modules)
	PrintModules(modules)
}
