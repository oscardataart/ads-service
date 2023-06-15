package datasources

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type client struct {
	httpClient http.Client
}

func NewRestClient() HTTPClient {
	return &client{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (client *client) Request(ctx context.Context, info *RequestInfo) (*ResponseInfo, error) {
	var msgDataReader strings.Reader

	request, err := http.NewRequest(info.HTTPMethod, info.Url, &msgDataReader)
	if err != nil {
		log.Printf("error creating request %v", err)
		return &ResponseInfo{}, err
	}

	response, err := client.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		log.Printf("error performing request %v", err)
		return &ResponseInfo{}, BadGateway
	}
	if isErrorStatus(response) {
		return &ResponseInfo{StatusCode: response.StatusCode},
			errors.New(fmt.Sprintf("Error performing rest request %v", response))
	}

	data, err := io.ReadAll(response.Body)

	return &ResponseInfo{Body: data, StatusCode: response.StatusCode}, err
}

func isErrorStatus(resp *http.Response) bool {
	return resp.StatusCode != http.StatusOK
}
