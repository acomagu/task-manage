package main

type Recode struct{}

func (f *Recode) Help() string {
	return "app Recode"
}

func (f *Recode) Run(args []string) int {

	return 0
}

func (f *Recode) Synopsis() string {
	return "display of all tasks"
}
