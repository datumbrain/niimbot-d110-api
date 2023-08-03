package main

import (
	"fmt"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func withLogging(wrappedHandler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)

		t1 := time.Now()
		wrappedHandler.ServeHTTP(lrw, r)
		t2 := time.Now()

		fmt.Printf("%s\t[HTTP] %s %s [%s] - %s <- %d %s\n",
			t1.UTC().Format(time.RFC3339), r.Method, r.URL.String(), getTimeDifference(t1, t2), r.RemoteAddr, lrw.statusCode, http.StatusText(lrw.statusCode),
		)
	})
}

func getTimeDifference(t1, t2 time.Time) string {
	dur := t2.Sub(t1)
	if dur < time.Second {
		return fmt.Sprintf("%.6fms", float64(dur.Nanoseconds())/1000000)
	}

	return fmt.Sprintf("%.3fs", float64(dur.Milliseconds())/1000)
}
