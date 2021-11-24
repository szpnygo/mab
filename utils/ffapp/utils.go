package ffapp

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostRaw(apiUrl string, data []byte, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	msgResult, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func Get(apiUrl string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
