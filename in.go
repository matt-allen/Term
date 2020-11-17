package term

type Input struct {
	args []string
}

func (i Input) GetArg(p int) string {
	return i.args[p+2]
}
