package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	httpAPIScheme      = "https"
	httpAPIHost        = "api.bitflyer.jp"
	httpAPIVersionPath = "/v1"
)

// Client is a client.
type Client struct {
	httpclient *http.Client
}

// NewClient creates a new client.
func NewClient() *Client {
	return &Client{
		httpclient: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
	}
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(method string, apipath string, query map[string]string, data []byte) ([]byte, error) {
	p := path.Join(httpAPIVersionPath, apipath)
	u := url.URL{Scheme: httpAPIScheme, Host: httpAPIHost, Path: p}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// query
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// header
	for key, value := range map[string]string{"key": "value"} {
		req.Header.Add(key, value)
	}

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rbody, nil
}

type Request struct {
	id int
}

type Response struct {
	id    int
	title string
}

func (c *Client) GetPosts(req *Request) (*Response, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := c.Do(http.MethodPost, "/posts", nil, data)
	if err != nil {
		return nil, err
	}

	var res Response
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func main() {
	c := NewClient()

	res, err := c.GetPosts(&Request{})
	if err != nil {
		log.Printf("%v", err)
	}

	fmt.Printf("%#v\n", res)
}
