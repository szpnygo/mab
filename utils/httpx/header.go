package httpx

func (h *httpx) Headers(headers map[string]string) *httpx {
	for k, v := range headers {
		h.header[k] = v
	}
	return h
}

func (h *httpx) AddHeader(key, value string) *httpx {
	h.header[key] = value
	return h
}
