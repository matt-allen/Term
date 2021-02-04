package term

import "os"

type Router struct {
	in     Input
	out    Output
	routes []Route
}

func NewRouter(r []Route, a []string) Router {
	o := Output{os.Stdout, os.Stderr}
	i := Input{a}
	return Router{i,o,r}
}

func (r Router) Handle() {
	if len(r.in.args) < 1 {
		r.out.Fatal("Not enough args")
	}
	o := r.in.args[1]
	if o == "help" {
		r.printHelp()
		return
	}
	rt := getRoute(r.routes, o)
	if rt == nil {
		r.out.Fatal("Invalid operation")
	}
	if !rt.EvaluateArgCount(r.in) {
		r.out.Fatal("Not enough args")
	}
	rt.EvaluateArgs(r.in, r.out)
	rt.Function(r.in, r.out)
}

func getRoute(r []Route, o string) *Route {
	for _, a := range r {
		if a.Operation == o {
			return &a
		}
	}
	return nil
}

func (r Router) printHelp() {
	r.out.Write("Available options:\n")
	for _, o := range r.routes {
		r.out.Write(o.Operation)
	}
}
