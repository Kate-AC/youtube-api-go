package conf

type Config struct {
  ApiKey string
  EndPoint string
  DbName string
  DbHost string
  DbPort int
  DbUser string
  DbPassword string
  DbTimeZone string
  SlackWebhookUrl string
  SlackBotName string
  SlackIconEmoji string
  SlackIconUrl string
  SlackChannel string
}

func NewConfig() Config {
  return Config{
    "xxxxxxxxxxx",
    "https://www.googleapis.com/youtube/v3",
    "xxxxxxxxxxx",
    "mysql",
    3306,
    "root",
    "root",
    "Asia/Tokyo",
    "https://hooks.slack.com/services/xxxxxxxxxxxxxxxx",
    "xxxxxxx",
    ":star:",
    "",
    "xxxxxxxx",
  }
}

