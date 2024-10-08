package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	endpoints := getEndpoints()
	schedule, err := strconv.Atoi(os.Getenv("SERVICE_TESTER_SCHEDULE_SECONDS"))
	if err != nil {
		fmt.Println("Schedule could not be converted into a integer, check value of SERVICE_TESTER_SCHEDULE_SECONDS", err)
		return
	}
	for {
		c := make(chan endpointStatus)
		for _, endpoint := range endpoints {
			go checkEndpoint(endpoint, c)
		}

		result := make([]endpointStatus, len(endpoints))
		for i, _ := range endpoints {
			result[i] = <-c
			fmt.Printf("%s, %s, %d, %s\n", result[i].time, result[i].url, result[i].statuscode, result[i].error)
		}
		time.Sleep(time.Duration(schedule) * time.Second)
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
