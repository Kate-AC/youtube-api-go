package lib

import (
  "time"
  "fmt"
  "strconv"
)

type FormatTime struct {}

func (formatTime *FormatTime) Now() string {
  now := time.Now()

  return fmt.Sprintf("%s-%s-%s %s:%s:%s",
    formatTime.padding(now.Year()),
    formatTime.padding(int(now.Month())),
    formatTime.padding(now.Day()),
    formatTime.padding(now.Hour()),
    formatTime.padding(now.Minute()),
    formatTime.padding(now.Second()))
}

func (formatTime *FormatTime) padding(value int) string {
  if 10 > value {
    return "0" + strconv.Itoa(value)
  }
  return strconv.Itoa(value)
}
