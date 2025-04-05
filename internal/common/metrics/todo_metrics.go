package metrics

type TodoMetrics struct {
}

// implement metricsClient interface
func (t TodoMetrics) Inc(key string, value int) {
	//TODO implement me
	panic("implement me")
}

// New todoMetrics (metricsClient)
func NewTodoMetrics() TodoMetrics {
	return TodoMetrics{}
}
