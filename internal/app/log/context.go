package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

func (l *commonLogger) InfoCtx(ctx context.Context, msg string, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithExtras(fields)).Info(msg)
}

func (l *commonLogger) TraceCtx(ctx context.Context, msg string, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithExtras(fields)).Trace(msg)
}

func (l *commonLogger) WarnCtx(ctx context.Context, msg string, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithExtras(fields)).Warn(msg)
}

func (l *commonLogger) DebugCtx(ctx context.Context, msg string, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithExtras(fields)).Debug(msg)
}

func (l *commonLogger) ErrorMessageCtx(ctx context.Context, msg string, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithException(&ExceptionData{
		Message: msg,
	})).Error()
}

func (l *commonLogger) ErrorCtx(ctx context.Context, err error, fields ...interface{}) {

	l.setFieldsWithContext(ctx, WithException(&ExceptionData{
		Type:       fmt.Sprintf("%T", err),
		Message:    err.Error(),
		Stacktrace: l.formatLogStack(string(debug.Stack())),
	})).Error()
}

func (l *commonLogger) FatalCtx(ctx context.Context, err error, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithException(&ExceptionData{
		Type:       fmt.Sprintf("%T", err),
		Message:    err.Error(),
		Stacktrace: l.formatLogStack(string(debug.Stack())),
	})).Fatal()
}

func (l *commonLogger) PanicCtx(ctx context.Context, err error, fields ...interface{}) {
	l.setFieldsWithContext(ctx, WithException(&ExceptionData{
		Type:       fmt.Sprintf("%T", err),
		Message:    err.Error(),
		Stacktrace: l.formatLogStack(string(debug.Stack())),
	})).Panic()
}

func (l *commonLogger) SetOptionsToCtx(ctx context.Context, optValues ...LoggerOption) context.Context {
	opts := &LoggerOptions{}
	optsFromCtx := l.OptionsFromCtx(ctx)
	if optsFromCtx != nil {
		opts = optsFromCtx
	}
	for _, setOption := range optValues {
		setOption(opts)
	}
	return context.WithValue(ctx, contextKey(LogOptionsContextKey), *opts)
}

func (l *commonLogger) OptionsFromCtx(ctx context.Context) *LoggerOptions {
	if ctx == nil {
		return nil
	}
	opts, ok := ctx.Value(contextKey(LogOptionsContextKey)).(LoggerOptions)
	if ok {
		return &opts
	}
	return nil
}

func (l *commonLogger) setFieldsWithContext(ctx context.Context, optValues ...LoggerOption) *logrus.Entry {
	opts := &LoggerOptions{}
	optsFromCtx := l.OptionsFromCtx(ctx)
	if optsFromCtx != nil {
		opts = optsFromCtx
	}

	for _, setOption := range optValues {
		setOption(opts)
	}
	return l.Console.WithField(LogOptionsField, *opts)
}
