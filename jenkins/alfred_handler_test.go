package jenkins

import (
	"testing"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_it_will_get_the_status_from_jenkins(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(200, `{"jobs":[{"name":"TheService","Url":"theUrl","color":"TheColor","healthReport":[{"description":"TheDescription","iconUrl":"theIconUrl","score":"TheScore"}],"lastBuild":{"number":10,"result":"Success"}}]}`))

	handler := Handler{"http://fakeJenkins.com"}

	response:= handler.GetStatus("kloans-loan")

	assert.Equal(t, `<items><item uid="TheService" arg="theUrl"><title>TheService [Last Build #10 - Success]</title><subtitle>TheDescription</subtitle><icon>images/theIconUrl</icon></item></items>`, response)

}

func Test_it_will_return_No_Description_if_no_health_report_is_found(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(200, `{"jobs":[{"name":"TheService","Url":"theUrl","color":"TheColor","healthReport":[],"lastBuild":{"number":10,"result":"Success"}}]}`))

	handler := Handler{"http://fakeJenkins.com"}

	response := handler.GetStatus("kloans-loan")

	assert.Equal(t, `<items><item uid="TheService" arg="theUrl"><title>TheService [Last Build #10 - Success]</title><subtitle>No Description</subtitle></item></items>`, response)
}

func Test_it_will_return_error_item_if_request_error(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(788, `{"jobl":"theUrl","color":"TheColor","healthReport":[],"lastBuild":{"number":"10","result":"Success"}}]}`))

	handler := Handler{"http://fakeJenkins.com"}

	response := handler.GetStatus("kloans-loan")

	assert.Equal(t, `<items><item arg=""><title>Error occurred while fetching jenkins jobs.</title><subtitle>invalid character ']' after top-level value</subtitle></item></items>`, response)
}