package data_api

import (
  "conf"
  "youtube_api"
  data_api_json "youtube_api/data_api/json"
)

type Search struct {
  Config conf.Config
}

func (search *Search) GetLatestVideo(channelId string) youtube_api.ResponseInterface {
  params := youtube_api.Params{}
  params.AddParam("part", "snippet")
  params.AddParam("type", "video")
  params.AddParam("maxResults", "1")
  params.AddParam("order", "date")
  params.AddParam("channelId", channelId)
  params.AddParam("key", (*search).Config.ApiKey)

  request := youtube_api.Request{
    "GET",
    "/search",
    (*search).Config,
    params.GetParams(),
    (&data_api_json.LatestVideo{}).GetJson(),
  }

  return request.Execute()
}

