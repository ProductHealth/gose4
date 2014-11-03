package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ProductHealth/gose4"
	"io/ioutil"
	"net/http"
)

type ClientConfiguration struct {
	Host string
	Port int
}

type HttpClient struct {
	config     ClientConfiguration
	httpClient http.Client
}

func NewClient(config ClientConfiguration) Client {
	client := http.Client{}
	return &HttpClient{config, client}
}

func (client *HttpClient) Healthcheck() (*gose4.TestResults, error) {
	requestUrl := fmt.Sprintf("http://%v:%v/service/healthcheck", client.config.Host, client.config.Port)
	var resp, err = http.Get(requestUrl)
	switch {
	case err != nil:
		return nil, err
	case resp.StatusCode != 200:
		return nil, errors.New(resp.Status)
	default:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := gose4.TestResults{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		} else {
			return &result, nil

		}
	}
}
