package main

import (
	"github.com/RangelReale/gostatsd/statsd"
	"log"
	"time"
)

func main() {
	log.Printf("Starting...")

	fagg := func(m statsd.MetricMap) {
		log.Printf("%s", m)
	}

	aggregator := statsd.NewMetricAggregator(statsd.MetricSenderFunc(fagg), 10*time.Second)
	go aggregator.Aggregate()

	f := func(m statsd.Metric) {
		log.Printf("%s", m)
		aggregator.MetricChan <- m
	}
	r := statsd.MetricReceiver{":8125", statsd.HandlerFunc(f)}
	r.ListenAndReceive()
}
