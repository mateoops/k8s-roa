package prometheus

type NodeMetrics struct {
	Name string
}

type NodeUsageMetrics struct {
	CpuUsage    int64
	MemoryUsage int64
	Node        NodeMetrics
}
