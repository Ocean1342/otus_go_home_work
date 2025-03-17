package tea_task

/* задача приготовить чай */

type CapOfTea struct {
	w HotWatter
	s Sugar
	t Tea
}

type HotWatter struct{}
type Sugar struct{}
type Tea struct{}

func NewCapOfTea(w HotWatter, s Sugar, t Tea) CapOfTea {
	return CapOfTea{w, s, t}
}

func ProcessTea() []CapOfTea {

	return []CapOfTea{}
}
