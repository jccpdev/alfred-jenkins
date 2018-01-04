package jenkins

import "encoding/xml"

type Feedback struct {
	XMLName xml.Name `xml:"items"`
	Item    []Item   `xml:"item"`
}

func (fb *Feedback) addItem(item Item) {
	fb.Item = append(fb.Item, item)
}

// Explanation of xml structure found here https://www.alfredforum.com/topic/5-generating-feedback-in-workflows/
type Item struct {
	Uid          string `xml:"uid,attr,omitempty"`
	Arg          string `xml:"arg,attr"`
	Valid        string `xml:"valid,attr,omitempty"`
	AutoComplete string `xml:"autocomplete,attr,omitempty"`
	Title        string `xml:"title"`
	Subtitle     string `xml:"subtitle"`
	Icon         string `xml:"icon,omitempty"`
}

