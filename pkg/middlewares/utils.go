package middlewares

import "net/http"

func requestInfo(req *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"path":   req.URL.Path,
		"query":  req.URL.RawQuery,
		"method": req.Method,
		"client": req.RemoteAddr,
	}
}

type wrappedWriter struct {
	http.ResponseWriter

	wroteStatus int
}

func (wr *wrappedWriter) WriteHeader(statusCode int) {
	wr.wroteStatus = statusCode
	wr.ResponseWriter.WriteHeader(statusCode)
}