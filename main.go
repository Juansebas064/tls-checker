package main

import (
	"fmt"
	"net/http"
	// "os"
	"io"
)

const SSLLABS_API = "https://api.ssllabs.com/api/v2"

func main() {
	fmt.Println(SSLLABS_API)
	analyzeHost("www.ssllabs.com")
}

func analyzeHost(hostAdress string) {
	query := fmt.Sprintf("%v/analyze?host=%v", SSLLABS_API, hostAdress)
	response, error := http.Get(query)

	if error != nil {
		fmt.Println("ERROR")
	}

	body, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println("ERROR")
	}
	fmt.Println(string(body))
}

func programLoop() {
	for {

	}
}