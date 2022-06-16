package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "test_total",
	Help: "Testing counter",
})

func increaseMetric() {
	counter.Inc()
}

func main() {
	go func() {
		for {
			increaseMetric()
			time.Sleep(1)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
