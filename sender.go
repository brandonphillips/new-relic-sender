package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli"
)

type Log struct {
	EventType         string `json:"eventType"`
	BuildId           string `json:"buildId"`
	BuildUrl          string `json:"buildUrl"`
	BuildTimestamp    string `json:"buildTimestamp"`
	Branch            string `json:"branch"`
	PullRequestId     string `json:"pullRequestId"`
	PullRequestLabels string `json:"pullRequestLabels"`
	Message           string `json:"message"`
}

func sendToNewRelicInsights(c *cli.Context) {
	// This information will always be passed from the codefresh pipeline
	buildId := c.String("CF_BUILD_ID")
	buildUrl := c.String("CF_BUILD_URL")
	buildTimestamp := c.String("CF_BUILD_TIMESTAMP")
	branch := c.String("CF_BRANCH")
	pullRequestId := c.String("CF_PULL_REQUEST_ID")

	// This parameter is provided by the codefresh but may not be populated depending on the git provider
	pullRequestLabels := c.String("CF_PULL_REQUEST_LABELS")

	// (Optional) default is US, change URL if user passes EU
	newRelicRegion := c.String("New-Relic-Region")
	baseUrl := "https://insights-collector.newrelic.com/v1/accounts/"
	if newRelicRegion == "EU" {
		// https://docs.newrelic.com/docs/using-new-relic/welcome-new-relic/get-started/introduction-eu-region-data-center#endpoints
		baseUrl = "https://insights-collector.eu01.nr-data.net/v1/accounts/"
	}

	// (Optional) Override the http post url for New Relic Insights
	newRelicInsightsUrlOverride := c.String("New-Relic-Insights-Url-Override")
	if newRelicInsightsUrlOverride != "" && newRelicInsightsUrlOverride != "${{New-Relic-Insights-Url-Override}}" {
		baseUrl = newRelicInsightsUrlOverride
	}

	// (Optional) Custom message for the log
	message := c.String("Message")

	// The account id must be populated to route to the appropriate insights data location
	accountId := c.String("New-Relic-Account-Id")

	// Build the fully qualified URL now that we have all the information
	eventUrl := "/events"
	extendedUrl := baseUrl + accountId + eventUrl

	// This insert key must be populated or you will get a 403 forbidden
	insertKey := c.String("X-Insert-Key")

	// Build the json payload for the event data
	jsonPayload := Log{
		EventType:         "Codefresh",
		BuildId:           buildId,
		BuildUrl:          buildUrl,
		BuildTimestamp:    buildTimestamp,
		Branch:            branch,
		PullRequestId:     pullRequestId,
		PullRequestLabels: pullRequestLabels,
		Message:           message,
	}
	jsonValue, err := json.Marshal(jsonPayload)
	if err != nil {
		fmt.Printf("Formatting JSON failed with error %s\n", err)
	}

	var byteBuffer bytes.Buffer
	g := gzip.NewWriter(&byteBuffer)
	g.Write(jsonValue)
	g.Close()

	// request, _ := http.NewRequest("POST", extendedUrl, bytes.NewBuffer(jsonValue))
	request, _ := http.NewRequest("POST", extendedUrl, &byteBuffer)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Encoding", "gzip")
	request.Header.Set("X-Insert-Key", insertKey)
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
