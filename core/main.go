package main

func main() {
	var modules = []Module{
		Module{
			id:   "bank",
			port: -1,
		},

		Module{
			id:   "mine",
			port: -1,
		},
	}

	InitializeModules(modules)
	DebugModules(modules)
}
