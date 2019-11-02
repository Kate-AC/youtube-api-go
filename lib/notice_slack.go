package lib

import (
  "conf"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
)

type NoticeSlack struct {
  Config conf.Config
}

func (noticeSlack *NoticeSlack) Post(message string) {
  params := struct {
    Text      string `json:"text"`
    Username  string `json:"username"`
    IconEmoji string `json:"icon_emoji"`
    IconURL   string `json:"icon_url"`
    Channel   string `json:"channel"`
  }{
    message,
    noticeSlack.Config.SlackBotName,
    noticeSlack.Config.SlackIconEmoji,
    noticeSlack.Config.SlackIconUrl,
    noticeSlack.Config.SlackChannel,
  }

  jsonData, _ := json.Marshal(params)
  response, _ := http.PostForm(
    noticeSlack.Config.SlackWebhookUrl,
    url.Values{"payload": {string(jsonData)}},
  )

  ioutil.ReadAll(response.Body)
  defer response.Body.Close()
}
