package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	events = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "demo_events_total", Help: "Demo events"},
		[]string{"type"},
	)
	activePlayers = prometheus.NewGauge(prometheus.GaugeOpts{Name: "demo_active_players", Help: "Active players"})
	reqLatency    = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "demo_request_latency_seconds",
		Help:    "Request latency",
		Buckets: prometheus.DefBuckets,
	})
)

func main() {
	rand.Seed(time.Now().UnixNano())
	prometheus.MustRegister(events, activePlayers, reqLatency)

	http.Handle("/metrics", promhttp.Handler())
	go simulate()

	log.Println("demoapp listening on :2112")
	if err := http.ListenAndServe(":2112", nil); err != nil {
		log.Fatal(err)
	}
}

func simulate() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	types := []string{"move", "chat"}
	for range ticker.C {
		t := types[rand.Intn(len(types))]
		events.WithLabelValues(t).Inc()
		activePlayers.Set(float64(10 + rand.Intn(30)))
		reqLatency.Observe(rand.Float64())
	}
}
