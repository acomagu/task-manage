package main

type Recode struct {
	Done bool `json:"done"`
	Data Data `json:"data"`
}

func (f *Recode) Help() string {
	return "task-manage recode"
}

type TaskList []string

func (list TaskList) GetRecodeList(recode []Recode, s bool, path string) []Recode {
	for _, v := range list {
		task := FindTask(path + v)
		recode = append(recode, Recode{
			s,
			task,
		})
	}
	return recode
}

func (f *Recode) Run(args []string) int {
	haveList := root.GetListHave()
	finishedList := root.GetListFinished()
	var recode []Recode
	recode = finishedList.GetRecodeList(recode, true, root.finished)
	recode = haveList.GetRecodeList(recode, false, root.have)
	creatore.Recode(recode)
	return 0
}

func (f *Recode) Synopsis() string {
	return "recode tasks"
}
