package EchoServer

import (
	"context"

	"github.com/opentracing/opentracing-go"
	olog "github.com/opentracing/opentracing-go/log"
)

func example1(ctx context.Context) {

	span, _ := opentracing.StartSpanFromContext(ctx, "example1")
	defer span.Finish()

	span.LogFields(
		olog.String("event", "example1event"),
		olog.String("value", "example1 value"),
	)
	{
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		example2(ctx)
	}
}

func example2(ctx context.Context) {

	span, _ := opentracing.StartSpanFromContext(ctx, "example2")
	defer span.Finish()

	span.LogFields(
		olog.String("event", "example1event"),
		olog.String("value", "example1 value"),
	)
	{
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		example3(ctx)
	}
}

func example3(ctx context.Context) {

	span, _ := opentracing.StartSpanFromContext(ctx, "example3")
	defer span.Finish()

	span.LogFields(
		olog.String("event", "example1event"),
		olog.String("value", "example1 value"),
	)
}
