package prometheus

type NodeMetrics struct {
	Name string
}

type PodMetrics struct {
	Name      string
	Namespace string
}

type ContainerMetrics struct {
	Name string
	Pod  PodMetrics
}

type NodeUsageMetrics struct {
	CpuUsage    int64
	MemoryUsage int64
	Node        NodeMetrics
}

type ContainerUsageMetrics struct {
	CpuUsage    int64
	MemoryUsage int64
	Container   ContainerMetrics
}
