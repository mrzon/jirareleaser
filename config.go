package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// JiraConfig is a config to communicate with JIRA account.
type JiraConfig struct {
	Token  string `json:"token,omitempty"`
	Email  string `json:"email,omitempty"`
	UserID string `json:"user_id,omitempty"`
	Test   bool   `json:"test,omitempty"`
}

var defaultConfig *JiraConfig

//fetch default config
func getConfig() *JiraConfig {
	if defaultConfig != nil { //ensure only run once, if the content change, restart the program.
		return defaultConfig
	}

	file, err := os.Open("config.json")
	defer file.Close()

	if err == nil {
		jiraConfig := &JiraConfig{}

		byteContent, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteContent, jiraConfig)
		defaultConfig = jiraConfig
		return defaultConfig
	}

	return nil
}
