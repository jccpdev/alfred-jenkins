package jenkins

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/xml"
	"strconv"
	"regexp"
)

type Handler struct {
	Url string
}

func (handler *Handler) GetStatus(query string) (string) {

	fb := Feedback{}

	url := handler.Url + "/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]"

	response, requestError := http.Get(url)

	if requestError != nil {
		return MakeError(requestError)
	}

	defer response.Body.Close()

	responseBytes, ioError := ioutil.ReadAll(response.Body)

	if ioError != nil {
		return MakeError(ioError)
	}

	res := Response{}

	unMarshalError := json.Unmarshal(responseBytes, &res)

	if unMarshalError != nil {
		return MakeError(unMarshalError)
	}

	for _, job := range res.Jobs {

		if job.Name != "" {

			matched, matchError := regexp.Match("(?i)("+query+")", []byte(job.Name))

			if matchError != nil {
				return MakeError(matchError)
			}

			if matched == true {
				fb.addItem(MakeItem(job))
			}

		}

	}

	fbString, err := xml.Marshal(fb)

	if err != nil {
		return MakeError(err)
	}

	return string(fbString)

}

func MakeItem(jobData Jobs) Item {

	var description string
	var iconUrl string = ""

	if len(jobData.HealthReport) > 0 {
		description = jobData.HealthReport[0].Description
		iconUrl = "images/" + jobData.HealthReport[0].IconUrl
	} else {
		description = "No Description"
	}

	buildNumber := strconv.Itoa(jobData.LastBuild.Number)

	title := jobData.Name + " [Last Build #" + buildNumber + " - " + jobData.LastBuild.Result + "]"

	return Item{jobData.Name, jobData.Url, "", "", title, description, iconUrl}

}

func MakeError(error error) string {
	return `<items><item arg=""><title>Error occurred while fetching jenkins jobs.</title><subtitle>` + error.Error() + `</subtitle></item></items>`
}
