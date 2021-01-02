package core

type AntContext struct {
	TargetProcessID string
}

func NewDefaultAntCtx() *AntContext {
	ctx := new(AntContext)
	return ctx
}
