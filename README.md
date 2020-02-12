JIRA Release

How do you plan a release for your service?

- Copy other released JIRA Task, 
- Copy existing template, or
- Write one every time

The process sometimes sequential towards the amount of service that you want to release in that week. A bit mundane, but necessary.

That's what jirarelease program do for you. 

- Specify Date of release
- Specify services that you want to release

1. Download and extract jirarelease to a folder
2. Ensure the binary is executable (`chmod +x ./jirarelease` if necessary)
3. Rename config-template.json to config.json
4. Fill up config.json with your actual information (use https://confluence.atlassian.com/cloud/api-tokens-938839638.html to retrieve the JIRA API Token)
5. Execute jirarelease to create a release JIRA task for you `./jirarelease YYYY-MM-DD <SERVICE1> <SERVICE2> ... <SERVICEN>`
  
It will automatically generated the story. After that, you can just fill up other manual information, such as the version, release time, and the release dependency.

Jirarelease utilizes your account to call JIRA backend API in the background, so you need to pass some config in order for it to be work.

PS: Can you think a better name for this simple tool? Please DM me for your idea. 

:golang: