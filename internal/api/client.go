package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"tls-checker/internal/model"
)

const (
	SSLLabsApi = "https://api.ssllabs.com/api/v2"
)

type ServiceError struct {
	Code    int
	Message string
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

// Make the API call to SSLLabs to retrieve host information
func AnalyzeHost(query *model.AnalyzeHostQuery) (*model.Host, error) {
	// Query parameters to make HTTP request
	params := url.Values{}
	params.Add("host", query.Host)
	params.Add("publish", query.Publish)
	params.Add("startNew", query.StartNew)
	params.Add("fromCache", query.FromCache)
	params.Add("maxAge", query.MaxAge)
	params.Add("all", query.All)
	params.Add("ignoreMismatch", query.IgnoreMismatch)

	// Build the string request
	req := fmt.Sprintf("%v/analyze?%v", SSLLabsApi, params.Encode())

	res, err := http.Get(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Return an error if the request was successful but the response was an error
	if res.StatusCode != 200 {
		myerr := ServiceError{
			Code:    res.StatusCode,
			Message: string(body),
		}
		return nil, myerr
	}

	defer res.Body.Close()

	var host model.Host
	json.Unmarshal(body, &host)
	return &host, nil
}
