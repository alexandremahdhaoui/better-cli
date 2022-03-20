package api

import (
	"fmt"
	"testing"
)

func getCLI() *CLI {
	arg0 := ArgSpec{
		Name:        "arg0",
		Description: "this is the 1st positional argument and must be an integer",
		Regex:       "int",
	}
	arg1 := ArgSpec{
		Name:        "arg1",
		Description: "this is the 2nd positional argument and must be a lower alphanumerical with dash",
		Regex:       "lower_alpha_dash",
	}

	flag0 := FlagSpec{
		Name:        "flag0",
		Description: "this is the flag0",
		Short:       "f",
		Long:        "flag0",
		PArgs:       ApiPSpec{&arg0, &arg1},
	}

	flag1 := FlagSpec{
		Name:        "flag1",
		Description: "this is the flag1",
		Short:       "F",
		Long:        "flag1",
		PArgs:       ApiPSpec{&arg0, &arg1},
	}
	cmd := Cmd{
		Name:        "sub",
		Description: "this is the sub command",
		PArgs:       ApiPSpec{&arg0, &arg1},
		Flags:       ApiFSpec{&flag0, &flag1},
		F: func(a Args, kw KWArgs) {
			_ = fmt.Sprintf("args %v", a)
			_ = fmt.Sprintf("kwargs %v", kw)
		},
	}

	scmd0 := SCmd{
		Name:        "super0",
		Description: "Contains sub command and super command",
		Cmd:         []*Cmd{&cmd},
	}

	scmd1 := SCmd{
		Name:        "super1",
		Description: "Contains sub command and super command",
		SCmd:        []*SCmd{&scmd0},
		Cmd:         []*Cmd{&cmd},
	}

	cli := CLI{
		Name:        "test",
		Description: "testCLI",
		Version:     "0.1.0",
		SCmd:        []*SCmd{&scmd0, &scmd1},
		Cmd:         []*Cmd{&cmd},
	}
	return &cli
}

func TestCLI_Run(t *testing.T) {
	cli := getCLI()
	cli.Run("super0 sub 0 asd-asd")
}
