package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingHandler struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           http.Handler
}

type statusCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *statusCodeWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func InstrumentingMiddleware(next http.Handler, counter metrics.Counter, latency metrics.Histogram) http.Handler {
	return &instrumentingHandler{
		requestCount:   counter,
		requestLatency: latency,
		next:           next,
	}
}

func (h *instrumentingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	scw := &statusCodeWriter{rw, http.StatusOK}

	defer func(start time.Time) {
		h.requestCount.With("method", r.Method, "path", r.URL.Path, "status_code", strconv.Itoa(scw.statusCode)).Add(1)
		h.requestLatency.With("method", r.Method, "path", r.URL.Path).Observe(time.Since(start).Seconds())
	}(time.Now())

	h.next.ServeHTTP(scw, r)
}
