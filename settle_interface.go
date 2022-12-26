package service

import "context"

// ISettleProcess 采用函数的编程
type ISettleProcess interface {
	// 校验参数
	CheckSettleParam(ctx context.Context)

	// Settle 结算
	Settle(ctx context.Context)
}
