package main

import (
  "conf"
  "db"
  "fmt"
  "lib"
  "repository"
  "youtube_api/data_api"
  data_api_json "youtube_api/data_api/json"
)

func main() {
  config := conf.NewConfig()
  mysql  := db.Mysql{config, nil}

  videosRepository   := repository.VideosRepository{mysql}
  channelsRepository := repository.ChannelsRepository{mysql}

  client := data_api.DataApiClient{config}
  noticeSlack := lib.NoticeSlack{config}

  channels := channelsRepository.FindAll()
  for _, channel := range channels {
    search              := client.Search()
    latestVideoResponse := search.GetLatestVideo(channel.Id)
    latestVideoBody     := latestVideoResponse.GetBody()
    latestVideoResult   := latestVideoBody.(*data_api_json.LatestVideoResult)

    found, _ := videosRepository.FindOneById(latestVideoResult.VideoId)

    if false == found {
      videosRepository.CreateVideo(latestVideoResult)

      message := fmt.Sprintf(`
        【%sの最新動画が公開されました。】
        タイトル: %s
        URL: https://www.youtube.com/watch?v=%s`,
        channel.Name, latestVideoResult.Title, latestVideoResult.VideoId)

      noticeSlack.Post(message)

      fmt.Println(channel.Id + " " + latestVideoResult.VideoId)
    } else {
      fmt.Println(channel.Id + " is nothing new videos.")
    }
  }
}

    /*
    videos        := client.Videos()
    videoResponse := videos.GetVideo(latestVideoResult.VideoId)
    videoBody     := videoResponse.GetBody()
    videoResult   := videoBody.(*youtube_api.VideoResult)
    fmt.Println(videoResult)
    */

    //_, videos := videosRepository.FindAll()
    /*
    */


