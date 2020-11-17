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
	if len(r.in.args) < 2 {
		r.out.Fatal("Not enough args")
	}
	rt := getRoute(r.routes, r.in.args[1])
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
