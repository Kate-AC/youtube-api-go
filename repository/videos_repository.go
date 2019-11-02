package repository

import (
  "db"
  "entity"
  "fmt"
  "time"
  "youtube_api/data_api/json"
)

type VideosRepository struct {
  Mysql db.Mysql
}

func (videosRepository *VideosRepository) FindAllStillNotCountedOneHourVideo() []entity.Video {
  var videos []entity.Video
  rows := (*videosRepository).Mysql.Select(`
    select *
    from videos
    where (published_at <= (now() - interval 1 hour))
    and one_hour_later_count is null
  `)

  for rows.Next() {
    video := entity.Video{}
    err := rows.Scan(
      &video.Id,
      &video.ChannelId,
      &video.Title,
      &video.PublishedAt,
      &video.OneHourLaterCount,
      &video.OneDayLaterCount,
      &video.CreatedAt,
      &video.UpdatedAt)

    if err != nil {
      panic(err.Error())
    }
    videos = append(videos, video)
  }

  return videos
}

func (videosRepository *VideosRepository) FindAllStillNotCountedOneDayVideo() []entity.Video {
  var videos []entity.Video
  rows := (*videosRepository).Mysql.Select(`
    select *
    from videos
    where (published_at <= (now() - interval 1 day))
    and one_day_later_count is null
  `)

  for rows.Next() {
    video := entity.Video{}
    err := rows.Scan(
      &video.Id,
      &video.ChannelId,
      &video.Title,
      &video.PublishedAt,
      &video.OneHourLaterCount,
      &video.OneDayLaterCount,
      &video.CreatedAt,
      &video.UpdatedAt)

    if err != nil {
      panic(err.Error())
    }
    videos = append(videos, video)
  }

  return videos
}

func (videosRepository *VideosRepository) FindOneById(id string) (bool, entity.Video) {
  rows := (*videosRepository).Mysql.Select(fmt.
    Sprintf("select * from videos where id = '%s' limit 1", id))

  video := entity.Video{}
  for rows.Next() {
    err := rows.Scan(
      &video.Id,
      &video.ChannelId,
      &video.Title,
      &video.PublishedAt,
      &video.OneHourLaterCount,
      &video.OneDayLaterCount,
      &video.CreatedAt,
      &video.UpdatedAt)

    if err != nil {
      panic(err.Error())
    }
  }

  return video.Id != "", video
}

func (videosRepository *VideosRepository) CreateVideo(latestVideoResult *json.LatestVideoResult) {
  (*videosRepository).Mysql.Execute(fmt.
    Sprintf(`
      insert into videos(
        id,
        channel_id,
        title,
        published_at,
        one_hour_later_count,
        one_day_later_count,
        created_at,
        updated_at
      ) values (?, ?, ?, ?, null, null, now(), now())
    `),
    latestVideoResult.VideoId,
    latestVideoResult.ChannelId,
    latestVideoResult.Title,
    latestVideoResult.PublishedAt.Format("2006-01-02 15:04:05"))
}

func (videosRepository *VideosRepository) UpdateOneHourLaterCount(viewCount string, videoId string) {
  (*videosRepository).Mysql.Execute(fmt.
    Sprintf(`
      update videos
      set one_hour_later_count = ?, updated_at = ?
      where id = ?
    `),
    viewCount,
    time.Now().Format("2006-01-02 15:04:05"),
    videoId)
}

func (videosRepository *VideosRepository) UpdateOneDayLaterCount(viewCount string, videoId string) {
  (*videosRepository).Mysql.Execute(fmt.
    Sprintf(`
      update videos
      set one_day_later_count = ?, updated_at = ?
      where id = ?
    `),
    viewCount,
    time.Now().Format("2006-01-02 15:04:05"),
    videoId)
}

