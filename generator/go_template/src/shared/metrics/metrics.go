package metrics

import (
	"encoding/json"

	"github.com/gustavohenrique/gometrics"
)

func Collect() string {
	collector := gometrics.New()
	metrics, _ := collector.Metrics() // collector.Process() or collector.Docker()
	bytes, _ := json.Marshal(metrics)
	return string(bytes)
}
