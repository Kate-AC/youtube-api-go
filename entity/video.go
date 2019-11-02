package entity

import (
  "time"
  "database/sql"
)

type Video struct {
  Id string
  ChannelId string
  Title string
  PublishedAt time.Time
  OneHourLaterCount sql.NullString
  OneDayLaterCount sql.NullString
  CreatedAt time.Time
  UpdatedAt time.Time
}

