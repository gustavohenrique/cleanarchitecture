package metrics

import (
	"encoding/json"

	"github.com/gustavohenrique/gometrics"
)

var collector = gometrics.New()

func Collect() string {
	metrics, _ := collector.Metrics() // collector.Process() or collector.Docker()
	bytes, _ := json.Marshal(metrics)
	return string(bytes)
}

func CpuUsagePercentage() float64 {
	metrics, _ := collector.Metrics()
	return metrics.CpuUsagePercentage
}
