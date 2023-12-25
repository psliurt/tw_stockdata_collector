package openapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type ApiRequestWrapper struct {
	sender      *ApiSender
	method      string
	url         string
	headers     map[string]string
	statusCode  int
	queryParams map[string]string
}

func (ar *ApiRequestWrapper) Method(m string) *ApiRequestWrapper {
	ar.method = m
	return ar
}

func (ar *ApiRequestWrapper) Url(u string) *ApiRequestWrapper {
	ar.url = u
	return ar
}

func (ar *ApiRequestWrapper) AddHeader(header, value string) *ApiRequestWrapper {
	ar.headers[header] = value
	return ar
}

func (ar *ApiRequestWrapper) AddQueryParam(param, value string) *ApiRequestWrapper {
	ar.queryParams[param] = value
	return ar
}

func (ar *ApiRequestWrapper) StatusCode() int {
	return ar.statusCode
}

func (ar *ApiRequestWrapper) SendJson(obj any) ([]byte, error) {

	bs, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}

	req, err := http.NewRequest(ar.method, ar.url, bytes.NewReader(bs))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range ar.headers {
		req.Header.Set(k, v)
	}
	rsp, err := ar.sender.client.Do(req)
	if err != nil {
		panic(err)
	}
	ar.statusCode = rsp.StatusCode
	defer rsp.Body.Close()
	// read body
	return io.ReadAll(rsp.Body)
}

func (ar *ApiRequestWrapper) SendQuery() ([]byte, error) {
	baseURL, _ := url.Parse(ar.url)
	params := url.Values{}
	for p, q := range ar.queryParams {
		params.Add(p, q)
	}

	baseURL.RawQuery = params.Encode()
	req, err := http.NewRequest(ar.method, baseURL.String(), nil)
	if err != nil {
		panic(err)
	}

	for k, v := range ar.headers {
		req.Header.Set(k, v)
	}

	rsp, err := ar.sender.client.Do(req)
	if err != nil {
		panic(err)
	}
	ar.statusCode = rsp.StatusCode
	defer rsp.Body.Close()
	// read body
	return io.ReadAll(rsp.Body)
}

var apiInstance *ApiSender
var createApiSenderOnce sync.Once

type ApiSender struct {
	client *http.Client
}

func Initialize() {
	createApiSenderOnce.Do(func() {
		apiInstance = createApiSender()
	})
}

func Instance() *ApiSender {
	Initialize()
	return apiInstance
}

func createApiSender() *ApiSender {
	c := &http.Client{
		Timeout: 30 * time.Second,
	}
	return &ApiSender{
		client: c,
	}
}

func (a *ApiSender) CreateRequest() *ApiRequestWrapper {
	return &ApiRequestWrapper{
		sender:      a,
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
	}
}
