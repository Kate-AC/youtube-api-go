package json

import (
  "time"
  api_json "youtube_api/json"
)

type LatestVideo struct {}

func (latestVideo *LatestVideo) GetJson() api_json.JsonInterface {
  return &LatestVideoJson{}
}

type LatestVideoJson struct {
  Kind          string `json:"kind"`
  Etag          string `json:"etag"`
  NextPageToken string `json:"nextPageToken"`
  RegionCode    string `json:"regionCode"`
  PageInfo      struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []struct {
    Kind string `json:"kind"`
    Etag string `json:"etag"`
    ID   struct {
      Kind    string `json:"kind"`
      VideoID string `json:"videoId"`
    } `json:"id"`
    Snippet struct {
      PublishedAt time.Time `json:"publishedAt"`
      ChannelID   string    `json:"channelId"`
      Title       string    `json:"title"`
      Description string    `json:"description"`
      Thumbnails  struct {
        Default struct {
          URL    string `json:"url"`
          Width  int    `json:"width"`
          Height int    `json:"height"`
        } `json:"default"`
        Medium struct {
          URL    string `json:"url"`
          Width  int    `json:"width"`
          Height int    `json:"height"`
        } `json:"medium"`
        High struct {
          URL    string `json:"url"`
          Width  int    `json:"width"`
          Height int    `json:"height"`
        } `json:"high"`
      } `json:"thumbnails"`
      ChannelTitle         string `json:"channelTitle"`
      LiveBroadcastContent string `json:"liveBroadcastContent"`
    } `json:"snippet"`
  } `json:"items"`
}

func (latestVideoJson *LatestVideoJson) GetResult() api_json.ResultInterface {
  item := (*latestVideoJson).Items[0]
  return &LatestVideoResult{
    item.Snippet.ChannelID,
    item.Snippet.ChannelTitle,
    item.Snippet.Title,
    item.ID.VideoID,
    item.Snippet.PublishedAt,
  }
}

type LatestVideoResult struct {
  ChannelId    string
  ChannelTitle string
  Title        string
  VideoId      string
  PublishedAt  time.Time
}

