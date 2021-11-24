package httpx

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	*httpx
	response *http.Response
}

func (r *response) String() (string, error) {
	defer r.response.Body.Close()
	body, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		r.log("get error %s", err.Error())
		return "", err
	}

	r.log("get response %s", string(body))
	return string(body), nil
}

func (r *response) Bytes() ([]byte, error) {
	defer r.response.Body.Close()
	body, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		r.log("get error %s", err.Error())
		return nil, err
	}

	r.log("get response %s", string(body))
	return body, nil
}

func (r *response) JSON(data interface{}) error {
	defer r.response.Body.Close()
	body, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		r.log("get error %s", err.Error())
		return err
	}

	r.log("get response %s", string(body))

	err = json.Unmarshal(body, &data)
	if err != nil {
		r.log("get error %s", err.Error())
		return err
	}

	return nil
}
