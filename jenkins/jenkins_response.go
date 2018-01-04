package jenkins

type Response struct {
	Jobs []Jobs `json:"jobs"`
}

type Jobs struct {
	Name         string         `json:"name"`
	Url          string         `json:"Url"`
	Color        string         `json:"color"`
	HealthReport []HealthReport `json:"healthReport"`
	LastBuild    LastBuild      `json:"lastBuild"`
}

type HealthReport struct {
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Score       int    `json:"icon"`
}

type LastBuild struct {
	Number int    `json:"number"`
	Result string `json:"result"`
}
