package youtube

type playerResponse struct {
	StreamingData struct {
		Formats []struct {
			URL              string `json:"url"`
			MimeType         string `json:"mimeType"`
			AudioQuality     string `json:"audioQuality"`
			ApproxDurationMs string `json:"approxDurationMs,omitempty"`
		} `json:"formats"`
		AdaptiveFormats []struct {
			URL       string `json:"url"`
			MimeType  string `json:"mimeType"`
			InitRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"initRange"`
			IndexRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"indexRange"`
			ApproxDurationMs string `json:"approxDurationMs"`
			AudioQuality     string `json:"audioQuality,omitempty"`
		} `json:"adaptiveFormats"`
	} `json:"streamingData"`
	VideoDetails struct {
		VideoID          string `json:"videoId"`
		Title            string `json:"title"`
		LengthSeconds    string `json:"lengthSeconds"`
		ShortDescription string `json:"shortDescription"`
		IsLiveContent    bool   `json:"isLiveContent"`
	} `json:"videoDetails"`
}
