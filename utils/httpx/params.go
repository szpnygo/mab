package httpx

import "net/url"

func (h *httpx) Params(params map[string]string) *httpx {
	base, _ := url.Parse(h.url)
	p := base.Query()
	for k, v := range params {
		p.Set(k, v)
	}
	base.RawQuery = p.Encode()
	h.url = base.String()
	return h
}

func (h *httpx) AddParam(key, value string) *httpx {
	return h.Params(map[string]string{key: value})
}
