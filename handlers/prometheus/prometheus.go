package prometheus

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

type PrometheusHandler struct {
	metricsEndpoint *string
	metricsPort     *string
}

var opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
	Name: "roa_some_counter",
	Help: "Description of this counter",
})

// TODO: Change this mock to real metrics
func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func NewPrometheusHandler() *PrometheusHandler {
	endpoint := viper.GetString("prometheusEndpoint")
	port := ":" + strconv.Itoa(viper.GetInt("prometheusEndpointPort"))

	return &PrometheusHandler{metricsEndpoint: &endpoint, metricsPort: &port}
}

func (prometheusHandler *PrometheusHandler) ExposeMetrics() {

	recordMetrics()

	http.Handle(*prometheusHandler.metricsEndpoint, promhttp.Handler())
	http.ListenAndServe(*prometheusHandler.metricsPort, nil)
}
