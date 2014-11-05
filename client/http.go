package client

import (
	"encoding/json"
	"errors"
	"fmt"
	gnet "github.com/ProductHealth/gommons/net"
	"github.com/ProductHealth/gose4"
	"io/ioutil"
	"net/http"
)

type httpClient struct {
	endpoint gnet.Endpoint
}

func New(endpoint gnet.Endpoint) Client {
	return &httpClient{endpoint}
}

func (client *httpClient) Healthcheck() (*gose4.TestResults, error) {
	requestUrl := fmt.Sprintf("http://%v:%v/service/healthcheck", client.endpoint.HostName(), client.endpoint.Port())
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
