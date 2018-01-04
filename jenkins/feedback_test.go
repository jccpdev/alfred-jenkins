package jenkins

import (
	"testing"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
)

func Test_it_can_make_a_feedback_xml_structure(t *testing.T) {

	expectedUid := "1234"
	expectedArg := "the arg"
	expectedValid := "yes"
	expectedAutoComplete := "The auto complete"
	expectedTitle := "The Title"
	expectedSubTitle := "The Subtitle"
	expectedIcon := "theIcon"

	fb := &Feedback{}

	fb.addItem(Item{
		expectedUid,
		expectedArg,
		expectedValid,
		expectedAutoComplete,
		expectedTitle,
		expectedSubTitle,
		expectedIcon})

	output, err := xml.Marshal(fb)

	if err != nil {
		t.Fatal(err)
	}

	resultFb := Feedback{}

	resultError := xml.Unmarshal(output, &resultFb)

	if resultError != nil {
		t.Fatal(resultError)
	}

	os.Stdout.Write(output)

	fmt.Printf("%s\n", resultFb.Item[0].Title)

	assert.Equal(t, resultFb.Item[0].Uid, expectedUid)
	assert.Equal(t, resultFb.Item[0].Arg, expectedArg)
	assert.Equal(t, resultFb.Item[0].Valid, expectedValid)
	assert.Equal(t, resultFb.Item[0].AutoComplete, expectedAutoComplete)
	assert.Equal(t, resultFb.Item[0].Title, expectedTitle)
	assert.Equal(t, resultFb.Item[0].Subtitle, expectedSubTitle)
	assert.Equal(t, resultFb.Item[0].Icon, expectedIcon)

}

func Test_make_item_with_optional_params(t *testing.T) {

	expectedArg := "the arg"
	expectedTitle := "The Title"
	expectedSubTitle := "The Subtitle"

	fb := &Feedback{}

	fb.addItem(Item{
		"",
		expectedArg,
		"",
		"",
		expectedTitle,
		expectedSubTitle,
		""})

	output, err := xml.Marshal(fb)

	if err != nil {
		t.Fatal(err)
	}

	resultFb := Feedback{}

	resultError := xml.Unmarshal(output, &resultFb)

	if resultError != nil {
		t.Fatal(resultError)
	}

	os.Stdout.Write(output)

	fmt.Printf("%s\n", resultFb.Item[0].Title)

	assert.Equal(t, resultFb.Item[0].Uid, "")
	assert.Equal(t, resultFb.Item[0].Arg, expectedArg)
	assert.Equal(t, resultFb.Item[0].Valid, "")
	assert.Equal(t, resultFb.Item[0].AutoComplete, "")
	assert.Equal(t, resultFb.Item[0].Title, expectedTitle)
	assert.Equal(t, resultFb.Item[0].Subtitle, expectedSubTitle)
	assert.Equal(t, resultFb.Item[0].Icon, "")

}
