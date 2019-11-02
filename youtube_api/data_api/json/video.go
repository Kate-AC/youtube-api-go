package json

import (
  api_json "youtube_api/json"
)

type Video struct {}

func (video *Video) GetJson() api_json.JsonInterface {
  return &VideoJson{}
}

type VideoJson struct {
  Kind     string `json:"kind"`
  Etag     string `json:"etag"`
  PageInfo struct {
    TotalResults   int `json:"totalResults"`
    ResultsPerPage int `json:"resultsPerPage"`
  } `json:"pageInfo"`
  Items []struct {
    Kind       string `json:"kind"`
    Etag       string `json:"etag"`
    ID         string `json:"id"`
    Statistics struct {
      ViewCount     string `json:"viewCount"`
      LikeCount     string `json:"likeCount"`
      DislikeCount  string `json:"dislikeCount"`
      FavoriteCount string `json:"favoriteCount"`
      CommentCount  string `json:"commentCount"`
    } `json:"statistics"`
  } `json:"items"`
}

func (videoJson *VideoJson) GetResult() api_json.ResultInterface {
  item := (*videoJson).Items[0]
  return &VideoResult{
    item.ID,
    item.Statistics.ViewCount,
    item.Statistics.LikeCount,
    item.Statistics.DislikeCount,
    item.Statistics.FavoriteCount,
    item.Statistics.CommentCount,
  }
}

type VideoResult struct {
  VideoId       string
  ViewCount     string
  LikeCount     string
  DislikeCount  string
  FavoriteCount string
  CommentCount  string
}

