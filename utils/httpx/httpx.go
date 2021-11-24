package httpx

import (
	"io"
	"log"
	"net/http"
)

type httpx struct {
	method      string
	url         string
	header      map[string]string
	body        io.Reader
	contentType string
	httpClient  *http.Client
	debug       bool
}

func Get(url string) *httpx {
	return newClient(url, "GET")
}

func Post(url string) *httpx {
	return newClient(url, "POST")
}

func newClient(url string, method string) *httpx {
	return &httpx{
		method:      method,
		url:         url,
		contentType: "application/json",
		header:      make(map[string]string),
		httpClient:  &http.Client{},
	}
}

func (h *httpx) Debug() *httpx {
	h.debug = true
	return h
}

func (h *httpx) URL() string {
	return h.url
}

func (h *httpx) Request() (*response, error) {
	r := &request{
		httpx: h,
	}

	return r.Request()
}

func (h *httpx) log(format string, v ...interface{}) {
	if h.debug {
		log.Printf("[httpx] "+format+"\n", v...)
	}
}
