package models

type adventure struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		}
	}
}
