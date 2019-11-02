package entity

import (
  "time"
)

type Channel struct {
  Id string
  Name string
  Admin int
  CreatedAt time.Time
  UpdatedAt time.Time
}

