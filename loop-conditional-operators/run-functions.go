package main

func runFunctions() {
	functions := []func(){
		Cave,
		Conditionals,
		BooleanOp,
		Switch,

		// loop,
	}

	for _, f := range functions {
		f()
		Repeat()
	}
}