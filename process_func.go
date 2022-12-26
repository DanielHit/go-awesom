package service

import "context"

type SettleResult struct{}

type AppADSettleImpl struct {
	// func处理过程
	ProgressFun []func(ctx context.Context) error
	// 数据源 input
	AppADDatasource
	// 结果 result
	SettleResult
}

func (a AppADSettleImpl) CheckSettleParam(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AppADSettleImpl) Settle(ctx context.Context) {
	for _, f := range a.ProgressFun {
		f(ctx)
	}
}

func InitProcessFun() {
	processor := &AppADSettleImpl{}
	// 真正的函数编程
	processor.ProgressFun = []func(ctx context.Context) error{
		processor.GetIncome,
		processor.GetIncome,
	}
}

// 构造器
type AppADDatasource struct {
	// 原始收入
	OriginAmount int64
}

func (process *AppADDatasource) GetIncome(ctx context.Context) error {
	return nil
}

func (process *AppADDatasource) CalculateOriginIncome(ctx context.Context) error {
	return nil
}
