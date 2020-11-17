package term

import (
	"os"
)

type Output struct {
	out *os.File
	err *os.File
}

func (o Output) Write(msg... string) error {
	return write(o.out, msg...)
}

func (o Output) Error(msg... string) error {
	return write(o.err, msg...)
}

func (o Output) Fatal(msg... string) {
	_ = o.Error(msg...)
	exit(1)
}

func (o Output) Success(msg... string) {
	_ = o.Write(msg...)
	exit(0)
}

func exit(code int) {
	os.Exit(code)
}

func write(out *os.File, msg... string) error {
	for _, s := range msg {
		_, err := out.WriteString(s + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
