package decorator

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

type MetricsClient interface {
	Inc(key string, value int)
}

type queryMetricsDecorator[C, R any] struct {
	client MetricsClient
	base   QueryHandler[C, R]
}

// implements QueryHandler interface
func (q queryMetricsDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	start := time.Now()
	actionName := generateActionName(cmd)
	defer func() {
		end := time.Since(start)
		q.client.Inc(fmt.Sprintf("querts.%s.duration", actionName), int(end.Seconds()))
		if err != nil {
			q.client.Inc(fmt.Sprintf("querts.%s.success", actionName), 1)
		} else {
			q.client.Inc(fmt.Sprintf("querts.%s.failure", actionName), 1)
		}
	}()

	return q.base.Handle(ctx, cmd)
}

//func generateActionName(cmd any) string {
//	return strings.Split(fmt.Sprintf("%#T", cmd), ".")[1]
//}
