package httpx

import (
	"net/http"
)

type request struct {
	*httpx
}

func (r *request) Request() (*response, error) {
	var req *http.Request
	var err error
	if r.method == "GET" {
		req, err = http.NewRequest(r.method, r.url, nil)
	} else {
		req, err = http.NewRequest(r.method, r.url, r.body)
	}
	if err != nil {
		r.log("get the error %s %s", r.method, err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", r.contentType)
	for key, value := range r.header {
		req.Header.Add(key, value)
	}

	r.log("%s %s", r.method, r.url)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		r.log("get the error %s %s", r.method, err.Error())
		return nil, err
	}

	r.log("%s success", r.method)

	return &response{
		httpx:    r.httpx,
		response: resp,
	}, nil
}

func (r *request) String() (string, error) {
	respone, err := r.Request()
	if err != nil {
		return "", err
	}
	return respone.String()
}

func (r *request) Bytes() ([]byte, error) {
	respone, err := r.Request()
	if err != nil {
		return nil, err
	}
	return respone.Bytes()
}

func (r *request) JSON(data interface{}) error {
	respone, err := r.Request()
	if err != nil {
		return err
	}
	return respone.JSON(data)
}
