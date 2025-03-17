package tea_task

import "testing"

func ProcessTeaTest(t *testing.T) {
	_ = NewCapOfTea(HotWatter{}, Sugar{}, Tea{})
	ProcessTea()
	//TODO: ожидать 10 чашек, например
}
