package term

type Input struct {
	args []string
}

func MockInput(i []string) Input {
	return Input{i}
}

func (i Input) GetArg(p int) string {
	return i.args[p + 2]
}

func (i Input) GetArgCount() int {
	return len(i.args) - 2
}
