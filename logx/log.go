package logx

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ContextMessage = "CtxMessage"
)

// WithLogContext 构建带message的context
func WithLogContext(ctx context.Context, val *CtxMessage) context.Context {
	return context.WithValue(ctx, ContextMessage, val)
}

// GetCtxMessageByLogContext 通过context获取ctxMessage
func GetCtxMessageByLogContext(ctx context.Context) *CtxMessage {
	value := ctx.Value(ContextMessage)
	msg, ok := value.(*CtxMessage)
	if !ok {
		return &CtxMessage{}
	}

	return msg
}

// ErrorContextf 打印错误级别日志
func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	m := GetCtxMessageByLogContext(ctx)
	fields := m.genLogFields()
	logx.Errorw(err.Error(), fields...)
}

// InfoContextf 打印info级别日志
func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	m := GetCtxMessageByLogContext(ctx)
	fields := m.genLogFields()
	logx.Infow(msg, fields...)
}
