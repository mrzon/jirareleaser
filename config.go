package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// JiraConfig is a config to communicate with JIRA account.
type JiraConfig struct {
	Token     string `json:"token,omitempty"`
	Email     string `json:"email,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Test      bool   `json:"test,omitempty"`
	Project   string `json:"project"`
	Component string `json:"component"`
}

var defaultConfig *JiraConfig

//fetch default config
func getJiraConfig() *JiraConfig {
	if defaultConfig != nil { //ensure only run once, if the content change, restart the program.
		return defaultConfig
	}

	// To get the path if user store the binary in $PATH
	file, err := os.Open(getBinaryPath() + "/config.json")
	if err != nil {
		file, err = os.Open("config.json")
	}

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

func getBinaryPath() string {
	e, err := os.Executable()
	if err != nil {
		return ""
	}
	return fmt.Sprint(path.Dir(e))
}
