package main

import (
    "fmt"
	"os"
	"strings"
	"net/http"
	"time"
)

func main() {
	endpoints := getEndpoints()
	c := make(chan endpointStatus)
	for _, endpoint := range endpoints {
		go checkEndpoint(endpoint, c)
	}

	result := make([]endpointStatus, len(endpoints))
	for i, _ := range endpoints {
		result[i] = <-c
		fmt.Printf("%s, %s, %d, %s\n", result[i].time, result[i].url, result[i].statuscode, result[i].error)
	}
}

func getEndpoints() []string {
	s := strings.Split(os.Getenv("SERVICE_TESTER_ENDPOINTS"), ",")
	for i := range s {
		s[i] = strings.TrimSpace(s[i])
	}
	return s
}

func checkEndpoint(url string, c chan endpointStatus) {
	res, err := http.Get(url)
	if err != nil {
		c <- endpointStatus{url, 999, time.Now(), err.Error()}
		return
	}
	c <- endpointStatus{url, res.StatusCode, time.Now(), ""}
}

type endpointStatus struct {
	url        string
	statuscode int
	time       time.Time
	error      string
}