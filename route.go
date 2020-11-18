package term

type Route struct {
	Operation string
	MinArgs   int
	Arg       []ArgSpec
	Function  func(i Input, o Output)
}

type ArgSpec struct {
	Position   int
	Name       string
	Validation func(s string) bool
	Failure    string
}

// Evaluate that this arg conditions have been satisfied
func (as ArgSpec) evaluate(i Input) bool {
	if (as.Position + 2) >= len(i.args) {
		return false
	}
	return as.Validation(i.args[as.Position + 2])
}

func (as ArgSpec) isRequired(min int) bool {
	return (as.Position + 1) <= min
}

// Check all args meet the outlined criteria, show fatal error if not
func (r Route) EvaluateArgs(i Input, o Output) bool {
	for _, a := range r.Arg {
		if !a.evaluate(i) && a.isRequired(r.MinArgs) {
			o.Fatal(a.Failure)
			return false
		}
	}
	return true
}

// Check that we have the correct number of args
func (r Route) EvaluateArgCount(i Input) bool {
	return r.MinArgs <= len(i.args) - 2
}
