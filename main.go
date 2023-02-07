package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewBuildInfoCollector(),
	)

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	randomDuration := promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
		Name: "random_duration_seconds",
		Help: "a random histogram representing http request latencies",
	})

	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ticker.C:
				randomDuration.(prometheus.ExemplarObserver).ObserveWithExemplar(
					float64(rnd.Intn(100)),
					prometheus.Labels{
						"UUID": uuid.NewString(),
					},
				)
			}
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{EnableOpenMetrics: true}))
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
