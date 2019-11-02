package youtube_api

import (
  "conf"
  "fmt"
  "net/http"
  "net/url"
  "os"
  api_json "youtube_api/json"
)

type RequestInterface interface {
  Execute() ResponseInterface
}

type Request struct {
  Method     string
  Resource   string
  Config     conf.Config
  Params     []Param
  JsonStruct api_json.JsonInterface
}

func (request *Request) Execute() ResponseInterface {
  values := url.Values{}
  list   := (*request).Params

  for i := 0; i < len(list); i++ {
    values.Add(list[i].key, list[i].value)
  }

  endPoint := (*request).Config.EndPoint + (*request).Resource

  // TODO: POSTとかにも分岐したい
  result, err := http.Get(endPoint + "?" + values.Encode())

  if err != nil {
    fmt.Printf("Response failed.")
    os.Exit(1)
  }

  return &Response{
    result,
    (*request).JsonStruct,
  }
}
