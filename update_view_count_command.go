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

  client       := data_api.DataApiClient{config}
  videosClient := client.Videos()
  noticeSlack  := lib.NoticeSlack{config}

  videos := videosRepository.FindAllStillNotCountedOneHourVideo()

  for _, video := range videos {
    videoResponse := videosClient.GetVideo(video.Id)
    videoBody     := videoResponse.GetBody()
    videoResult   := videoBody.(*data_api_json.VideoResult)
    videosRepository.UpdateOneHourLaterCount(videoResult.ViewCount, video.Id)

    _, channel := channelsRepository.FindOneById(video.ChannelId)

    message := fmt.Sprintf(`
      *【%sの最新動画の1時間後の再生数を取得しました。】*
      タイトル: %s
      URL: https://www.youtube.com/watch?v=%s
      *再生数: %s*`,
      channel.Name, video.Title, video.Id, videoResult.ViewCount)

    noticeSlack.Post(message)
    fmt.Println(channel.Id + " " + video.Id)
  }

  videos = videosRepository.FindAllStillNotCountedOneDayVideo()

  for _, video := range videos {
    videoResponse := videosClient.GetVideo(video.Id)
    videoBody     := videoResponse.GetBody()
    videoResult   := videoBody.(*data_api_json.VideoResult)
    videosRepository.UpdateOneDayLaterCount(videoResult.ViewCount, video.Id)

    _, channel := channelsRepository.FindOneById(video.ChannelId)

    message := fmt.Sprintf(`
      *【%sの最新動画の1日後の再生数を取得しました。】*
      タイトル: %s
      URL: https://www.youtube.com/watch?v=%s
      *再生数: %s*`,
      channel.Name, video.Title, video.Id, videoResult.ViewCount)

    noticeSlack.Post(message)
    fmt.Println(channel.Id + " " + video.Id)
  }
}

