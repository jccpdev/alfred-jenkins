package jenkins

import (
	"testing"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_it_will_get_the_status_from_jenkins(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(200, `{"jobs":[{"name":"TheService","Url":"theUrl","color":"TheColor","healthReport":[{"description":"TheDescription","iconUrl":"theIconUrl","score":"TheScore"}],"lastBuild":{"number":10,"result":"Success"}}]}`))

	handler := Handler{"http://fakeJenkins.com"}

	response:= handler.GetStatus("TheService")

	assert.Equal(t, `<items><item uid="TheService" arg="theUrl"><title>TheService [Last Build #10 - Success]</title><subtitle>TheDescription</subtitle><icon>images/theIconUrl</icon></item></items>`, response)

}

func Test_it_will_filter_the_search_results_with_query(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(200, `{"jobs":[{"name":"Green Job","Url":"theUrl","color":"TheColor","healthReport":[{"description":"TheDescription","iconUrl":"theIconUrl","score":"TheScore"}],"lastBuild":{"number":10,"result":"Success"}},{"name":"Orange Job","Url":"theUrl","color":"TheColor","healthReport":[{"description":"TheDescription","iconUrl":"theIconUrl","score":"TheScore"}],"lastBuild":{"number":10,"result":"Success"}}]}`))

	handler := Handler{"http://fakeJenkins.com"}

	response:= handler.GetStatus("Orange")

	fmt.Println(response)

	assert.Equal(t, `<items><item uid="Orange Job" arg="theUrl"><title>Orange Job [Last Build #10 - Success]</title><subtitle>TheDescription</subtitle><icon>images/theIconUrl</icon></item></items>`, response)

}


func Test_it_will_return_No_Description_if_no_health_report_is_found(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(200, `{"jobs":[{"name":"TheService","Url":"theUrl","color":"TheColor","healthReport":[],"lastBuild":{"number":10,"result":"Success"}}]}`))

	handler := Handler{"http://fakeJenkins.com"}

	response := handler.GetStatus("TheService")

	assert.Equal(t, `<items><item uid="TheService" arg="theUrl"><title>TheService [Last Build #10 - Success]</title><subtitle>No Description</subtitle></item></items>`, response)
}

func Test_it_will_return_error_item_if_request_error(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(788, `{"jobl":"theUrl","color":"TheColor","healthReport":[],"lastBuild":{"number":"10","result":"Success"}}]}`))

	handler := Handler{"h://fakeJenkins.com"}

	response := handler.GetStatus("kloans-loan")

	assert.Equal(t, `<items><item arg=""><title>Error occurred while fetching jenkins jobs.</title><subtitle>Get h://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]: no responder found</subtitle></item></items>`, response)
}

func Test_it_will_return_error_item_if_parsing_error(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://fakeJenkins.com/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]",
		httpmock.NewStringResponder(788, `{"jobl":"th`))

	handler := Handler{"http://fakeJenkins.com"}

	response := handler.GetStatus("kloans-loan")

	assert.Equal(t, `<items><item arg=""><title>Error occurred while fetching jenkins jobs.</title><subtitle>unexpected end of JSON input</subtitle></item></items>`, response)
}