package httpx

import (
	"bytes"
	"encoding/json"
)

func (h *httpx) Body(data []byte) *request {
	h.body = bytes.NewBuffer(data)
	return &request{
		httpx: h,
	}
}

func (h *httpx) Content(data string) *request {
	h.body = bytes.NewBuffer([]byte(data))
	return &request{
		httpx: h,
	}
}

func (h *httpx) StructWithError(data interface{}) (*request, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	h.body = bytes.NewBuffer([]byte(jsonData))

	return &request{
		httpx: h,
	}, nil
}

func (h *httpx) Struct(data interface{}) *request {
	jsonData, _ := json.Marshal(data)
	h.body = bytes.NewBuffer([]byte(jsonData))

	return &request{
		httpx: h,
	}
}
