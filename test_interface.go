package service

// 实现抽象
type SettleProcess interface {
	Run()
	GetName() string
	DoWork(taskName string, lastTimes int) string
}

type GameAdSettleProcess struct {
	Name string
	Age  int
}

type GamePaymentSettleProcess struct {
	Name string
	Age  int
}

func (g GamePaymentSettleProcess) Run() {
	//TODO implement me
	panic("implement me")
}

func (g GamePaymentSettleProcess) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (g GamePaymentSettleProcess) DoWork(taskName string, lastTimes int) string {
	//TODO implement me
	panic("implement me")
}

func (S GameAdSettleProcess) Run() {
	//TODO implement me
	panic("implement me")
}

func (S GameAdSettleProcess) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (S GameAdSettleProcess) DoWork(taskName string, lastTimes int) string {
	//TODO implement me
	panic("implement me")
}
