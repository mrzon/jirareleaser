package main

import (
	"reflect"
	"testing"
)

func Test_parseJiraReleaserData(t *testing.T) {
	var tests = []struct {
		name  string
		args1 []string
		args2 *JiraConfig
		want  *jiraReleaserData
	}{
		{
			"Empty param",
			[]string{""},
			nil,
			nil,
		},
		{
			"No project as param",
			[]string{"", "YYY", "X", "Y", "Z"},
			&JiraConfig{Project: "ABC"},
			&jiraReleaserData{
				config:   &JiraConfig{Project: "ABC"},
				services: []string{"X", "Y", "Z"},
				date:     "YYY",
			},
		},
		{
			"Project as param",
			[]string{"", "--p", "STL", "YYY", "X", "Y", "Z"},
			&JiraConfig{Project: "ABC"},
			&jiraReleaserData{
				config:   &JiraConfig{Project: "STL"},
				services: []string{"X", "Y", "Z"},
				date:     "YYY",
			},
		},
		{
			"Project as param2",
			[]string{"", "-project", "STL", "YYY", "X", "Y", "Z"},
			&JiraConfig{Project: "ABC"},
			&jiraReleaserData{
				config:   &JiraConfig{Project: "STL"},
				services: []string{"X", "Y", "Z"},
				date:     "YYY",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseJiraReleaserData(tt.args1, tt.args2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseJiraReleaserData() = %v, want %v", got, tt.want)
			}
		})
	}
}
