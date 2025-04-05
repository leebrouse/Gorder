package metrics

import "log"

type TodoMetrics struct {
}

// implement metricsClient interface
func (t TodoMetrics) Inc(key string, value int) {
	//TODO implement me
	log.Print("To do Metrics logger")
}

// New todoMetrics (metricsClient)
func NewTodoMetrics() TodoMetrics {
	return TodoMetrics{}
}
