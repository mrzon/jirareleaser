package main

import (
	"fmt"
	"os"
)

func main() {
	numOfArgs := len(os.Args)

	if numOfArgs < 3 {
		fmt.Println("JIRA Release needs at least date and one released service")
		return
	}

	releaseDate := os.Args[1]
	var services []string = os.Args[2:]

	jiraConfig := getConfig()
	for _, service := range services {
		template := getTemplate(jiraConfig.UserID, releaseDate, service)
		response := jiracall(service, template, jiraConfig)

		fmt.Println(service + " " + releaseDate + " https://29022131.atlassian.net/browse/" + response)
	}

}
