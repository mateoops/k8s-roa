package prometheus

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO: its'a only a test

var opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	Name: "roa_some_counter",
	Help: "Description of this counter",
})

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func ExposePrometheusMetrics(endpoint string, port int) {

	recordMetrics()

	portAsString := ":" + strconv.Itoa(port)

	http.Handle(endpoint, promhttp.Handler())
	http.ListenAndServe(portAsString, nil)
}
