package server

func (h *HTTPServer) Stop() error {
	return h.stop()
}
