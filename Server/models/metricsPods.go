package models

type MetricsNodes struct {
	NodeName string  `json:"node_name"`
	Cpu      float64 `json:"cpu"`
	Memory   float64 `json:"memory"`
}

type MetricsPods struct {
	PodName   string  `json:"pod_name"`
	Namespace string  `json:"namespace"`
	Container string  `json:"container"`
	Cpu       float64 `json:"cpu"`
	Memory    float64 `json:"memory"`
	Size      string  `json:"size"`
}

type Metrics struct {
	Nodes []MetricsNodes `json:"nodes"`
	Pods  []MetricsPods  `json:"pods"`
}

var (
	metricsNodes []MetricsNodes
	metricsPods  []MetricsPods
	metrics      []Metrics
)

func GetMetricsNodes() []MetricsNodes {
	return metricsNodes
}

func SetMetricsNods(MetricsNodes []MetricsNodes) {
	metricsNodes = MetricsNodes
}

func GetMetricsPods() []MetricsPods {
	return metricsPods
}

func SetMetricsPods(MetricsPods []MetricsPods) {
	metricsPods = MetricsPods
}

func GetMetrics() []Metrics {
	return metrics
}

func SetMetrics(Metrics []Metrics) {
	metrics = Metrics
}
