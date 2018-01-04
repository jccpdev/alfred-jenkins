package jenkins

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func Test_it_can_parse_a_jenkins_json_response(t *testing.T){

	fakeResponseJson := `{"jobs":[{"name":"TheService","Url":"theUrl","color":"TheColor","healthReport":[{"description":"TheDescription","iconUrl":"theIconUrl","score":"TheScore"}],"lastBuild":{"number":"10","result":"Success"}},{"name":"TheService2","Url":"theUrl","color":"TheColor","healthReport":[{"description":"TheDescription","iconUrl":"theIconUrl","score":"TheScore"}],"lastBuild":{"number":"10","result":"Success"}}]}`

	res := Response{}
	json.Unmarshal([]byte(fakeResponseJson), &res)

	assert.Len(t, res.Jobs, 2)

}