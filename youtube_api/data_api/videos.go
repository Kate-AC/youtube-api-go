package data_api

import (
  "conf"
  "youtube_api"
  data_api_json "youtube_api/data_api/json"
)

type Videos struct {
  Config conf.Config
}

func (videos *Videos) GetVideo(videoId string) youtube_api.ResponseInterface {
  params := youtube_api.Params{}
  params.AddParam("part", "statistics")
  params.AddParam("id", videoId)
  params.AddParam("key", (*videos).Config.ApiKey)

  request := youtube_api.Request{
    "GET",
    "/videos",
    (*videos).Config,
    params.GetParams(),
    (&data_api_json.Video{}).GetJson(),
  }

  return request.Execute()
}
