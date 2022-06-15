package main

import (
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"gotest.tools/assert"
)


func checkMetrics(t *testing.T, goldenFile string, metricName string) error {
	f, err := os.Open(goldenFile)
	defer f.Close()
	if err != nil {
		return err
	}
	return testutil.CollectAndCompare(counter, f, metricName)
}

func Test_increaseMetric(t *testing.T) {
	assert.NilError(t, checkMetrics(t, "metrics_before.golden", "test_total"))
	increaseMetric()
	assert.NilError(t, checkMetrics(t, "metrics_after.golden", "test_total"))
}
