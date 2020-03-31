package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Log struct {
	EventType    string `json:"eventType"`
	ProjectName  string `json:"projectName"`
	PipelineName string `json:"pipelineName"`
	BuildId      string `json:"buildId"`
	BuildMessage string `json:"message"`
}

func main() {
	jsonPayload := Log{
		EventType:    "Codefresh",
		ProjectName:  "Knapsack Pro Demo",
		PipelineName: "knapsack-pro-demo",
		BuildId:      "1214416",
		BuildMessage: "This is just a test3",
	}
	jsonValue, err := json.Marshal(jsonPayload)
	if err != nil {
		fmt.Printf("Formatting JSON failed with error %s\n", err)
	}

	accountId := "1410943"
	baseUrl := "https://insights-collector.newrelic.com/v1/accounts/"
	eventUrl := "/events"

	extendedUrl := baseUrl + accountId + eventUrl

	var byteBuffer bytes.Buffer
	g := gzip.NewWriter(&byteBuffer)
	g.Write(jsonValue)
	g.Close()

	// request, _ := http.NewRequest("POST", extendedUrl, bytes.NewBuffer(jsonValue))
	request, _ := http.NewRequest("POST", extendedUrl, &byteBuffer)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Encoding", "gzip")
	request.Header.Set("X-Insert-Key", "NRII-v5y7YgtjLAkopJQgAGSPGV-jPriJIgUc")
	client := &http.Client{}
	response, err := client.Do(request)
	// response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
