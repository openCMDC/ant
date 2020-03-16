package core

type AntContext interface {
	TargetProcessID() string
}

type DefaultCtx struct {
	targetId string
}

func (d DefaultCtx) TargetProcessID() string {
	return d.targetId
}

func NewDefaultAntCtx(pid string) AntContext {
	ctx := new(DefaultCtx)
	ctx.targetId = pid
	return ctx
}
