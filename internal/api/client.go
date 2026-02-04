package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tls-checker/internal/model"
)

const (
	SSLLABS_API = "https://api.ssllabs.com/api/v2"
)

func AnalyzeHost(hostAdress string) (*model.Host, error) {
	// Query to make HTTP request
	req := fmt.Sprintf("%v/analyze?host=%v", SSLLABS_API, hostAdress)

	res, err := http.Get(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
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