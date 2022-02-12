package drmfilter

type Response struct {
	Response []ResponseItem `json:"response"` // not sure, should we omitempty here?
}

type ResponseItem struct {
	Image string `json:"image"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type Request struct {
	Payload []struct {
		Country      string `json:"country,omitempty"`
		Description  string `json:"description,omitempty"`
		Drm          bool   `json:"drm,omitempty"`
		EpisodeCount int    `json:"episodeCount,omitempty"`
		Genre        string `json:"genre,omitempty"`
		Image        struct {
			ShowImage string `json:"showImage"`
		} `json:"image,omitempty"`
		Language      string      `json:"language,omitempty"`
		NextEpisode   interface{} `json:"nextEpisode,omitempty"`
		PrimaryColour string      `json:"primaryColour,omitempty"`
		Seasons       []struct {
			Slug string `json:"slug"`
		} `json:"seasons,omitempty"`
		Slug      string `json:"slug"`
		Title     string `json:"title"`
		TvChannel string `json:"tvChannel"`
	} `json:"payload"`
	Skip         int `json:"skip"`
	Take         int `json:"take"`
	TotalRecords int `json:"totalRecords"`
}
