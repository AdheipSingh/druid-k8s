package collector

import (
	"druid-exporter/utils"

	"github.com/prometheus/client_golang/prometheus"
)

// metricCollector includes the list of metrics
type metricCollector struct {
	RunningTaskMetric   *prometheus.Desc
	CompletedTaskMetric *prometheus.Desc
	WaitingTaskMetric   *prometheus.Desc
	PendingTaskMetric   *prometheus.Desc
}

// Collector return the defined metrics with prometheus description
func Collector() *metricCollector {
	return &metricCollector{
		RunningTaskMetric: prometheus.NewDesc("druid_running_tasks",
			"Shows number of running tasks",
			nil, prometheus.Labels{
				"running": "tasks",
			},
		),
		CompletedTaskMetric: prometheus.NewDesc("druid_completed_tasks",
			"Shows number of Completed tasks",
			nil, prometheus.Labels{
				"compcleted": "tascks",
			},
		),
		WaitingTaskMetric: prometheus.NewDesc("druid_waiting_tasks",
			"number of Completed tasks",
			nil, prometheus.Labels{
				"waiting": "tasks",
			},
		),
		PendingTaskMetric: prometheus.NewDesc("druid_pending_tasks",
			"number of Completed tasks",
			nil, prometheus.Labels{
				"pending": "tasks",
			},
		),
	}
}

// Describe method shall ingest the metric value passed.
func (collector *metricCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.RunningTaskMetric
	ch <- collector.CompletedTaskMetric
	ch <- collector.PendingTaskMetric
	ch <- collector.WaitingTaskMetric

}

// Collect ingests
func (collector *metricCollector) Collect(ch chan<- prometheus.Metric) {

	running := utils.HTTPGetMetric("http://localhost:8081/druid/indexer/v1/runningTasks")
	runningTasks := len(running)
	ch <- prometheus.MustNewConstMetric(collector.RunningTaskMetric, prometheus.CounterValue, float64(runningTasks))

	completed := utils.HTTPGetMetric("http://localhost:8081/druid/indexer/v1/completeTasks")
	completedTasks := len(completed)
	ch <- prometheus.MustNewConstMetric(collector.CompletedTaskMetric, prometheus.GaugeValue, float64(completedTasks))

	waiting := utils.HTTPGetMetric("http://localhost:8081/druid/indexer/v1/waitingTasks")
	waitingTasks := len(waiting)
	ch <- prometheus.MustNewConstMetric(collector.WaitingTaskMetric, prometheus.GaugeValue, float64(waitingTasks))

	pending := utils.HTTPGetMetric("http://localhost:8081/druid/indexer/v1/pendingTasks")
	pendingTasks := len(pending)
	ch <- prometheus.MustNewConstMetric(collector.PendingTaskMetric, prometheus.GaugeValue, float64(pendingTasks))

}
