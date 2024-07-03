package prometheus

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

type PrometheusHandler struct {
	endpoint *string
	port     *string
}

func NewPrometheusHandler() *PrometheusHandler {
	endpoint := viper.GetString("prometheusEndpoint")
	port := ":" + strconv.Itoa(viper.GetInt("prometheusEndpointPort"))

	return &PrometheusHandler{endpoint: &endpoint, port: &port}
}

func (prometheusHandler *PrometheusHandler) ExposeMetrics() {

	http.Handle(*prometheusHandler.endpoint, promhttp.Handler())
	http.ListenAndServe(*prometheusHandler.port, nil)
}
