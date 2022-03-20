package main

import (
	"fmt"
	"main/api"
)

func getCliExample() *api.CLI {
	arg0 := api.ArgSpec{
		Name:        "arg0",
		Description: "this is the 1st positional argument and must be an integer",
		Regex:       "int",
	}
	arg1 := api.ArgSpec{
		Name:        "arg1",
		Description: "this is the 2nd positional argument and must be a lower alphanumerical with dash",
		Regex:       "lower_alpha_dash",
	}

	flag0 := api.FlagSpec{
		Name:        "flag0",
		Description: "this is the flag0",
		Short:       "f",
		Long:        "flag0",
		PArgs:       api.ApiPSpec{&arg0, &arg1},
	}

	flag1 := api.FlagSpec{
		Name:        "flag1",
		Description: "this is the flag1",
		Short:       "F",
		Long:        "flag1",
		PArgs:       api.ApiPSpec{&arg0, &arg1},
	}
	cmd := api.Cmd{
		Name:        "sub",
		Description: "this is the sub command",
		PArgs:       api.ApiPSpec{&arg0, &arg1},
		Flags:       api.ApiFSpec{&flag0, &flag1},
		F: func(a api.Args, kw api.KWArgs) {
			fmt.Printf("This a little function that simply returns the positional arguments and flag options\n")
			fmt.Println("\tFunc:\tExpects 2 positional argument")
			fmt.Println("\tFlags:\tExpects 2 positional arguments")
			fmt.Printf("\tGot:\n")
			fmt.Printf("\t\tPArgs:\t%v\n", a)
			fmt.Printf("\t\tFlags:\t%v\n", kw)
		},
	}

	scmd0 := api.SCmd{
		Name:        "super0",
		Description: "Contains sub command and super command",
		Cmd:         []*api.Cmd{&cmd},
	}

	scmd1 := api.SCmd{
		Name:        "super1",
		Description: "Contains sub command and super command",
		SCmd:        []*api.SCmd{&scmd0},
		Cmd:         []*api.Cmd{&cmd},
	}

	cli := api.CLI{
		Name:        "test",
		Description: "testCLI",
		Version:     "0.1.0",
		SCmd:        []*api.SCmd{&scmd0, &scmd1},
		Cmd:         []*api.Cmd{&cmd},
	}
	return &cli
}

func main() {
	cli := getCliExample()
	cmd := "super0 sub 43 asd-asd --flag0 0 iolo -F 1 yolo"
	fmt.Printf("Running CLI with following arguments: %s\n\n", cmd)
	cli.Run(cmd)
}
