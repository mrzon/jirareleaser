package main

import "fmt"

func getTemplate(jiraConfig *JiraConfig, releaseDate string, service string) string {
	template := `
	{
		"fields": {
			"project": {
				"key": "%s"
			},
			"summary": "[RELEASE][%s][%s][Main Release]",
			"description": "Check related issues for changelog.\nTime and Version of Release:\nTime: Timestamp of actual release to production\nOld version:\nNew version:\n\nDependency:\n1. ..\n2. ..\n\nRisk:\nPlease give warning on this JIRA task if your change have major risk (new major feature, service refactors, major config change, etc)\n\nReasons of Rollback:\nAdd reasons of rollback/ link to document if any\n...\n...\n\nPowered by [jiraReleaser|https://29022131.atlassian.net/wiki/spaces/GO/pages/1266488670/JIRA+Releaser+CLI+Program]",
			"issuetype": {
				"name": "Task"
			},
			"components": [
				{
				"name": "BACKEND"
				}
			],
			"assignee": {
				"name": "%s"
			}
		}
	}`

	return fmt.Sprintf(template, jiraConfig.Project, releaseDate, service, jiraConfig.UserID)
}
