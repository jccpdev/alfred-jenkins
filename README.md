# Alfred Jenkins

## How to use

1) Download alfred workflow file and double click it or add it to jenkins
2) Open input keyword workflow and put jenkins url with username and api key where it says `"JENKINS_URL_HERE"` 

Sample jenkins url `http://username:api-key@jenkins.your-jenkins-server.com`

## How to edit

1) Have go installed
2) Clone Project
3) Install the go dependencies listed bellow
4) When done making modifications run `go build`
5) Copy the `alfred-jenkins/jenkins/alfred-jenkins/alfred-jenkins` file and move it to the root directory of the installed alfred workflow.

### Required External Goland Packages
* "github.com/jarcoal/httpmock"
* "github.com/stretchr/testify/assert"
* "github.com/jccpdev/alfred-go"