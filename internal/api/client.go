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

	req := fmt.Sprintf("%v/analyze?%v", SSLLabsApi, params.Encode())

	res, err := http.Get(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

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

// func GetEndpointData(hostAdress string, ipAddress string) (*model.Endpoint, error) {
// 	query := fmt.Sprintf("%v/getEndpointData?host=%v&s=%v", SSLLABS_API, hostAdress, ipAddress)

// 	response, error := http.Get(query)
// 	if error != nil {
// 		return nil, error
// 	}

// 	body, error := io.ReadAll(response.Body)
// 	if error != nil {
// 		return nil, error
// 	}
// 	defer response.Body.Close()

// 	var host model.Endpoint
// 	json.Unmarshal(body, &host)
// 	return &host, nil
// }
