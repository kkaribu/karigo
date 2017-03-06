package karigo

import "net/http"

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra
// methods about the response.
type ResponseWriter interface {
	http.ResponseWriter

	// Status returns the HTTP status of the response, or 0 if it hasn't been
	// set yet.
	Status() int

	// Written returns wether a status code has been written or not.
	Written() bool

	// Unwrap returns the original proxied target.
	Unwrap() http.ResponseWriter
}

// WrapResponseWriter wraps an http.ResponseWriter and returns a
// ResponseWriter.
func WrapResponseWriter(w http.ResponseWriter) ResponseWriter {
	bw := responseWriter{ResponseWriter: w}
	return &bw
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (r *responseWriter) WriteHeader(status int) {
	if r.status == 0 {
		r.status = status
		r.ResponseWriter.WriteHeader(status)
	}
}

func (r *responseWriter) Write(buf []byte) (int, error) {
	r.WriteHeader(http.StatusOK)
	n, err := r.ResponseWriter.Write(buf)
	return n, err
}

func (r *responseWriter) Status() int {
	return r.status
}

func (r *responseWriter) Written() bool {
	return r.status != 0
}

func (r *responseWriter) Unwrap() http.ResponseWriter {
	return r.ResponseWriter
}
