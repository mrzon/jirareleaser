package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//Response from JIRA API {"id":"410128","key":"STL-3811","self":"https://29022131.atlassian.net/rest/api/2/issue/410128

type JiraResponse struct {
	Id   string `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Self string `json:"self,omitempty"`
}

func jiracall(service, content string, config *JiraConfig) string {

	if config.Test {
		return "TEST-" + service
	}

	url := "https://29022131.atlassian.net/rest/api/2/issue/"

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(content))

	if err != nil {
		return "Failed to create Jira Task for Service " + service
	}
	auth := config.Email + ":" + config.Token

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	resp, err := client.Do(req)

	if err != nil {
		return "Failed to create Jira Task for Service " + service
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)

	jiraResponse := &JiraResponse{}
	json.Unmarshal(respBytes, jiraResponse)
	return jiraResponse.Key
}
