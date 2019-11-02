package repository

import (
  "db"
  "entity"
  "fmt"
)

type ChannelsRepository struct {
  Mysql db.Mysql
}

func (channelsRepository *ChannelsRepository) FindAll() []entity.Channel {
  var channels []entity.Channel
  rows := (*channelsRepository).Mysql.Select("select * from channels")

  for rows.Next() {
    channel := entity.Channel{}
    err := rows.Scan(
      &channel.Id,
      &channel.Name,
      &channel.Admin,
      &channel.CreatedAt,
      &channel.UpdatedAt)

    if err != nil {
      panic(err.Error())
    }
    channels = append(channels, channel)
  }

  return channels
}

func (channelsRepository *ChannelsRepository) FindOneById(channelId string) (bool, entity.Channel) {
  rows := (*channelsRepository).Mysql.Select(fmt.
    Sprintf("select * from channels where id = '%s' limit 1", channelId))

  channel := entity.Channel{}
  for rows.Next() {
    err := rows.Scan(
      &channel.Id,
      &channel.Name,
      &channel.Admin,
      &channel.CreatedAt,
      &channel.UpdatedAt)

    if err != nil {
      panic(err.Error())
    }
  }

  return channel.Id != "", channel
}

