package main

import (
	"fmt"
	"os"
)

type jiraReleaserData struct {
	config   *JiraConfig
	services []string
	date     string
}

func main() {
	appData := parseJiraReleaserData(os.Args, getJiraConfig())

	if appData == nil {
		fmt.Println("JIRA Release needs at least date and one released service")
		fmt.Println("How to use: jirareleaser [--p | -project <JIRA Project>] YYYY-MM-DD SERVICE_1 SERVICE_2 ... SERVICE_N")
		return
	}

	if appData.config == nil {
		fmt.Println("JIRA Releaser needs configuration file. Please make one.")
		return
	}

	for _, service := range appData.services {
		template := getTemplate(appData.config, appData.date, service)
		response := jiracall(service, template, appData.config)

		fmt.Println(service + " " + appData.date + " https://29022131.atlassian.net/browse/" + response)
	}
}

func parseJiraReleaserData(args []string, config *JiraConfig) *jiraReleaserData {
	numOfArgs := len(args)

	if numOfArgs < 3 {

		return nil
	}

	flagIndex := 1
	projectFlag := args[1]
	projectInFlag := ""
	if projectFlag == "--p" || projectFlag == "-project" {
		flagIndex = 3
		projectInFlag = args[2]

		if numOfArgs < 5{
			return nil
		}
	}

	releaseDate := args[flagIndex]
	var services = args[flagIndex+1:]

	if config != nil && projectInFlag != "" {
		config.Project = projectInFlag
	}

	return &jiraReleaserData{
		config:   config,
		services: services,
		date:     releaseDate,
	}
}
