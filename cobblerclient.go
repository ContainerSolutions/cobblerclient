/*
Copyright 2015 Container Solutions

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cobblerclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kolo/xmlrpc"
)

const bodyTypeXML = "text/xml"

type HTTPClient interface {
	Post(string, string, io.Reader) (*http.Response, error)
}

type Client struct {
	httpClient HTTPClient
	config     ClientConfig
	token      string
}

type ClientConfig struct {
	Url      string
	Username string
	Password string
}

func NewClient(httpClient HTTPClient, c ClientConfig) Client {
	return Client{
		httpClient: httpClient,
		config:     c,
	}
}

func (c *Client) Call(method string, args ...interface{}) (interface{}, error) {
	var result interface{}

	reqBody, err := xmlrpc.EncodeMethodCall(method, args...)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Post(c.config.Url, bodyTypeXML, bytes.NewReader([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := xmlrpc.NewResponse(body)
	if err := resp.Unmarshal(&result); err != nil {
		return nil, err
	}

	if resp.Failed() {
		return nil, resp.Err()
	}

	return result, nil
}

// Performs a login request to Cobbler using the credentials provided
// in the configuration in the initializer.
func (c *Client) Login() (bool, error) {
	var result interface{}
	result, err := c.Call("login", c.config.Username, c.config.Password)
	if err != nil {
		return false, err
	}

	c.token = result.(string)
	return true, nil
}

// Sync the system.
// Returns true if the sync was successful, or false if it was not.
// Returns an error if anything went wrong
func (c *Client) Sync() (bool, error) {
	_, err := c.Call("sync", c.token)
	if err != nil {
		return false, err
	}

	return true, nil
}
