package service

type ClosFun func(string) string

type IHello interface {
	Say(a, b int) string
	GetName() string
}

type HelloImplA struct{}

func (h HelloImplA) Say(a, b int) string {
	//TODO implement me
	panic("implement me")
}

func (h HelloImplA) GetName() string {
	//TODO implement me
	panic("implement me")
}

type NewStr struct {
	F  ClosFun
	IC IHello
}

func DoFun() {

	f := func(a string) string {
		return a + "fuck"
	}

	a := NewStr{
		F:  f,
		IC: &HelloImplA{},
	}

	// struct build diff
	a.IC.Say(1, 2)
	a.IC.GetName()

	// use func to do it
	a.F("a")
}
