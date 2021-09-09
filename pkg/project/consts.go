package project

import "context"

type TraceContext string

const (
	TraceID TraceContext = "TRACE_ID"
)

func TraceFromCtx(ctx context.Context) string {
	id, ok := ctx.Value(TraceID).(string)
	if ok {
		return id
	}
	return ""
}
