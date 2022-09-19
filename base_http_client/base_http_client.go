package base_http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/best-expendables-v2/logger"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"moul.io/http2curl"
)

const (
	defaultRetry = 4
	defaultDelay = 10 * time.Second
)

type baseHttpClient struct {
	httpClient   *http.Client
	isReturnCURL bool
	defaultRetry int
	defaultDelay time.Duration
}

type BaseHttpClient interface {
	SetReturnCURL(isReturned bool)
	SetDefaultRetry(defaultRetry int)
	SetDefaultDelay(defaultDelay time.Duration)
	SendRequest(ctx context.Context, method string, url string, options, payload interface{}, headers map[string]string) (*HttpResponse, error)
	SendRequestWithAttempt(ctx context.Context, method string, url string, options, payload interface{}, headers map[string]string) (*HttpResponse, error)
}

func NewBaseHttpClient(httpClient *http.Client) BaseHttpClient {
	return &baseHttpClient{
		httpClient:   httpClient,
		defaultRetry: defaultRetry,
		defaultDelay: defaultDelay,
	}
}

type HttpResponse struct {
	Request    *http.Request `json:"request"`
	StatusCode int           `json:"statusCode"`
	Body       []byte        `json:"body"`
}

func (s *baseHttpClient) SetReturnCURL(isReturnCURL bool) {
	s.isReturnCURL = isReturnCURL
}

func (s *baseHttpClient) SetDefaultRetry(defaultRetry int) {
	s.defaultRetry = defaultRetry
}

func (s *baseHttpClient) SetDefaultDelay(defaultDelay time.Duration) {
	s.defaultDelay = defaultDelay
}

func (s baseHttpClient) SendRequest(ctx context.Context, method string, url string, options, payload interface{}, headers map[string]string) (*HttpResponse, error) {
	requestBody, err := s.processPayload(ctx, payload, headers)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can not get data by endpoint '%s'", url))
	}
	request = request.WithContext(ctx)
	for key, val := range headers {
		request.Header.Add(key, val)
	}
	if options != nil {
		optionsQuery, err := query.Values(options)
		if err != nil {
			return nil, err
		}
		for k, values := range request.URL.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		request.URL.RawQuery = optionsQuery.Encode()
	}
	output := &HttpResponse{
		Request: request,
	}
	s.returnCURL(request)
	resp, err := s.httpClient.Do(request)
	if err != nil {
		return output, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	output.StatusCode = resp.StatusCode
	output.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return output, errors.Wrap(err, "can not read body")
	}
	if resp.StatusCode >= 400 {
		return output, errors.Errorf("can not get data from %s api: %s", url, string(output.Body))
	}
	return output, nil
}

func (s baseHttpClient) returnCURL(request *http.Request) {
	if s.isReturnCURL {
		command, _ := http2curl.GetCurlCommand(request)
		fmt.Println(command)
	}
}

func (s baseHttpClient) SendRequestWithAttempt(ctx context.Context, method string, url string, options, payload interface{}, headers map[string]string) (*HttpResponse, error) {
	resp, err := s.SendRequest(ctx, method, url, options, payload, headers)
	if err != nil || resp.StatusCode >= 400 {
		for i := 0; i <= s.defaultRetry; i++ {
			logger.Error(errors.Wrap(err, "Retrying the request"))
			resp, err = s.SendRequest(ctx, method, url, options, payload, headers)
			if err == nil && resp.StatusCode < 400 {
				break
			}
			time.Sleep(s.defaultDelay)
		}
	}
	return resp, err
}

func (s baseHttpClient) processPayload(ctx context.Context, payload interface{}, headers map[string]string) (io.Reader, error) {
	if payload == nil {
		return bytes.NewBuffer(nil), nil
	}
	var output io.Reader
	for contentType, value := range headers {
		if strings.ToLower(contentType) == "content-type" {
			if strings.ToLower(value) == "application/x-www-form-urlencoded" {
				data, err := query.Values(payload)
				if err != nil {
					return nil, err
				}
				encodedData := data.Encode()
				output = strings.NewReader(encodedData)
				return output, nil
			}
		}
	}
	bPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	output = bytes.NewBuffer(bPayload)
	return output, nil
}
