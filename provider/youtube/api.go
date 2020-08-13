package youtube

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deryrahman/player-cli/provider"
)

const (
	URL_META     = "https://www.youtube.com/get_video_info?&video_id="
	AUDIO_MEDIUM = "AUDIO_QUALITY_MEDIUM"
	AUDIO_LOW    = "AUDIO_QUALITY_LOW"
)

type youtubeProvider struct{}

func NewYoutubeProvider() provider.Provider {
	return &youtubeProvider{}
}

// GetAudio implementation for youtube provider
// currently only support query from exavt videoID
func (p *youtubeProvider) GetAudio(videoID string) (*provider.AudioMeta, error) {
	resp, err := getPlayer(videoID)
	if err != nil {
		return nil, err
	}

	for _, f := range resp.StreamingData.AdaptiveFormats {
		if strings.Contains(f.MimeType, "audio/") && f.AudioQuality == AUDIO_LOW && f.URL != "" {
			audioMeta := &provider.AudioMeta{
				MimeType: f.MimeType,
				Title:    resp.VideoDetails.Title,
				URL:      f.URL,
			}

			return audioMeta, nil
		}
	}

	return nil, provider.ErrNotFound
}

func fetchMeta(videoID string) (string, error) {
	resp, err := http.Get(URL_META + videoID)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	query, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(query), nil
}

func getPlayer(videoID string) (*playerResponse, error) {
	query, err := fetchMeta(videoID)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("?" + query)
	if err != nil {
		return nil, err
	}

	var resp playerResponse
	if err := json.Unmarshal([]byte(u.Query().Get("player_response")), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
